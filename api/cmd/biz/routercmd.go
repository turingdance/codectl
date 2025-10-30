package biz

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/spf13/cobra"
	"github.com/turingdance/codectl/app/logic"
	"github.com/turingdance/infra/slicekit"
	"github.com/turingdance/infra/stringx"
)

var dirsrc string = ""
var dirdst string = ""
var author string = ""
var routerfile string = "router.go"
var tplfile string = "./router.tpl"

// 排除在外的
var excludes []string = []string{}

type Route struct {
	Package string
	Module  string
	Func    string
	Path    string
	Method  []string
	Comment string
}

func (p *Route) Println() {
	fmt.Printf("moduel=%s,func=%s,path=%s,method=%s,comment=%s\n", p.Module, p.Func, p.Path, p.Method, p.Comment)
}

type RouteTree struct {
	Route
	Children []*Route
}

// func (ctrl *Account ) Register (
var regfunc *regexp.Regexp = regexp.MustCompile(`\s*func\s+\(\s*\w+\s+\*([A-Z]+\w+)\s*\)\s+([A-Z]+\w+)\s*\(.*`)

// @Summary 注册用户
var regsummary *regexp.Regexp = regexp.MustCompile(`\s*@Summary\s+(.*)`)

// @Router /account/register [POST,GET]
var regrouter *regexp.Regexp = regexp.MustCompile(`\s*@Router\s+(\S+)\s+(\[?.*\]?)`)

// type Account struct{}
var regstruct *regexp.Regexp = regexp.MustCompile(`\s*type\s+(\S+)\s+struct\s*\{.*`)

// package Account
var regpackage *regexp.Regexp = regexp.MustCompile(`\s*package\s+(\S+)`)

// // 简单描述
var regcomment *regexp.Regexp = regexp.MustCompile(`.*\/\/\s*(.*).*`)

// 下划线单词转为小写驼峰单词
func camel(s string) string {
	return stringx.CamelLcFirst(s)
}
func newroute(pkg string) *Route {
	return &Route{
		Package: pkg,
		Method:  []string{"GET", "POST", "OPTIONS"},
	}
}

// 解析所有文件,构建信息结构体
func buildroutes(dirsrc string, excludes ...string) (routes []*Route, err error) {
	if dirsrc == "" {
		err = errors.New("请指定接口代码路径")
		return
	}
	routes = make([]*Route, 0)
	err = filepath.WalkDir(dirsrc, func(fpath string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return err
		}
		//basepath := strings.ReplaceAll(fpath, dirsrc, "")
		// 过滤掉需要隔离的文件
		if slicekit.Some(excludes, func(arr []string, ele string) bool {
			if strings.Contains(fpath, ele) {
				return true
			} else {
				return false
			}
		}) {
			return nil
		}
		bts, err := os.ReadFile(fpath)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(bytes.NewReader(bts))
		var proute *Route = nil
		pkg := ""
		_ = pkg
		for scanner.Scan() {
			txt := scanner.Text()
			// 开始
			if regpackage.MatchString(txt) { //package
				arr := regpackage.FindStringSubmatch(txt)
				pkg = arr[1]
				proute = newroute(pkg)
			} else if regstruct.MatchString(txt) { // 结构体
				structs := regstruct.FindStringSubmatch(txt)
				proute.Module = structs[1]
				routes = append(routes, proute)
				// 开启新的结构体
				proute = newroute(pkg)
			} else if regfunc.MatchString(txt) { // 碰到函数
				funcs := regfunc.FindStringSubmatch(txt)
				proute.Module = funcs[1]
				proute.Func = funcs[2]
				// 如果是空的
				if proute.Path == "" {
					proute.Path = path.Join("//", camel(proute.Module), camel(proute.Func))
				}
				routes = append(routes, proute)
				// 开启新的结构体
				proute = newroute(pkg)
			} else if regsummary.MatchString(txt) { // 处理@Summary 函数
				arr := regsummary.FindStringSubmatch(txt)
				proute.Comment = arr[1]
			} else if regrouter.MatchString(txt) { // 处理@Router /a/b/c [POST,DELETE]
				arr := regrouter.FindStringSubmatch(txt)
				proute.Path = arr[1]
				if strings.Contains(arr[2], "[") {
					arr[2] = strings.TrimPrefix(arr[2], "[")
					arr[2] = strings.TrimSuffix(arr[2], "]")
					arr[2] = strings.Trim(arr[2], " ")
				}
				if arr[2] != "" {
					proute.Method = strings.Split(arr[2], ",")
				}
			} else if regcomment.MatchString(txt) { //处理注释
				arr := regcomment.FindStringSubmatch(txt)
				if proute != nil && proute.Comment == "" {
					proute.Comment = arr[1]
				}
			} else {

			}
		}
		return err
	})

	return
}

// 为路径打分,mark priay for rule
func score(path string) int {
	var scorechar string = "{:"
	var ret int = 0
	for _, v := range scorechar {
		if strings.Contains(path, string(v)) {
			ret += 1
		}
	}
	return ret
}

var tplrouter string = `
//don't modify !!!!
// create at ${datetime}
// creeate by ${author}
//go:generate  codectl router -a ${author} -s . -d . -n ${routerfile}
package ${package}

import (
	"net/http"
	"github.com/turingdance/infra/restkit"
)
type Route struct {
	Package string
	Module  string
	Func    string
	Path    string
	Method  []string
	Comment string
	HandlerFunc http.HandlerFunc
}

var (
{{- range $k,$v := . }}
	{{$module := $v.Module|camel}}
	// {{$v.Comment}}
	{{$module}}Ctrl = &{{if ne $v.Package "${package}" }}{{$v.Package}}.{{end}}{{$v.Module}}{}
{{end}}
)
var MapCtrl map[string]any = map[string]any{
	{{- range $k,$v := . }}
	{{$module := $v.Module|camel}}
	// {{$v.Comment}}
		"{{$v.Module}}":{{$module}}Ctrl,
	{{end}}
}
var Routes []Route= []Route{
	{{- range $k,$v := . }}
	{{$module := $v.Module|camel}}
		{{- range $m,$n := $v.Children }}
		{Package:"{{$n.Package}}" ,Module:"{{$n.Module}}",HandlerFunc:{{$module}}Ctrl.{{$n.Func}},Func:"{{$n.Func}}",Path: "{{$n.Path}}",Method:[]string{	{{- range $x,$y := $n.Method }}"{{$y}}",{{end}} },Comment:"{{$n.Comment}}"},
		{{end}}
	{{end}}
}

var DefaultRouter *restkit.Router = restkit.NewRouter().PathPrefix("/")
// 初始化路由
func InitRouter(router *restkit.Router) {
	{{- range $k,$v := . }}
	{{$module := $v.Module|camel}}
	// {{$v.Comment}}
	{{$module}}Ctrl := &{{$v.Module}}{}
	{{$module}}router := router.Subrouter().PathPrefix("{{$v.Path}}")
	{{- range $g,$h := $v.Children }}
	//{{$h.Comment}}
	{{$module}}router.HandleFunc("{{$h.Path}}", {{$module}}Ctrl.{{$h.Func}}).Methods({{range $h.Method}}"{{.}}",{{end}})
	{{end}}
	{{end}}
}
func init() {
	InitRouter(DefaultRouter)
}
`

// replace keyword to ruled str
func replace(input string, rule map[string]string) string {
	for k, v := range rule {
		input = strings.ReplaceAll(input, k, v)
	}
	return input
}

type ModuleItem struct {
	Module  string
	Comment string
	Path    string
}

// 生成代码
func gencode(dirdst string, routes []*Route) (err error) {
	bts, err := os.ReadFile(tplfile)
	if err != nil {
		// 如果是文件不存在问题 则直接
		if os.IsNotExist(err) {
			os.WriteFile(tplfile, []byte(tplrouter), 0755)
		} else {
			return err
		}
	} else {
		tplrouter = string(bts)
	}
	//NodeRoute
	routetree := make([]*RouteTree, 0)
	// 这里是处理model
	for _, v := range routes {
		if v.Func == "" && v.Module != "" {
			routetree = append(routetree, &RouteTree{
				Route:    *v,
				Children: make([]*Route, 0),
			})
		}
	}

	for _, v := range routes {
		if v.Func != "" {
			for _, node := range routetree {
				if node.Module == v.Module {
					node.Children = append(node.Children, v)
				}
			}
		}
	}
	for _, v := range routetree {
		slicekit.SortStable(v.Children, func(e1, e2 *Route) bool {
			score1 := score(e1.Path)
			score2 := score(e2.Path)
			return score1 > score2
		})
	}
	dirdst, _ = filepath.Abs(dirdst)
	pkg := filepath.Base(dirdst)
	datatime := time.Now().Format("2006-01-02 15:04:05")
	tpl, err := template.New("root").Funcs(template.FuncMap{"join": func(str []string) string {
		return strings.Join(str, ",")
	},
		"json": func(input any) string {
			bts, _ := json.Marshal(input)
			return string(bts)
		},
		"camel": camel,
	}).Parse(replace(tplrouter, map[string]string{
		"${package}":    pkg,
		"${datetime}":   datatime,
		"${author}":     author,
		"${dirdst}":     dirdst,
		"${dirsrc}":     dirsrc,
		"${routerfile}": routerfile,
	}))
	if err != nil {
		return err
	}
	dstfilename := path.Join(dirdst, routerfile)
	_, err = os.Stat(dstfilename)
	// 如果文件不存在,则创建
	if err == nil || os.IsNotExist(err) {
		err = os.Rename(dstfilename, dstfilename+".bak")
	}
	dstfile, err := os.Create(dstfilename)
	if err != nil {
		return err
	}
	dstfile.Truncate(0)
	defer dstfile.Close()
	err = tpl.Execute(dstfile, routetree)
	return err

}

func gen(dirsrc string, dirdst string, excludes ...string) error {
	excludes = append(excludes, dstfile)
	routes, err := buildroutes(dirsrc, excludes...)

	if err != nil {
		return err
	}
	return gencode(dirdst, routes)
}

// 子命令定义 运行方法 go run main.go version 编译后 ./hugo version
var routerCmd = &cobra.Command{
	Use:   "router", // Use这里定义的就是命令的名称
	Short: "generate route by annotation",
	Long:  `generate route by annotation`,
	Run: func(cmd *cobra.Command, args []string) { //这里是命令的执行方法
		prj, _ := logic.TakeDefaultProject()
		if prj == nil {
			fmt.Println("当前参数尚未配置")
			return
		}
		if dirsrc != "" && dirdst == "" {
			dirdst = dirsrc
		}
		if author == "" {
			author = prj.Author
		}
		//扫描目录下的每一个文件
		if err := gen(dirsrc, dirdst, excludes...); err != nil {
			fmt.Println("gen route ", filepath.Join(dirsrc), "->", filepath.Join(dirdst, routerfile), err.Error())
		} else {
			fmt.Println("gen route ", filepath.Join(dirsrc), "->", filepath.Join(dirdst, routerfile), " ok")
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		//这个在命令执行前执行
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		//这个在命令执行后执行
	},
	// 还有其他钩子函数
}

func init() {
	rootCmd.AddCommand(routerCmd)
	routerCmd.Flags().StringVarP(&dirsrc, "src", "s", "", "dir of source")
	routerCmd.Flags().StringVarP(&dirdst, "dst", "d", "", "dir for save")
	routerCmd.Flags().StringVarP(&tplfile, "tpl", "t", "./router.tpl", "tpl for router")
	routerCmd.Flags().StringVarP(&author, "author", "a", "", "author of code")
	routerCmd.Flags().StringVarP(&routerfile, "name", "n", "router.gen.go", "name of router file")
	routerCmd.Flags().StringArrayVarP(&excludes, "exclude", "x", []string{}, "file will be exclude for scan...")
}

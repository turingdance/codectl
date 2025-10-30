package biz

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra" // 安装依赖 go get -u github.com/spf13/cobra/cobra
	"github.com/spf13/viper"
	"github.com/turingdance/codectl/app/conf"
	"github.com/turingdance/codectl/infra"
	"github.com/turingdance/infra/filekit"
	"github.com/turingdance/infra/logger"
	"github.com/turingdance/infra/oskit"
	"github.com/turingdance/infra/slicekit"
	"gopkg.in/yaml.v3"
)

type tplctrl struct {
	tpldir string
}

type TplInfo struct {
	Name    string `yaml:"name"`
	Memo    string `yaml:"memo"`
	File    string `yaml:"file"`
	Url     string `yaml:"url"`
	Lang    string `yaml:"lang"`
	Package string `yaml:"package"` // 原始包名称
}

/*
  - name: golang+vue3
    memo: golang+vue3+ts+tailwindcss+pina
    file: golang+vue3.zip
    url: https://github.com/turingdance/golang+vue3.git
*/
var tplsArr []TplInfo = []TplInfo{{
	Name:    "golang+vue3",
	Memo:    "golang+vue3+ts+tailwindcss+pina",
	File:    "golang+vue3.zip",
	Url:     "https://github.com/turingdance/golang-vue3.git",
	Lang:    "golang",
	Package: "turingdance.com/turing",
},
}

type TemplateDataStruct struct {
	Templates []TplInfo `yaml:"templates"`
}

func synctpl(tpls ...TplInfo) (result []TplInfo, err error) {
	// 文件存在
	viper.SetConfigFile(conf.TplDbFile)
	viper.SetConfigType("yaml")
	// 存在
	tmpdata := &TemplateDataStruct{
		Templates: make([]TplInfo, 0),
	}
	// 如果不存在,那么直接加载默认数据
	if !filekit.Exists(conf.TplDbFile) {
		tmpdata.Templates = append(tmpdata.Templates, tplsArr...)
	} else { // 如果存在,从配置文件加载
		viper.ReadInConfig()
		existdata := &TemplateDataStruct{
			Templates: make([]TplInfo, 0),
		}
		err = viper.Unmarshal(existdata)
		if err != nil {
			return
		}
		// 加载老数据
		tmpdata.Templates = append(tmpdata.Templates, existdata.Templates...)
	}

	// 合并新的数据
	tmpdata.Templates = append(tmpdata.Templates, tpls...)
	// 去重
	tmpdata.Templates = infra.Unique(tmpdata.Templates)
	// 保存到配置文件中
	yamlData, _err := yaml.Marshal(tmpdata)
	if _err != nil {
		return tpls, _err
	}
	if err = viper.ReadConfig(bytes.NewBuffer(yamlData)); err != nil {
		return
	}
	err = viper.WriteConfig()
	return tmpdata.Templates, err
}

func resettpl(tpls ...TplInfo) (result []TplInfo, err error) {
	// 文件存在
	viper.SetConfigFile(conf.TplDbFile)
	viper.SetConfigType("yaml")
	// 存在
	tmpdata := &TemplateDataStruct{
		Templates: make([]TplInfo, 0),
	}
	tmpdata.Templates = append(tmpdata.Templates, tplsArr...)
	tmpdata.Templates = append(tmpdata.Templates, tpls...)
	// 去重
	tmpdata.Templates = infra.Unique(tmpdata.Templates)
	// 保存到配置文件中
	yamlData, _err := yaml.Marshal(tmpdata)
	if _err != nil {
		return tpls, _err
	}
	if err = viper.ReadConfig(bytes.NewBuffer(yamlData)); err != nil {
		return
	}
	err = viper.WriteConfig()
	return tmpdata.Templates, err
}

func listtpl(tpls ...TplInfo) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"序号", "名称", "包名称", "适用语言", "描述", "地址"})
	for index, tpl := range tpls {
		table.Append([]string{
			strconv.Itoa(index + 1),
			tpl.Name,
			tpl.Package,
			tpl.Lang,
			tpl.Memo,
			tpl.Url,
		})
	}
	table.Render()
}

// 列表
func (s *tplctrl) list(args []string) error {
	tpls, err := synctpl()
	listtpl(tpls...)
	return err
}

// 添加
func (s *tplctrl) setupv1(args []string) (err error) {
	src := args[0]
	if strings.HasPrefix(src, "http://") || strings.HasPrefix(src, "https://") {
		return s.add_net(args)
	} else {
		return s.add_local(args)
	}
}

// 添加
func (s *tplctrl) add(args []string) (err error) {
	var tpl TplInfo
	for tpl.Name == "" {
		fmt.Println("请输入模板信息")
		fmt.Print("name>")
		fmt.Scanln(&tpl.Name)
	}
	for tpl.Memo == "" {
		fmt.Println("请输入模板描述,技术栈等")
		fmt.Print("memo>")
		fmt.Scanln(&tpl.Memo)
	}
	for tpl.Lang == "" {
		fmt.Println("请输入对应后端语言,golang/java/python等")
		fmt.Print("lang>")
		fmt.Scanln(&tpl.Lang)
	}
	for tpl.Package == "" {
		fmt.Println("请后端包名称,如 github.com/turingdance/codectl")
		fmt.Print("package>")
		fmt.Scanln(&tpl.Package)
	}

	for tpl.Url == "" {
		fmt.Println("请输入模板下载地址")
		fmt.Print("url>")
		fmt.Scanln(&tpl.Url)
	}
	// 添加到数据配置文件
	tpls, err := synctpl(tpl)
	if err != nil {
		return
	}
	listtpl(tpls...)
	return nil

}

func (s *tplctrl) add_local(args []string) (err error) {
	src := args[0]
	basedir := filepath.Base(src)
	dstdir := filepath.Join(s.tpldir, basedir)
	if len(args) > 1 {
		dstdir = filepath.Join(s.tpldir, args[1])
	}
	err = filekit.CopyDir(src, dstdir)
	if err != nil {
		return err
	}
	return s.list(args)
}

func (s *tplctrl) copy(args []string) (err error) {
	from := args[0]
	dst := args[1]
	dirfromabs := path.Join(conf.DirTpldata, from+"/")
	dirtoabs := path.Join(conf.DirTpldata, dst)
	err = filekit.CopyDir(dirfromabs, dirtoabs)
	if err == nil {
		fmt.Printf("copy ...✅\n%s \n=>\n %s \n", dirfromabs, dirtoabs)
	}
	return err
}

// 添加网络模版
func (s *tplctrl) add_net(args []string) (err error) {
	dstdir := ""
	if len(args) > 1 {
		dstdir = args[1]
	} else {
		urls, err := url.Parse(args[0])
		if err != nil {
			return err
		}
		paths := strings.Split(urls.Path, "/")
		dstname := paths[len(paths)-1]
		dstdir = strings.TrimSuffix(dstname, ".git")
	}
	dstdir = strings.ReplaceAll(dstdir, "/", "_")
	dstdir = strings.ReplaceAll(dstdir, "\\", "_")
	dstdir = path.Join(s.tpldir, dstdir)
	return s.clone([]string{args[0], dstdir})
}

// 添加
func (s *tplctrl) use(args []string) (err error) {
	tplId := args[0]
	viper.SetConfigType("yaml")
	viper.SetConfigFile(conf.ConfigFile)
	viper.Set("tplid", tplId)
	viper.WriteConfig()

	return err
}

// 添加 tpl del 1 2 3
func (s *tplctrl) del(args []string) error {
	tpls, _ := synctpl()
	ids := []int{}
	for _, id := range args {
		_id, _ := strconv.Atoi(id)
		ids = append(ids, _id-1)
	}
	tplnew := []TplInfo{}
	for i, v := range tpls {
		if !slicekit.Contains(ids, i) {
			tplnew = append(tplnew, v)
		}
	}
	list, err := resettpl(tplnew...)
	if err != nil {
		return err
	}
	listtpl(list...)
	return nil
}

// 添加
func (s *tplctrl) clone(args []string) error {
	argcmd := []string{
		"clone",
	}
	argcmd = append(argcmd, args...)
	resultch, errch, stopch := s.git(argcmd)
	for {
		select {
		case m := <-resultch:
			logger.Debug(m)
		case e := <-errch:
			logger.Debug(e)
		case e := <-stopch:
			return errors.New(e)
		}
	}
}
func (s *tplctrl) git(args []string) (resultch, errorch, stopch chan string) {
	ctx := context.Background()
	return oskit.ExecWithChanel(ctx, "git", args)
}
func NewtplCtl() *tplctrl {
	return &tplctrl{
		tpldir: conf.DirTpldata,
	}
}

var defaulttplctrl *tplctrl = NewtplCtl()
var tplmapfun map[string]func([]string) error = map[string]func([]string) error{
	"list": defaulttplctrl.list,
	"add":  defaulttplctrl.add,
	// "use":   defaulttplctrl.use,
	"del": defaulttplctrl.del,
	// "cp":    defaulttplctrl.copy,
	// "copy":  defaulttplctrl.copy,
	// "clone": defaulttplctrl.clone,
}

// 子命令定义 运行方法 go run main.go version 编译后 ./hugo version
var tplCmd = &cobra.Command{
	Use:   "tpl", // Use这里定义的就是命令的名称
	Short: "tpl manager",
	Long: `
tpl list
    list all template exist in tpl dir  
tpl del [tplid]
    delete tpl
tpl add tplpath [dstdir]
    if tplpath is a net url,download it and make it as a templete
	if tplpath is a local dir path like dir/of/tpl,copy it and make it as a templete
`,
	Run: func(cmd *cobra.Command, args []string) { //这里是命令的执行方法
		if len(args) == 0 {
			args = append(args, "list")
		} else if len(args) == 1 {
			args = append(args, "")
		}
		//args = append(args, "", "")
		cmdstr := args[0]

		paramargs := args[1:]
		if fun, ok := tplmapfun[cmdstr]; ok {
			if err := fun(paramargs); err != nil {
				logger.Error(err.Error())
			}
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
var (
	prjId int32 = 0
)

func init() {
	rootCmd.AddCommand(tplCmd)
}

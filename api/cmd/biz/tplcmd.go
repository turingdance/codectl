package biz

import (
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
	"github.com/turingdance/infra/filekit"
	"github.com/turingdance/infra/logger"
	"github.com/turingdance/infra/oskit"
)

type tplctrl struct {
	tpldir string
}

func (s *tplctrl) walkdir(filepath string) ([]string, error) {
	files, err := os.ReadDir(filepath) // files为当前目录下的所有文件名称【包括文件夹】
	if err != nil {
		return []string{}, err
	}
	var allfile []string
	for _, v := range files {
		allfile = append(allfile, v.Name())
	}
	return allfile, nil
}

// 列表
func (s *tplctrl) list(args []string) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"序号", "模板ID"})
	dirs, err := s.walkdir(s.tpldir)
	if err != nil {
		return err
	}
	total := len(dirs)
	for index, tplId := range dirs {
		table.Append([]string{
			strconv.Itoa(index + 1),
			tplId,
		})
	}
	fmt.Printf("tpldir = %s\ntotal = %d\n", s.tpldir, total)
	table.Render()
	return err
}

// 添加
func (s *tplctrl) setup(args []string) (err error) {
	src := args[0]
	if strings.HasPrefix(src, "http://") || strings.HasPrefix(src, "https://") {
		return s.add_net(args)
	} else {
		return s.add_local(args)
	}
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

// 添加
func (s *tplctrl) del(args []string) error {
	for _, dir := range args {
		rpath := filepath.Join(s.tpldir, dir)
		if filekit.IsExist(rpath) {
			os.RemoveAll(rpath)
		}
	}
	s.list(args)
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
	"list":  defaulttplctrl.list,
	"setup": defaulttplctrl.setup,
	"use":   defaulttplctrl.use,
	"del":   defaulttplctrl.del,
	"cp":    defaulttplctrl.copy,
	"copy":  defaulttplctrl.copy,
	"clone": defaulttplctrl.clone,
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
tpl setup tplpath [dstdir]
    if tplpath is a net url,download it and make it as a templete
	if tplpath is a local dir path like dir/of/tpl,copy it and make it as a templete
tpl clone [giturl] [dstdir]
    clone giturl  and save it at current dir for eg  tpl  clone https://github.com/techidea8/tpl-vue3-go.git
tpl cp tplname newtplname
tpl copy tplname newtplname
    copy tpl named tplname and save as newtplname
`,
	Run: func(cmd *cobra.Command, args []string) { //这里是命令的执行方法
		if len(args) == 0 {
			cmd.Help()
			return
		}
		if len(args) == 1 {
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

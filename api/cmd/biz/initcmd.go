package biz

import (
	"fmt"
	"time"

	"github.com/spf13/cobra" // 安装依赖 go get -u github.com/spf13/cobra/cobra
	"github.com/turingdance/codectl/app/conf"
	"github.com/turingdance/codectl/app/logic"
	"github.com/turingdance/codectl/app/model"
	"github.com/turingdance/infra/dbkit"
	"github.com/turingdance/infra/slicekit"
)

type initctrl struct{}

// 添加
func (s *initctrl) init(args []string) error {
	prj := &model.Project{}
	s.bindprj("please init project ", prj)
	if prj.Package == "" {
		prj.Package = conf.DefaultPackage
	}
	if prj.TplId == "" {
		prj.Package = conf.DefaultTplId
	}
	if prj.Name == "" {
		prj.Name = "biz"
	}
	if prj.Lang == "" {
		prj.Lang = conf.DefaultLang
	}
	prj.SortIndex = int32(time.Now().Unix())
	logic.Create(prj)
	return nil
}

func (s *initctrl) bindprj(title string, prj *model.Project) error {
	fmt.Println(title)
	fmt.Println("what's the name for project, eg: mall")
	fmt.Print("name >")
	fmt.Scanln(&prj.Name)

	fmt.Println("what's the title for project ,eg: 在线商城")
	fmt.Print("title >")
	fmt.Scanln(&prj.Title)

	fmt.Println("the dir where the project save eg: /path/save/dir ")
	fmt.Print("dirsave >")
	_dirsave := ""
	fmt.Scanln(&_dirsave)
	if _dirsave != "" {
		prj.Dirsave = _dirsave
	}

	fmt.Println("data source name for project  eg: mysql://user:password@host:port/dbmall?")
	fmt.Print("dsn >")
	_dsn := ""
	fmt.Scanln(&_dsn)

	if _dsn != "" {
		prj.Dsn = _dsn
	}
	if prj.Dsn != "" {
		info, err := dbkit.Parse(prj.Dsn)
		if err != nil {
			return err
		} else {
			prj.DbType = string(info.DbType)
			prj.DbName = info.Dbname
		}
	}

	fmt.Println("prefix of table ,such as sys_, or kf_>")
	fmt.Print("prefix>")
	var _prefix = ""
	fmt.Scanln(&_prefix)
	if _prefix != "" {
		prj.Prefix = _prefix
	}

	fmt.Println("package of app ,eg : turingdance.com/turing/app")
	fmt.Print("package>")
	var _package = ""
	fmt.Scanln(&_package)
	if _package != "" {
		prj.Package = _package
	}

	fmt.Println("which tpl for use,please ")
	defaulttplctrl.list([]string{})
	fmt.Print("tplId>")
	fmt.Scanln(&prj.TplId)

	fmt.Println("language eg:golang/java,default is golang")
	fmt.Print("lang>")
	fmt.Scanln(&prj.Lang)

	fmt.Println("are you confirm project use those info?Y/N")
	fmt.Print(">")
	var yesornot = "N"
	fmt.Scanln(&yesornot)
	if slicekit.Contains[string]([]string{"Y", "y", "YES", "yes", "Yes"}, yesornot) {
		//这里创建项目
		prj.SortIndex = int32(time.Now().Unix())
		logic.Create(prj)
		defaultprjctrl.list([]string{})
	} else {

	}
	return nil
}

// 子命令定义 运行方法 go run main.go version 编译后 ./hugo version
var initctrlCmd = &cobra.Command{
	Use:   "init", // Use这里定义的就是命令的名称
	Short: "init a project",
	Long: `
init 
    init a project
`,
	Run: func(cmd *cobra.Command, args []string) { //这里是命令的执行方法
		if len(args) < 1 {

		} else {
			args = append(args, "")
			projectmapfun[args[0]](args[1:])
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

	rootCmd.AddCommand(initctrlCmd)
}

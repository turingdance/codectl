package biz

import (
	"fmt"
	"time"

	"github.com/spf13/cobra" // 安装依赖 go get -u github.com/spf13/cobra/cobra
	"github.com/spf13/viper"
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
	viper.SetConfigType("yaml")
	viper.SetConfigFile(conf.ConfigFile)
	viper.ReadInConfig()
	viper.Unmarshal(prj)
	s.bindprj("now we are create a project ", prj)
	prj.SortIndex = int32(time.Now().Unix())
	logic.Create(prj)
	return nil
}

func (s *initctrl) bindprj(title string, prj *model.Project) error {
	fmt.Println(title)
	fmt.Println("what's the name for project, eg: mall,default:", prj.Name)
	fmt.Print("name>")
	fmt.Scanln(&prj.Name)

	fmt.Println("what's the title for project ,eg: 在线商城,default:", prj.Title)
	fmt.Print("title>")
	fmt.Scanln(&prj.Title)

	fmt.Println("the dir where the project save eg: /path/save/dir ,default:", prj.Dirsave)
	fmt.Print("saveas>")
	_dirsave := ""
	fmt.Scanln(&_dirsave)
	if _dirsave != "" {
		prj.Dirsave = _dirsave
	}
	fmt.Println("data source name for project  eg: mysql://user:password@host:port/dbmall?,default is ", prj.Dsn)
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

	fmt.Println("prefix of table ,such as sys_, or kf_,default is", prj.Prefix)
	fmt.Print("prefix>")
	var _prefix = ""
	fmt.Scanln(&_prefix)
	if _prefix != "" {
		prj.Prefix = _prefix
	}
	fmt.Println("package of app ,eg : turingdance.com/turing/app,default is ", prj.Package)
	fmt.Print("package>")
	var _package = ""
	fmt.Scanln(&_package)
	if _package != "" {
		prj.Package = _package
	}

	fmt.Println("which tpl for use, default is ", prj.TplId)
	defaulttplctrl.list([]string{})
	fmt.Print("tplId>")
	fmt.Scanln(&prj.TplId)

	fmt.Println("language eg:golang/java,default is ", prj.Lang)
	fmt.Print("lang>")
	fmt.Scanln(&prj.Lang)

	fmt.Println("are you sure to create project use those info?Y/N")
	fmt.Print("Y/N>")
	var yesornot = "N"
	fmt.Scanln(&yesornot)
	if slicekit.Contains[string]([]string{"Y", "y", "YES", "yes", "Yes"}, yesornot) {
		//这里创建项目
		prj.SortIndex = int32(time.Now().Unix())
		defaultprjctrl.list([]string{})
		//翻转
		reverse(prj)
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
		(&initctrl{}).init(args)
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

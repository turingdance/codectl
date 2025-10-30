package biz

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra" // 安装依赖 go get -u github.com/spf13/cobra/cobra
	"github.com/spf13/viper"
	"github.com/turingdance/codectl/app/conf"
	"github.com/turingdance/codectl/app/logic"
	"github.com/turingdance/infra/dbkit"
)

// 设置默认数据
type configctrl struct{}

func (s *configctrl) updateOne(key string, value string) error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(conf.ConfigFile)
	viper.Set(key, value)
	viper.WriteConfig()
	return nil
}
func (s *configctrl) update(kvpear map[string]string) (err error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(conf.ConfigFile)
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	for key, value := range kvpear {
		_k := strings.TrimSpace(key)
		_v := strings.TrimLeft(value, " ")
		viper.Set(_k, _v)
	}
	err = viper.WriteConfig()
	if err != nil {
		return
	}
	return err
}

func (s *configctrl) show() error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(conf.ConfigFile)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("error %v", err)
		return err
	}
	data := map[string]string{}
	err = viper.Unmarshal(&data)
	if err != nil {
		fmt.Printf("error %v", err)
		return err
	}

	for k, v := range data {
		fmt.Printf("%s=%s\n", k, v)
	}
	return nil
}
func (s *configctrl) listset() error {
	prj, err := logic.TakeDefaultProject()
	if err != nil {
		return err
	}
	fmt.Println("appname=", prj.Name)
	fmt.Println("auth   =", prj.Author)
	fmt.Println("tile   =", prj.Title)
	fmt.Println("dsn    =", prj.Dsn)
	fmt.Println("dirsave=", prj.Dirsave)
	fmt.Println("prefix =", prj.Prefix)
	fmt.Println("package=", prj.Package)
	return nil
}

func (s *configctrl) setonebyone() error {
	prj, err := logic.TakeDefaultProject()
	if err != nil {
		return err
	}
	var input = ""
	for {
		fmt.Println("appname:the name of app,must be english ,will use as service name ")
		fmt.Println("default: ", prj.Name)
		fmt.Print(">")
		fmt.Scanln(&input)
		if input != "" {
			prj.Name = input
			input = ""
			break
		}
		if prj.Name != "" {
			break
		}
	}

	for {
		fmt.Println("author: author of this app ")
		fmt.Println("default: ", prj.Author)
		fmt.Print(">")
		fmt.Scanln(&input)
		if input != "" {
			prj.Author = input
			input = ""
		}
		if prj.Author != "" {
			break
		}
	}
	for {

		fmt.Println("title: the title of project")
		fmt.Println("default: ", prj.Title)
		fmt.Print(">")
		fmt.Scanln(&input)
		if input != "" {
			prj.Title = input
			input = ""
			break
		}
		if prj.Title != "" {
			break
		}
	}

	for {
		fmt.Println("dsn: required, eg:appdata.db or mysql://user:password@tcp(host:port)/db?quest")
		fmt.Println("default: ", prj.Dsn)
		fmt.Print(">")
		fmt.Scanln(&input)
		if input != "" {
			prj.Dsn = input
			dbcfg, _ := dbkit.Parse(input)
			prj.DbName = dbcfg.Dbname
			prj.DbType = string(dbcfg.DbType)
			input = ""
			break
		}
		if prj.Dsn != "" {
			break
		}
	}

	fmt.Println("prefix:prefix of table,such as sys_, or kf_>")
	fmt.Print(">")
	fmt.Scanln(&input)
	if input != "" {
		prj.Prefix = input
		input = ""
	}

	fmt.Println("package:,eg turingdance.com/turing")
	fmt.Print(">")
	fmt.Scanln(&input)
	if input != "" {
		prj.Package = input
		input = ""
	} else {
		if prj.Package == "" {
			prj.Package = "turingdance.com/turing"
		}
	}
	logic.UpdateDefaultProject(prj)
	return nil
}

// 子命令定义 运行方法 go run main.go version 编译后 ./hugo version
var setctrlCmd = &cobra.Command{
	Use:   "set", // Use这里定义的就是命令的名称
	Short: "set default project params",
	Long: `
		set default project params one by ONE
	`,
	Run: func(cmd *cobra.Command, args []string) { //这里是命令的执行方法
		ctrl := &configctrl{}
		if lsset {
			ctrl.listset()
		} else {
			ctrl.setonebyone()
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

// 是否展示全部配置
var lsset = false

func init() {
	setctrlCmd.Flags().BoolVarP(&lsset, "ls", "l", false, "展示当前项目配置")
	rootCmd.AddCommand(setctrlCmd)
}

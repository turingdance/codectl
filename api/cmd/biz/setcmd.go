package biz

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra" // 安装依赖 go get -u github.com/spf13/cobra/cobra
	"github.com/spf13/viper"
	"github.com/turingdance/codectl/app/conf"
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

// 子命令定义 运行方法 go run main.go version 编译后 ./hugo version
var setctrlCmd = &cobra.Command{
	Use:   "set", // Use这里定义的就是命令的名称
	Short: "set default project params",
	Long: `
set 
    show all default value
set name1=value1 name2=value2
    set default value of project
`,
	Run: func(cmd *cobra.Command, args []string) { //这里是命令的执行方法
		ctrl := &configctrl{}
		if len(args) < 1 {
			fmt.Println("all params list as follow")
			ctrl.show()
		} else {
			tmp := map[string]string{}
			for _, v := range args {
				index := strings.Index(v, "=")
				tmp[v[:index]] = v[(index + 1):]
			}
			ctrl.update(tmp)
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
	rootCmd.AddCommand(setctrlCmd)
}

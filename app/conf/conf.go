package conf

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
	"github.com/turingdance/codectl/app/model"
	"github.com/turingdance/infra/filekit"
	"gopkg.in/yaml.v3"
)

type ENVDEF string

const PROD ENVDEF = "prod"
const DEV ENVDEF = "dev"

type AppConf struct {
	DbType   string
	Dsn      string
	Env      ENVDEF
	LogFile  string
	LogLevel string
	Prefix   string
}

const (
	DefaultTplId   = "turing-vue3-go-v2"
	DefaultPackage = "turngdance.com/turing/app"
	DefaultLang    = "golang"
)
const _dirappconfig = ".codectl"

var DirUserConfig, _ = os.UserConfigDir()
var (
	DirAppConfig = path.Join(DirUserConfig, _dirappconfig)
	DirTpldata   = path.Join(DirUserConfig, _dirappconfig, "tpldir")
	ConfigFile   = path.Join(DirUserConfig, _dirappconfig, "config.yaml")
)
var DefaultAppConf *AppConf = &AppConf{
	DbType:  "sqlite",
	Dsn:     path.Join(DirAppConfig, "codectl.db"),
	Env:     PROD,
	Prefix:  "gen_",
	LogFile: "gen.log",
}
var IdleProject *model.Project = &model.Project{
	Author:  "winlion@turingdance.com",
	Dirsave: "./",
	Name:    "project",
	TplId:   "tpl-vue3-go-v2",
	Dsn:     "./app.db",
	Lang:    "golang",
	Title:   "我的应用",
	Package: "turingdance.com/turing/app",
}

func init() {

	os.MkdirAll(DirTpldata, os.ModeDir)
	// 初始化默认的
	if !filekit.IsExist(ConfigFile) {
		viper.SetConfigType("yaml")
		byts, err := yaml.Marshal(IdleProject)
		if err != nil {
			fmt.Printf("读取配置失败: %v", err)
			return
		}
		if err := viper.ReadConfig(bytes.NewBuffer(byts)); err != nil {
			fmt.Printf("读取配置失败: %v", err)
			return
		}
		viper.WriteConfigAs(ConfigFile)
	}
}

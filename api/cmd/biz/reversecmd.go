package biz

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra" // 安装依赖 go get -u github.com/spf13/cobra/cobra
	"github.com/turingdance/codectl/app/conf"
	"github.com/turingdance/codectl/app/logic"
	"github.com/turingdance/codectl/app/model"
	"github.com/turingdance/infra/dbkit"
	"github.com/turingdance/infra/logger"
	"github.com/turingdance/infra/stringx"
)

type reversectrl struct {
	export []exportctrl
}

func NewReverseCtrl() *reversectrl {
	return &reversectrl{
		export: make([]exportctrl, 0),
	}
}
func (s *reversectrl) Init() {

}
func reverse(args []string) (err error) {
	// 配置文件展示加载的规则
	if rulefile != "" {
		conf.ResetMapperRule(rulefile)
	}
	//首先获得全部表格
	prj, err := logic.TakeCurrentProject()
	if err != nil {
		return err
	}
	if dirsave != "" {
		prj.Dirsave = dirsave
	}
	loglevel := logger.InfoLevel
	if env == string(conf.PROD) {
		loglevel = logger.ErrorLevel
	}
	if prj.ID == 0 {
		err = fmt.Errorf("缺少项目配置")
		return
	}
	// 导出的数据库
	exportdb, err := dbkit.OpenDb(dbkit.DBTYPE(prj.DbType), prj.Dsn, dbkit.WithWriter(os.Stdout), dbkit.SetLogLevel(loglevel))
	if err != nil {
		return err
	}
	tables, err := logic.BuildTableFromSchema(exportdb, prj.DbName)
	if err != nil {
		return err
	}
	for _, tb := range tables {
		if !strings.Contains(tb.Name, prj.Prefix) {
			logger.Debugf("ignore  %s because of %s", tb.Name, prj.Prefix)
			continue
		}
		// module name
		modulename := stringx.UnderlineToCamelCase(strings.TrimPrefix(tb.Name, prj.Prefix))

		vo := &logic.PrepareVo{
			Project:    prj,
			BizDbEngin: exportdb,
			TableName:  tb.Name,
			ModuleName: modulename,
			BizTitle:   tb.Title,
			Methods:    model.SimpleMethods,
		}
		table, err := logic.PrepareExportTable(vo)
		if err != nil {
			return err
		}
		err = logic.ExportTable(table, conf.DirTpldata)
		if err != nil {
			return err
		}
		fmt.Println("✅generate code " + table.Name + "->" + table.Module + "✓")
	}

	return nil
}

// 子命令定义 运行方法 go run main.go version 编译后 ./hugo version
var reverseCmd = &cobra.Command{
	Use:   "reverse", // Use这里定义的就是命令的名称
	Short: "reverse all table of project to code",
	Long:  `reverse all table of project to code such as golang/java..`,
	Run: func(cmd *cobra.Command, args []string) { //这里是命令的执行方法
		if err := reverse(args); err != nil {
			logger.Error(err.Error())
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		//这个在命令执行前执行
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		//这个在命令执行后执行
		defaultexportctrl.Init()
	},
	// 还有其他钩子函数
}

func init() {
	reverseCmd.Flags().StringVarP(&rulefile, "rule", "r", "", "rule config file,like mapper.yml")
	reverseCmd.Flags().StringVarP(&dirsave, "dirsave", "d", "", "dir for save")
	rootCmd.AddCommand(reverseCmd)
}

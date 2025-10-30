package biz

import (
	"errors"
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

// 通配符匹配函数（递归版）
// pattern: 包含通配符的模式串（支持 * 和 ?）
// s: 目标字符串
func isMatchRecursive(pattern, s string) bool {
	// 递归终止条件：模式串为空时，目标串也必须为空才匹配
	if pattern == "" {
		return s == ""
	}

	// 第一个字符是否匹配（普通字符相等 或 模式串为 ?）
	firstMatch := s != "" && (pattern[0] == s[0] || pattern[0] == '?')

	// 处理 * 通配符：* 可以匹配 0 个或多个字符
	if pattern[0] == '*' {
		// 分支1：* 匹配 0 个字符，直接跳过 * 继续匹配后续
		// 分支2：* 匹配 1 个或多个字符，继续用 * 匹配目标串的下一个字符
		return isMatchRecursive(pattern[1:], s) || (s != "" && isMatchRecursive(pattern, s[1:]))
	}

	// 非 * 情况：当前字符匹配后，继续匹配后续字符
	return firstMatch && isMatchRecursive(pattern[1:], s[1:])
}

func reverse(prj *model.Project, _ ...string) (err error) {
	// 配置文件展示加载的规则
	if rulefile != "" {
		conf.ResetMapperRule(rulefile)
	}

	if dirsave != "" {
		prj.Dirsave = dirsave
	}
	loglevel := logger.InfoLevel
	if env == string(conf.PROD) {
		loglevel = logger.ErrorLevel
	}
	if prj.Dsn == "" {
		err = fmt.Errorf("缺少项目配置")
		return
	}
	// 导出的数据库
	exportdb, err := dbkit.OpenDb(prj.Dsn, dbkit.WithWriter(os.Stdout), dbkit.SetLogLevel(loglevel))
	if err != nil {
		return err
	}

	dbinfo, err := prj.DbInfo()
	if err != nil {
		return err
	}
	tables, err := logic.BuildTableFromSchema(exportdb, dbinfo.Dbname)
	if err != nil {
		return err
	}
	if len(tables) == 0 {
		return errors.New("there is found  no table")
	}
	logger.Info("generate code table num = %d", len(tables))
	for _, tb := range tables {
		if !strings.Contains(tb.Name, prj.Prefix) {
			logger.Infof("ignore  %s because of has no prefix %s", tb.Name, prj.Prefix)
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
		tpldir := prj.TplId // path.Join(conf.DirTpldata, prj.TplId)
		err = logic.ExportTable(table, tpldir)
		if err != nil {
			return err
		}
		fmt.Println("generate code " + table.Name + "->" + table.Module + "✓")
	}

	return nil
}

// 子命令定义 运行方法 go run main.go version 编译后 ./hugo version
var reverseCmd = &cobra.Command{
	Use:   "reverse", // Use这里定义的就是命令的名称
	Short: "reverse all table of project to code",
	Long:  `reverse all table of project to code such as golang/java..`,
	Run: func(cmd *cobra.Command, args []string) { //这里是命令的执行方法

		if prj, err := logic.TakeDefaultProject(); err != nil {
			fmt.Println(err.Error())
		} else {
			if dirsave != "" {
				prj.Dirsave = dirsave
			}
			if appname != "" {
				prj.Name = appname
			}
			if tablepatern != "" {
				prj.Prefix = tablepatern
			}
			if dstpkg != "" {
				prj.Package = dstpkg
			}
			if dsn != "" {
				prj.Dsn = dsn
			}

			if err := reverse(prj, args...); err != nil {
				fmt.Println(err.Error())
			}
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
	reverseCmd.Flags().StringVarP(&rulefile, "mapper", "r", "", `字段映射规则文件,如 mysql-golang.yml
	mysql-golang:
		TINYINT: int8
		SMALLINT: int16`)
	reverseCmd.Flags().StringVarP(&dirsave, "dirsave", "d", "", "生成的代码存储在哪个陌路如:app/path/to")
	reverseCmd.Flags().StringVarP(&appname, "appname", "n", "", "生成应用名字,请使用英文,如:cms")
	reverseCmd.Flags().StringVarP(&tablepatern, "table", "t", "", "表名称,支持通配符*,?,如 cms_*")
	reverseCmd.Flags().StringVarP(&trimprefix, "trimprefix", "x", "", "去掉表结构前缀,多个用,隔开.如cms_,rpt_")
	reverseCmd.Flags().StringVarP(&dsn, "dsn", "s", "", "数据库链接字符串 mysql://username:password@tcp(127.0.0.1:3306)/dbtestt?")
	reverseCmd.Flags().StringVarP(&dstpkg, "dstpkg", "g", "", "生成系统的包名称")
	rootCmd.AddCommand(reverseCmd)
}

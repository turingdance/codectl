package biz

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra" // 安装依赖 go get -u github.com/spf13/cobra/cobra
	"github.com/turingdance/codectl/app/logic"
	"github.com/turingdance/codectl/infra"
	"github.com/turingdance/infra/filekit"
)

// 子命令定义 运行方法 go run main.go version 编译后 ./hugo version
var initctrlCmd = &cobra.Command{
	Use:   "init", // Use这里定义的就是命令的名称
	Short: "init a project",
	Long: `
	init [appname]
		init a  project  named appname
`,
	Run: func(cmd *cobra.Command, args []string) { //这里是命令的执行方法
		if len(args) > 0 {
			destAppName = args[0]
		}
		fmt.Println("选择应用框架,eg:1 2 ")
		tpls, err := synctpl()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		listtpl(tpls...)
		var choseIndex int = 0
		fmt.Print("$")
		fmt.Scanln(&choseIndex)
		if choseIndex < 1 || choseIndex > len(tpls) {
			fmt.Println("应用框架选择错误")
		}
		tpl := tpls[choseIndex-1]
		fmt.Printf("已选择框架%s,开始下载...\n", tpl.Name)
		prj, err := logic.TakeDefaultProject()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		absDir, _ := filepath.Abs(destDirforapp)
		prj.Name = destAppName
		prj.Dirsave = filepath.Join(absDir, destAppName)
		if filekit.Exists(prj.Dirsave) {
			if force {
				os.RemoveAll(prj.Dirsave)
			} else {
				fmt.Println(prj.Dirsave, "该目录已经存在", "请删除后再试")
				return
			}

		}
		prj.TplId = filepath.Join(absDir, destAppName, "server", "tpl")
		prj.Lang = tpl.Lang
		logic.UpdateDefaultProject(prj)

		infra.Clone(tpl.Url, prj.Dirsave)
		// 移除
		os.RemoveAll(filepath.Join(absDir, destAppName, ".git"))
		// 替换到新的包
		if destPkgName != "" {
			// 如果是golang
			serverDir := filepath.Join(absDir, destAppName, "server")
			err := filepath.Walk(serverDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					return infra.ReplaceInFile(path, tpl.Package, destPkgName)
				}
				return nil
			})

			if err != nil {
				fmt.Printf("Error walking the path: %s\n", err)
				os.Exit(1)
			}
			fmt.Println("String replacement completed successfully.")
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
var destDirforapp = "./"
var destAppName = "turingapp"
var destPkgName = ""
var force = false

func init() {
	initctrlCmd.Flags().StringVarP(&destDirforapp, "dirsave", "d", "./", "生成的代码存储在哪个目录")
	initctrlCmd.Flags().StringVarP(&destPkgName, "package", "k", "", "新的包名称")
	initctrlCmd.Flags().BoolVarP(&force, "force", "f", false, "是否强制删除已经存在的项目")
	// initctrlCmd.Flags().StringVarP(&destAppName, "appname", "n", "codectlapp", "当前应用名称")
	rootCmd.AddCommand(initctrlCmd)
}

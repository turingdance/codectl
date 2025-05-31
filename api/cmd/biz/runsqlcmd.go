package biz

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra" // 安装依赖 go get -u github.com/spf13/cobra/cobra
	"github.com/turingdance/codectl/app/logic"
	"github.com/turingdance/infra/dbkit"
	"github.com/turingdance/infra/logger"
	"gorm.io/gorm"
)

type sqlctrl struct {
}

func (s *sqlctrl) run(args []string) error {

	return nil
}

// splitSQLStatements 将SQL文件内容分割成单个SQL语句
func splitSQLStatements(content string) []string {
	var statements []string
	var currentStatement strings.Builder
	inQuotes := false
	quoteChar := rune(0)
	escaped := false
	commentMode := false
	commentStart := 0

	for i, char := range content {
		// 处理注释
		if !inQuotes {
			if !commentMode && i < len(content)-1 {
				// 检查是否开始单行注释
				if char == '-' && content[i+1] == '-' {
					commentMode = true
					commentStart = i
					continue
				}
				// 检查是否开始多行注释
				if char == '/' && content[i+1] == '*' {
					commentMode = true
					commentStart = i
					continue
				}
			}

			if commentMode {
				// 检查是否结束单行注释
				if char == '\n' && (content[commentStart] == '-' || content[commentStart] == '#') {
					commentMode = false
					continue
				}
				// 检查是否结束多行注释
				if char == '*' && i < len(content)-1 && content[i+1] == '/' {
					commentMode = false
					continue
				}
				continue
			}
		}

		// 处理转义字符
		if escaped {
			currentStatement.WriteRune(char)
			escaped = false
			continue
		}

		// 检查是否为转义字符
		if char == '\\' {
			escaped = true
			currentStatement.WriteRune(char)
			continue
		}

		// 处理引号
		if char == '\'' || char == '"' || char == '`' {
			if !inQuotes {
				// 开始引号
				inQuotes = true
				quoteChar = char
			} else if char == quoteChar {
				// 结束引号
				inQuotes = false
			}
			currentStatement.WriteRune(char)
			continue
		}

		// 处理语句分隔符
		if char == ';' && !inQuotes {
			statements = append(statements, currentStatement.String())
			currentStatement.Reset()
			continue
		}

		// 其他字符
		currentStatement.WriteRune(char)
	}

	// 添加最后一个语句
	if stmt := currentStatement.String(); strings.TrimSpace(stmt) != "" {
		statements = append(statements, stmt)
	}

	return statements
}

// executeInTransaction 在事务中执行SQL语句
func executeInTransaction(db *gorm.DB, statements []string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for i, stmt := range statements {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}

			fmt.Printf("执行语句 %d: %s...\n", i+1, truncateString(stmt, 50))
			result := tx.Exec(stmt)
			if result.Error != nil {
				return fmt.Errorf("执行语句 %d 失败: %v\n语句内容: %s", i+1, result.Error, stmt)
			}
			fmt.Printf("语句 %d 执行成功\n", i+1)
		}
		return nil
	})
}

// executeStatements 执行SQL语句，不使用事务
func executeStatements(db *gorm.DB, statements []string) error {
	for i, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}

		fmt.Printf("执行语句 %d: %s...\n", i+1, truncateString(stmt, 50))
		result := db.Exec(stmt)
		if result.Error != nil {
			return fmt.Errorf("执行语句 %d 失败: %v\n语句内容: %s", i+1, result.Error, stmt)
		}
		fmt.Printf("语句 %d 执行成功\n", i+1)
	}
	return nil
}

// truncateString 截断字符串并添加省略号
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
func _runfile(db *gorm.DB, sqlfile string, transaction bool) (err error) {
	sqlContent, err := os.ReadFile(sqlfile)
	if err != nil {
		return err
	}

	sqlStatements := splitSQLStatements(string(sqlContent))

	// 执行SQL语句
	if transaction {
		err = executeInTransaction(db, sqlStatements)
	} else {
		err = executeStatements(db, sqlStatements)
	}
	return
}
func runfile(db *gorm.DB, sqlfile string, transaction bool) {
	if err := _runfile(db, sqlfile, transaction); err != nil {
		logger.Errorf("run %s error=%v ×", sqlfile, err)
	} else {
		logger.Infof("run %s, success √", sqlfile)
	}
}

// 子命令定义 运行方法 go run main.go version 编译后 ./hugo version
var runsqlCmd = &cobra.Command{
	Use:   "runsql", // Use这里定义的就是命令的名称
	Short: "sql runner",
	Long: `
runsql 
	runsql --dir ./ it will walkdir and run a sql file
	runsql --file ./pathtofile.sql it will pathtofile.sql 
`,
	Run: func(cmd *cobra.Command, args []string) { //这里是命令的执行方法
		prj, err := logic.TakeDefaultProject()
		if err != nil {
			logger.Error(err.Error())
			return
		}
		db, err := dbkit.OpenDb(prj.Dsn, dbkit.Debug(debug))
		if err != nil {
			logger.Error(err.Error())
			return
		}
		if sqlfile != "" {
			runfile(db, sqlfile, transaction)
		} else if filedir != "" {
			filepath.WalkDir(filedir, func(path string, d fs.DirEntry, err error) error {
				if d.IsDir() {
					return nil
				}
				if strings.HasSuffix(path, ".sql") {
					runfile(db, path, transaction)
				}
				return nil
			})
		} else {
			cmd.Help()
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
var filedir string = "."
var sqlfile string = ""
var transaction bool = false
var debug bool = false

func init() {
	runsqlCmd.Flags().StringVar(&filedir, "dir", "", "dir  it will walk and find all *.sql file")
	runsqlCmd.Flags().StringVar(&sqlfile, "file", "", "file it will run")
	runsqlCmd.Flags().BoolVar(&debug, "debug", false, "debug of not")
	runsqlCmd.Flags().BoolVar(&transaction, "transaction", true, "use transaction or not")
	rootCmd.AddCommand(runsqlCmd)
}

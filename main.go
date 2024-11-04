package main

import (
	// 安装依赖 go get -u github.com/spf13/cobra/cobra

	"github.com/turingdance/codectl/api/cmd/biz"
	"github.com/turingdance/codectl/api/cmd/tpl"
)

// 入口函数
func main() {
	tpl.Release()
	biz.Execute()
}

package infra

import (
	"context"
	"errors"
	"fmt"

	"github.com/turingdance/infra/oskit"
)

// 添加
func Clone(args ...string) error {
	argcmd := []string{
		"clone",
	}
	argcmd = append(argcmd, args...)
	resultch, errch, stopch := Git(argcmd)
	for {
		select {
		case m := <-resultch:
			fmt.Println(m)
		case e := <-errch:
			fmt.Println(e)
		case e := <-stopch:
			return errors.New(e)
		}
	}
}
func Git(args []string) (resultch, errorch, stopch chan string) {
	ctx := context.Background()
	return oskit.ExecWithChanel(ctx, "git", args)
}

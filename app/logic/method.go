package logic

import (
	"github.com/turingdance/codectl/app/model"
	"github.com/turingdance/infra/slicekit"
)

func BuildMethod(methodArr []string) []model.Method {
	tmp := make([]model.Method, 0)
	for _, method := range model.AllSuportMethods {
		tmp = append(tmp, model.Method{
			Name:   method.Name,
			Title:  method.Title,
			Enable: slicekit.Contains(methodArr, method.Name),
		})
	}
	return tmp

}

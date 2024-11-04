package biz

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"regexp"
	"testing"
)

var rule1 *regexp.Regexp = regexp.MustCompile(`//\s+((?:\,?(?:post|get|put|delete|options))+)\s+((?:/\w+)+)`)
var regmap map[string]*regexp.Regexp = map[string]*regexp.Regexp{
	"// post /acc/create": rule1,
	"// gen by codectl ,donot modify ,https://github.com/turingdance/codectl.git": rule1,
	"post /acc/create":            rule1,
	"// post,get /acc/create":     rule1,
	"// post,get,put /acc/create": rule1,
}

func String(a any) string {
	s, _ := json.Marshal(a)
	return string(s)
}
func TestReg(t *testing.T) {
	for k, v := range regmap {
		arr := v.FindStringSubmatch(k)
		fmt.Println(k, String(arr))
		//t.Log(arr)
	}
}

func Test02(t *testing.T) {
	str1 := "// post,get,put /acc/create"
	patern := regexp.MustCompile(`//\s+(post|get|put[\,post|\,get|\,put]*)\s+[\/\w]+`)
	result := patern.FindStringSubmatch(str1)
	fmt.Println(str1, len(result), result)
}

func Test03(t *testing.T) {
	pt := regexp.MustCompile(`//\s+((?:post|get|put)\S*)\s+((?:/?\w+)+)`)
	arr := []string{
		"// post,get,put /a/b",
		"// post,get /c/d",
		"// post /e/f",
		"// get /h/j",
	}
	for _, v := range arr {
		aa := pt.FindStringSubmatch(v)
		fmt.Println(v, aa)
	}
}

func Test004(t *testing.T) {
	dst := "."
	t2, _ := filepath.Abs(dst)
	fmt.Println(filepath.Base(t2))

}

func Test005(t *testing.T) {
	dst := "./api/test/0034/"
	t2, _ := filepath.Abs(dst)
	fmt.Println(filepath.Base(t2))

}

func Test006(t *testing.T) {
	dst := "../api/test/0034"
	t2, _ := filepath.Abs(dst)
	fmt.Println(filepath.Base(t2))

}

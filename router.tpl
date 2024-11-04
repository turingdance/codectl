//don't modify !!!!
// create at ${datetime}
// creeate by ${author}
//go:generate  codectl router -a ${author} -s . -d . -n ${routerfile}
package ${package}

import (
	"net/http"
	"github.com/turingdance/infra/restkit"
)
type Route struct {
	Package string
	Module  string
	Func    string
	Path    string
	Method  []string
	Comment string
	HandlerFunc http.HandlerFunc
}
var (
{{- range $k,$v := . }}
	{{$module := $v.Module|camel}}
	// {{$v.Comment}}
	{{$module}}Ctrl = &{{if ne $v.Package "${package}" }}{{$v.Package}}.{{end}}{{$v.Module}}{}
{{end}}
)
var Routes []Route= []Route{
	{{- range $k,$v := . }}
	{{$module := $v.Module|camel}}
		{{- range $m,$n := $v.Children }}
		Route{Package:"{{$n.Package}}" ,Module:"{{$n.Module}}",HandlerFunc:{{$module}}Ctrl.{{$n.Func}},Func:"{{$n.Func}}",Path: "{{$n.Path}}",Method:[]string{	{{- range $x,$y := $n.Method }}"{{$y}}",{{end}} },Comment:"{{$n.Comment}}"},
		{{end}}
	{{end}}
}
package biz

//环境变量
var env string
var configfile string
var rulefile string = ""
var dirsave string = ""

//-source .  -dstpkg turingdance.com/turing/devfast/api/rest/sys -dstfile routefunc.go
// 需要解析router 的
var source string = "."

//需要的包路径
var dstpkg string = "turingdance.com/turing/app"

// 生成的router 文件

var dstfile = "router_file.go"

// 数据库的reverse
var dsn = ""

// 去掉前缀
var trimprefix = ""

// 生成的appname
var appname = ""

//正则表达式匹配的表名称
var tablepatern = ".*"

// 是否显示version
var showversion = false
var VERSION = "1.0.0.20251029"

type ReverseRequest struct {
	Dsn         string
	Trimprefix  string
	AppName     string
	TablePatern string
	Dirsave     string
	Rulefile    string
}

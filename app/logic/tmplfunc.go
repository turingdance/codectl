package logic

import (
	"html"
	"html/template"
	"strings"

	"github.com/turingdance/infra/slicekit"
	"github.com/turingdance/infra/stringx"
)

// type: MySQL 数据类型字符串（如 "varchar", "int", "datetime" 等）
// 返回值: 对应的 Go 类型字符串（如 "string", "int", "time.Time" 等）
func mysqltogorm(typ string) string {
	// 预处理：去除空格、转为小写，提取主类型（忽略长度/精度信息，如 "varchar(255)" → "varchar"）
	typ = strings.TrimSpace(typ)
	typ = strings.ToLower(typ)

	// 提取主类型（分割掉括号内的长度/精度信息）
	idx := strings.Index(typ, "(")
	if idx != -1 {
		typ = typ[:idx]
	}

	// 处理无符号类型（如 "int unsigned" → 主类型为 "int"，标记为无符号）
	isUnsigned := strings.Contains(typ, "unsigned")
	typ = strings.ReplaceAll(typ, "unsigned", "")
	typ = strings.TrimSpace(typ)

	// 根据 MySQL 主类型映射 Go 类型
	switch typ {
	// 整数类型
	case "tinyint":
		if isUnsigned {
			return "int"
		}
		return "int"
	case "smallint":
		if isUnsigned {
			return "uint"
		}
		return "int"
	case "mediumint":
		if isUnsigned {
			return "int" // mediumint 范围接近 uint32 子集
		}
		return "int"
	case "int", "integer":
		if isUnsigned {
			return "uint"
		}
		return "int"
	case "bigint":
		if isUnsigned {
			return "uint"
		}
		return "bigint"

	// 浮点/小数类型
	case "float":
		return "float"
	case "double", "real":
		return "float"
	case "decimal", "numeric":
		// 高精度小数推荐使用 decimal 库类型
		return "float" // 需要引入 github.com/shopspring/decimal

	// 字符串类型
	case "char", "varchar", "text", "mediumtext", "longtext":
		return "string"
	case "enum", "set":
		return "string" // 枚举/集合用字符串存储

	// 二进制类型
	case "binary", "varbinary", "blob", "mediumblob", "longblob":
		return "bytes"

	// 时间类型
	case "date", "time", "datetime", "timestamp":
		return "time"

	// 布尔类型（MySQL 无原生 bool，用 tinyint(1) 模拟）
	case "bool":
		return "bool"

	// JSON 类型（MySQL 5.7+）
	case "json":
		return "json.RawMessage" // 推荐用 json.RawMessage 处理 JSON 数据

	// 特殊类型
	case "uuid":
		return "string" // UUID 以字符串形式存储
	case "bit":
		return "bytes" // 位类型用字节数组处理（长度 1 时可考虑 bool）

	// 未匹配的类型
	default:
		return ""
	}
}

// gorm4postgresqltype 输入 PostgreSQL 原始数据类型，返回对应的 GORM（Go）类型
// typ: PostgreSQL 数据类型字符串（如 "varchar", "integer", "jsonb" 等）
// 返回值: 对应的 Go 类型字符串（如 "string", "int", "json.RawMessage" 等）
func postgresqltogrom(typ string) string {
	// 预处理：去除空格、转为小写，提取主类型（忽略长度/精度/修饰符）
	typ = strings.TrimSpace(typ)
	typ = strings.ToLower(typ)

	// 处理数组类型（如 "integer[]" → 提取 "integer" 并标记为数组）
	isArray := strings.HasSuffix(typ, "[]")
	if isArray {
		typ = strings.TrimSuffix(typ, "[]")
		typ = strings.TrimSpace(typ)
	}

	// 提取主类型（忽略长度/精度，如 "varchar(255)" → "varchar"）
	idx := strings.Index(typ, "(")
	if idx != -1 {
		typ = typ[:idx]
	}

	// 基础类型映射
	baseType := ""
	switch typ {
	// 整数类型
	case "smallint", "int2":
		baseType = "int"
	case "integer", "int", "int4":
		baseType = "int"
	case "bigint", "int8":
		baseType = "int"
	case "smallserial", "serial2": // 自增小整数
		baseType = "int"
	case "serial", "serial4": // 自增整数
		baseType = "int"
	case "bigserial", "serial8": // 自增大整数
		baseType = "uint"

	// 浮点/小数类型
	case "real", "float4":
		baseType = "float"
	case "double precision", "float8":
		baseType = "float"
	case "numeric", "decimal": // 高精度小数
		baseType = "float" // 依赖 github.com/shopspring/decimal

	// 字符串类型
	case "char", "character", "varchar", "character varying", "text":
		baseType = "string"
	case "citext": // 大小写不敏感字符串
		baseType = "string"

	// 二进制类型
	case "bytea":
		baseType = "[]byte"

	// 布尔类型
	case "boolean", "bool":
		baseType = "bool"

	// 时间类型
	case "date":
		baseType = "time"
	case "time", "time without time zone":
		baseType = "time"
	case "timestamp", "timestamp without time zone":
		baseType = "time"
	case "timestamp with time zone", "timestamptz":
		baseType = "time"

	// JSON 类型
	case "json":
		baseType = "json.RawMessage"
	case "jsonb": // 二进制 JSON（推荐）
		baseType = "json.RawMessage"

	// 特有类型
	case "uuid": // UUID 类型
		baseType = "string" // 或使用 gorm.UUID
	case "hstore": // 键值对类型
		baseType = "map[string]string" // 简单映射，复杂场景需自定义类型
	case "inet": // IP 地址类型
		baseType = "string" // 以字符串存储 IP 地址
	case "macaddr": // MAC 地址类型
		baseType = "string"
	case "bit": // 位类型
		baseType = "bytes"
	case "bit varying", "varbit": // 可变长度位类型
		baseType = "bytes"

	// 未匹配的类型
	default:
		return ""
	}

	// 处理数组类型（如 integer[] → []int）
	if isArray {
		return "[]" + baseType
	}

	return baseType
}
func unescape(s string) template.HTML {
	// 先解码 HTML 实体（&amp;→&）
	decoded := html.UnescapeString(s)
	// 返回 template.HTML 类型，告诉模板该内容已安全，无需再次转义
	return template.HTML(decoded)
}

var funcMaps = template.FuncMap{
	"ucfirst":          stringx.Ucfirst,
	"lcfirst":          stringx.Lcfirst,
	"jsstr":            stringx.JSStr,
	"jstxt":            stringx.JSStr,
	"js":               stringx.JS,
	"lower":            stringx.Lower,
	"upper":            stringx.Upper,
	"upercamel":        stringx.UnderlineToUperCamelCase,
	"camel":            stringx.UnderlineToCamelCase,
	"contains":         strings.Contains,
	"has":              slicekit.HasSubStr,
	"unescape":         unescape,
	"replaceall":       strings.ReplaceAll,
	"mysqltogorm":      mysqltogorm,
	"postgresqltogrom": postgresqltogrom,
}

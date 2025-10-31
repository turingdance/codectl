package logic

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/turingdance/codectl/app/conf"
	"github.com/turingdance/codectl/app/model"
	"github.com/turingdance/infra/cond"
	"github.com/turingdance/infra/dbkit"
	"github.com/turingdance/infra/filekit"
	"github.com/turingdance/infra/logger"
	"github.com/turingdance/infra/slicekit"
	"github.com/turingdance/infra/timekit"
	"gorm.io/gorm"
)

type ExportVO struct {
	Name    string
	Title   string
	Module  string
	Method  []model.Method
	Columns []model.Column
	Package string
	Project model.Project
	TplId   string
	Primary model.Column //主键信息
	Table   *model.Table
	Types   []string
}

type PrepareVo struct {
	Project    *model.Project //项目信息
	BizDbEngin *gorm.DB       //数据库连接信息
	TableName  string         //表结构信息
	ModuleName string         //模块名称
	BizTitle   string         //业务名称
	Methods    []string       //简单的方法
}
type CallbackFunc func(file string)

// 准备导出
func PrepareExportTable(vo *PrepareVo) (table *model.Table, err error) {
	prj := vo.Project
	dbengin := vo.BizDbEngin
	methods := vo.Methods
	table = &model.Table{
		ProjectID: prj.ID,
		Name:      vo.TableName,
	}
	dbconf, err := dbkit.ParseMysql(prj.Dsn)
	if err != nil {
		return nil, err
	}
	dbname := dbconf.Dbname
	prj.DbName = dbconf.Dbname
	prj.DbType = string(dbconf.DbType)
	// 取数据,有没有现成的
	table, err = Take(table, *cond.NewCondWrapper())
	// 如果报错了 直接报错
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	tableexist := (table != nil && table.ID > 0)
	// 如果已经存在了,对数据进行添加和修改
	mapexistcolumn := make(map[string]model.Column, 0)
	if tableexist {
		// 获取clomun
		// 首先获得字段和row之间的map
		fields, _, err := ListAllColumnByTableId(table.ID)
		if err != nil {
			return table, err
		}
		for _, field := range fields {
			mapexistcolumn[field.RawData.ColumnName] = field
		}
	} else { // 如果表不存在,那么创建表结构
		table.Module = vo.ModuleName
		table.Name = vo.TableName
		table.Title = vo.BizTitle
		table.ProjectID = prj.ID
		table.Method = BuildMethod(methods)
		table, err = Create(table)
		if err != nil {
			return
		}
	}
	// 新的数据列表,如果数据有变动则更新
	newcolumns, err := BuildColumnFromSchema(dbengin, dbname, vo.TableName)
	if err != nil {
		return
	}
	mapnewcolumn := make(map[string]model.Column, 0)
	keydbtypetolang := fmt.Sprintf("%s-%s", prj.DbType, prj.Lang)
	types := make([]string, 0)
	for _, newfield := range newcolumns {
		// 如果不存存在,说明作为添加存在
		if oldfield, exist := mapexistcolumn[newfield.RawData.ColumnName]; !exist {
			newfield.RawData.BuildColumn(&newfield)
			// 新的数据类型
			datatype := conf.DataTypeTOLangMapperRule[keydbtypetolang][newfield.RawData.DataType]
			if datatype == "" {
				datatype = newfield.RawData.DataType
			}
			newfield.DataType = datatype
			mapnewcolumn[newfield.RawData.ColumnName] = newfield
			types = append(types, newfield.DataType)
		} else {
			oldfield.RawData = newfield.RawData
			oldfield.RawData.BuildColumn(&oldfield)
			// 获得实际数据类型映射
			datatype := conf.DataTypeTOLangMapperRule[keydbtypetolang][oldfield.RawData.DataType]
			if datatype == "" {
				datatype = newfield.RawData.DataType
			}
			oldfield.DataType = datatype
			mapexistcolumn[newfield.RawData.ColumnName] = oldfield
			types = append(types, oldfield.DataType)
		}
	}
	tmpcreate := make([]model.Column, 0)
	tmpupdate := make([]model.Column, 0)
	// 继续保存到数据
	for _, field := range mapnewcolumn {
		field.TableID = table.ID
		tmpcreate = append(tmpcreate, field)
	}
	if len(tmpcreate) > 0 {
		err := DbEngin.Model(&model.Column{}).Create(tmpcreate).Error
		if err != nil {
			return table, err
		}
	}
	//
	for _, field := range mapexistcolumn {
		tmpupdate = append(tmpupdate, field)
		DbEngin.Model(&field).Updates(&field)
	}
	tmpcreate = append(tmpcreate, tmpupdate...)
	//kaishi export
	table.Columns = tmpcreate
	table.ProjectID = prj.ID
	table.Project = *prj
	table.Method = BuildMethod(methods)
	table.Types = types
	return table, nil
}

type BizType string

const (
	INITPROJECT  BizType = "init"
	EXPORTATABLE BizType = "export"
)

// tpldir 使用的模板路径
func ExportTable(table *model.Table, tpldir string, biz BizType, onfilegennerate ...CallbackFunc) error {
	return Render(table, tpldir, biz, onfilegennerate...)
}

const rootname = "./root.html"

func Render(table *model.Table, tpldir string, biz BizType, onfilegennerate ...CallbackFunc) (err error) {
	tmpls := template.New(rootname)
	tmpls = tmpls.Funcs(funcMaps)
	tmpls, err = tmpls.ParseGlob(tpldir + "/*.html")
	if err != nil {
		return err
	}
	// 初始化备份目录
	batchIndex := timekit.DateTimeNow(timekit.YYYYMMDDhhmmsspure)
	bakdir := filepath.Join(table.Project.Dirsave, ".bak", batchIndex)
	if !filekit.IsExist(bakdir) {
		os.MkdirAll(bakdir, 0755)
	}
	for _, tpl := range tmpls.Templates() {
		tplName := tpl.Name()
		//如果是inita table  业务,那么只处理包含 export_字段的模板
		if biz == EXPORTATABLE {
			if !strings.Contains(tplName, string(biz)+"_") {
				continue
			}
		}
		//过滤掉以html结尾的
		if strings.HasSuffix(tplName, ".html") {
			continue
		}
		//将
		dstFile := strings.ReplaceAll(tplName, "[module]", strings.ToLower(table.Module))
		dstFile = strings.ReplaceAll(dstFile, "[model]", strings.ToLower(table.Module))
		dstFile = strings.ReplaceAll(dstFile, "[prjname]", strings.ToLower(table.Project.Name))
		dstFile = strings.ReplaceAll(dstFile, "[Module]", table.Module)
		pkgpath := strings.ReplaceAll(table.Project.Package, ".", "/")
		dstFile = strings.ReplaceAll(dstFile, "[pkgpath]", pkgpath)

		dstFile = filepath.Join(table.Project.Dirsave, dstFile)
		dstFile = strings.TrimSuffix(dstFile, ".tpl")
		os.MkdirAll(filepath.Dir(dstFile), os.ModeDir)
		// 如果存在直接备份

		if filekit.Exists(dstFile) {
			bakfile := filepath.Join(bakdir, strings.ReplaceAll(dstFile, table.Project.Dirsave, ""))
			dir := filepath.Dir(bakfile)
			if !filekit.IsExist(dir) {
				os.MkdirAll(dir, 0755)
			}
			buf, err := os.ReadFile(dstFile)
			if err != nil {
				logger.Error(err.Error())
				return err
			}
			err = os.WriteFile(bakfile, buf, 0755)
			if err != nil {
				logger.Error(err.Error())
				return err
			}
		}
		f, e := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, 0766)
		if e != nil {
			return e
		}

		//文件需要再次清空
		err = f.Truncate(0)
		if err != nil {
			f.Close()
			logger.Error(err.Error())
			return
		}
		primary := model.Column{}
		for _, col := range table.Columns {
			if col.IsPrimaryKey {
				primary = col
				break
			}
		}
		rows := table.Columns
		rows1 := slicekit.Sort(rows, func(col1, col2 model.Column) bool {
			return col1.RawData.OrdinalPosition < col2.RawData.OrdinalPosition
		})
		tpldata := ExportVO{
			Name:    table.Name,
			Title:   table.Title,
			Module:  table.Module,
			Method:  table.Method,
			Columns: rows1,
			Primary: primary,
			Package: table.Project.Package,
			TplId:   table.Project.TplId,
			Table:   table,
			Types:   table.Types,
			Project: table.Project,
		}
		err = tpl.ExecuteTemplate(f, tplName, tpldata)
		f.Close()
		if err != nil {
			logger.Error("error", err.Error())
			continue
		}
		buf, _ := os.ReadFile(dstFile)

		content := string(buf)
		content = strings.ReplaceAll(content, "&lt;", "<")
		err = os.WriteFile(dstFile, []byte(content), 0766)
		if err != nil {
			continue
		}
		for _, callback := range onfilegennerate {
			callback(dstFile)
		}
	}
	os.Remove(rootname)
	return nil
}

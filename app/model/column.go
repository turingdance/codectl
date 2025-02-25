package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"

	"gorm.io/gorm"
)

const TableNameColumn = "column"

type RawData struct {
	ColumnName      string `gorm:"column_name"`
	DataType        string `gorm:"data_type"`
	CharMaxLength   int    `gorm:"char_max_length"`
	NumberPrecision int    `gorm:"number_precision"`
	NumberScale     int    `gorm:"number_scale"`
	ColumnComment   string `gorm:"column_comment"`
	ColumnKey       string `gorm:"column_key"`
	ColumnType      string `gorm:"column_type"`
	IsNullAble      string `gorm:"is_nullable"`
	Extra           string `gorm:"extra"`
	OrdinalPosition int    `gorm:"ordinal_position"`
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *RawData) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("数据类型不识别")
	}
	err := json.Unmarshal(bytes, j)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j RawData) Value() (driver.Value, error) {
	bts, err := json.Marshal(&j)
	return bts, err
}

/*
[{"id":0,"tableId":0,"dataColumn":"","dataType":"","dataSize":0,"title":"","domType":"","isPrimaryKey":"","autoIncrement":0,"isIndex":"","suportSearch":0,"placeholder":"","option":"","suportCreate":0,"suportUpdate":0,"serializer":"","rawdata":{"ColumnName":"id","DataType":"bigint","CharMaxLength":0,"NumberPrecision":20,"NumberScale":0,"ColumnComment":"ID","ColumnKey":"PRI","ColumnType":"bigint(20) unsigned","Extra":"auto_increment","OrdinalPosition":"1"}},
{"id":0,"tableId":0,"dataColumn":"","dataType":"","dataSize":0,"title":"","domType":"","isPrimaryKey":"","autoIncrement":0,"isIndex":"","suportSearch":0,"placeholder":"","option":"","suportCreate":0,"suportUpdate":0,"serializer":"","rawdata":{"ColumnName":"title","DataType":"varchar","CharMaxLength":255,"NumberPrecision":0,"NumberScale":0,"ColumnComment":"名称","ColumnKey":"","ColumnType":"varchar(255)","Extra":"","OrdinalPosition":"2"}},
{"id":0,"tableId":0,"dataColumn":"","dataType":"","dataSize":0,"title":"","domType":"","isPrimaryKey":"","autoIncrement":0,"isIndex":"","suportSearch":0,"placeholder":"","option":"","suportCreate":0,"suportUpdate":0,"serializer":"","rawdata":{"ColumnName":"age","DataType":"int","CharMaxLength":0,"NumberPrecision":10,"NumberScale":0,"ColumnComment":"年龄","ColumnKey":"","ColumnType":"int(11)","Extra":"","OrdinalPosition":"3"}},
{"id":0,"tableId":0,"dataColumn":"","dataType":"","dataSize":0,"title":"","domType":"","isPrimaryKey":"","autoIncrement":0,"isIndex":"","suportSearch":0,"placeholder":"","option":"","suportCreate":0,"suportUpdate":0,"serializer":"","rawdata":{"ColumnName":"birth","DataType":"datetime","CharMaxLength":0,"NumberPrecision":0,"NumberScale":0,"ColumnComment":"生日","ColumnKey":"","ColumnType":"datetime","Extra":"","OrdinalPosition":"4"}}]
*/
func (r RawData) BuildColumn(c *Column) *Column {
	c.DataColumn = r.ColumnName
	c.DataType = r.DataType
	c.Title = r.ColumnComment
	c.Title = r.ColumnComment
	c.IsPrimaryKey = (r.ColumnKey == "PRI")
	c.IsIndex = (r.ColumnKey != "" && r.ColumnKey != "PRI")
	c.AutoIncrement = (r.Extra == "auto_increment")
	if r.CharMaxLength > 0 {
		c.DataSize = int32(r.CharMaxLength)
	}
	if r.NumberPrecision > 0 {
		c.DataSize = int32(r.NumberPrecision)
	}
	c.DomType = "text"
	if strings.Contains(c.DataType, "int") {
		c.DomType = "number"
	} else if strings.Contains(c.DataType, "DateTime") || strings.Contains(c.DataType, "Time") {
		c.DomType = "datetime"
	} else if strings.Contains(c.DataType, "Date") {
		c.DomType = "date"
	} else {
		c.DataType = "text"
	}
	c.IsNullAble = (strings.Compare(r.IsNullAble, "YES") == 0)
	//logger.Debugf(r.ColumnName, r.IsNullAble,"YES",c.IsNullAble)
	c.RawData = r
	return c
}

// Column mapped from table <column>
type Column struct {
	ID            int32   `gorm:"column:id;type:int;primaryKey" json:"id"`
	TableID       int32   `gorm:"column:table_id;type:int" json:"tableId"`
	DataColumn    string  `gorm:"column:data_column;type:string;size:50" json:"dataColumn"`
	DataType      string  `gorm:"column:data_type;type:string;size:50" json:"dataType"`
	DataSize      int32   `gorm:"column:data_size;type:int" json:"dataSize"`
	Title         string  `gorm:"column:title;type:string;size:150" json:"title"`
	DomType       string  `gorm:"column:dom_type;type:string;size:20" json:"domType"`
	IsPrimaryKey  bool    `gorm:"column:is_primary_key;type:int" json:"isPrimaryKey"`
	AutoIncrement bool    `gorm:"column:auto_increment;type:int" json:"autoIncrement"`
	IsIndex       bool    `gorm:"column:is_index;type:string;size:50" json:"isIndex"`
	IsNullAble    bool    `gorm:"column:is_nullable;type:int" json:"isNullAble"`
	Placeholder   string  `gorm:"column:placeholder;type:string;size:50" json:"placeholder"`
	Option        string  `gorm:"column:option;type:string;size:250" json:"option"`
	SuportSearch  bool    `gorm:"column:suport_search;type:int;default:1" json:"suportSearch"`
	SuportCreate  bool    `gorm:"column:suport_create;type:int;default:1" json:"suportCreate"`
	SuportUpdate  bool    `gorm:"column:suport_update;type:int;default:1" json:"suportUpdate"`
	Sortable      bool    `gorm:"column:sortable;type:int;default:1" json:"sortable"`
	Hidden        bool    `gorm:"column:hidden;type:int;default:0" json:"hidden"`
	Serializer    string  `gorm:"column:serializer;type:string;size:50" json:"serializer"`
	RawData       RawData `gorm:"column:rawdata;type:string;size:250" json:"rawdata"`
}

// TableName Column's table name
func (s *Column) BeforeCreate(*gorm.DB) error {

	return nil
}

// TableName Column's table name
// func (*Column) TableName() string {
// 	return TableNameColumn
// }

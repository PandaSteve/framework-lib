package BaseLib

import (
	"github.com/PandaSteve/framework-lib/src/Lib/DataSource"
	"github.com/PandaSteve/framework-lib/src/Lib/MySql/Config"
	"strconv"
)

type Param struct {
	Name string
}

func (_this *Param) Equal(val interface{}) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: _this.Name, Type: MySqlConfigLib.Where, Value: val, Expression: _this.Name + " = ? "}
}

func (_this *Param) GreaterThan(val interface{}) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: _this.Name, Type: MySqlConfigLib.Where, Value: val, Expression: _this.Name + " < ? "}
}

func (_this *Param) LessThan(val interface{}) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: _this.Name, Type: MySqlConfigLib.Where, Value: val, Expression: _this.Name + " > ? "}
}

func (_this *Param) Like(val interface{}) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: _this.Name, Type: MySqlConfigLib.Where, Value: val, Expression: _this.Name + " like ? "}
}

func (_this *Param) InString(val ...string) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: _this.Name, Type: MySqlConfigLib.Where, Value: val, Expression: _this.Name + " in ? "}
}

func (_this *Param) InNumber(val ...int) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: _this.Name, Type: MySqlConfigLib.Where, Value: val, Expression: _this.Name + " in ? "}
}

func (_this *Param) Max() DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: _this.Name, Type: MySqlConfigLib.Field, Expression: " max(" + _this.Name + ") "}
}

func (_this *Param) Sum() DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: _this.Name, Type: MySqlConfigLib.Field, Expression: " sum(" + _this.Name + ") "}
}

func (_this *Param) Count() DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: _this.Name, Type: MySqlConfigLib.Field, Expression: " count(" + _this.Name + ") "}
}

func (_this *Param) Page(count int, page int) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: _this.Name, Type: MySqlConfigLib.Page, Expression: " limit  " + strconv.Itoa(count*(page-1)) + "," + strconv.Itoa(count)}
}

func (_this *Param) OrderByAsc() DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: _this.Name, Type: MySqlConfigLib.Order, Expression: _this.Name + " asc"}
}

func (_this *Param) OrderByDesc() DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: _this.Name, Type: MySqlConfigLib.Order, Expression: _this.Name + " desc"}
}

func (_this *Param) Group() DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: _this.Name, Type: MySqlConfigLib.Group, Expression: _this.Name}
}

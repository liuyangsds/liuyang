package liuyang

import (
	"fmt"
	"reflect"
)

//将结构体的字段名称和字段值以字符串相连接，供sql语句使用，返回字段、值、bool
func StructFieldAndValueToString(i interface{}) (string, string, bool) {
	//获取struct的field
	fobj := reflect.TypeOf(i)

	//判断参数类型如果不是struct就返回，即使是struct对象的指针也返回，因为指针传过来也不知道它是什么类型的指针。只能知道它是ptr类型。
	if fobj.Kind() != reflect.Struct {
		return "参数不是struct类型", "", false
	}

	//声明变量，接收struct的field字符串
	var fieldStr string
	for n := 0; n < fobj.NumField(); n++ {
		if n < fobj.NumField() - 1 {
			fieldStr += "`" + fobj.Field(n).Name + "`,"
		} else {
			fieldStr += "`" + fobj.Field(n).Name + "`"
		}
	}

	//获取struct的value
	vobj := reflect.ValueOf(i)
	//声明变量，接收struct的value字符串
	var valueStr string
	for m := 0; m < vobj.NumField(); m++ {
		if m < vobj.NumField() - 1 {
			valueStr += "'" + fmt.Sprintf("%v",vobj.Field(m)) + "',"
		} else {
			valueStr += "'" + fmt.Sprintf("%v",vobj.Field(m)) + "'"
		}
	}
	return fieldStr, valueStr, true
}


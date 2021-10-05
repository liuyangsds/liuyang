package liuyang

import (
	"fmt"
	"reflect"
)

//获取参数类型，以kind()返回，这样方便和其它类型判断
func GetType(i interface{}) reflect.Kind {
	return reflect.TypeOf(i).Kind()
}

//获取变量类型，返回系统类型的名称
func Get_T(i interface{}) string {//struct
	obj := reflect.TypeOf(i)
	return obj.Kind().String()
}
//获取变量类型，返回实际类型的名称(具体的实际类型，更细节)
func Get_TT(i interface{}) string {//main.Fangchong
	return fmt.Sprintf("%T", i)
}


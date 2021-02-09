package liuyang

import (
	"fmt"
	"reflect"
)

//可打印map或struct类型的数据，并打印其键和值
func Print_i(i interface{}) {

	obj := reflect.TypeOf(i)//Type类型

	if obj.Kind() == reflect.Map {
		//使用接口断言的方式，判断其类型是否是map[string]interface{}
		mapList,mapErr := i.(map[string]interface{})//返回两个参数，第一个是map值，第二个是bool类型
		if mapErr == true {
			fmt.Println("==================打印",obj.Kind(),"类型的",obj,"数据开始==================")
			for key,value := range mapList {
				fmt.Println(key,"\t",value)
			}
			fmt.Println("==================打印",obj.Kind(),"类型的",obj,"数据结束==================")
			return
		}
		fmt.Println("无法打印：",mapList)
		return
	}
	if obj.Kind() == reflect.Struct {
		//s,e := i.(Person)//这里无需判断其类型是否是Person了，因为只要是struct类型的值，那么一律打印即可。
		//fmt.Println(e)//true
		//fmt.Println(s)//{刘阳 男 25 0 0 }

		//得到struct的值
		obj_v := reflect.ValueOf(i)

		fmt.Println("==================打印",obj.Kind(),"类型的",obj,"数据开始==================")
		for i := 0; i < obj.NumField(); i++ {
			fmt.Println(obj.Field(i).Name,"\t",obj_v.Field(i))
		}
		fmt.Println("==================打印",obj.Kind(),"类型的",obj,"数据结束==================")
		return
	}

	fmt.Println("其他类型：",i)
}
//打印参数的类型
func Print_type(i interface{}) {
	obj := reflect.TypeOf(i)
	fmt.Println("当前变量的类型为：",obj.Kind())
}
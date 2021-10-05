package liuyang

import (
	"encoding/json"
	"errors"
	"reflect"
)

//封装struct转map、struct转struct、map转struct、map转map。
//复杂数据结构间的转换函数
func JsonToObject(input interface{}, output interface{}) error {
	//首先要判断两个参数的类型，只有指针才可以进行以下操作，否则直接返回错误
	//ptr1 := reflect.TypeOf(input).//普通方式获取类型，判断时需要使用.Kind()
	//ptr2 := reflect.TypeOf(output)//普通方式获取类型，判断时需要使用.Kind()
	//用封装函数获取类型
	if GetType(input) != reflect.Ptr || GetType(output) != reflect.Ptr {
		return errors.New("参数类型不是指针，请取变量地址进行传参")
	}

	//将input(struct或map)指针编码为json数据
	jsonByte,err := json.Marshal(input)
	if err != nil {
		return err
	}
	//解析json并赋到output(struct或map)指针上
	merr := json.Unmarshal(jsonByte, output)
	if merr != nil {
		return merr
	}

	return nil
}

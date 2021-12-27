package liuyang

import (
	"fmt"
	"reflect"
)

//打印map或struct类型的数据的键和值
func Print_I(ia interface{}) {

	obj := reflect.TypeOf(ia) //Type类型

	if obj.Kind() == reflect.Map {
		//使用接口断言的方式，判断其类型是否是map[string]interface{}
		mapList, mapErr := ia.(map[string]interface{}) //返回两个参数，第一个是map值，第二个是bool类型
		if mapErr == true {
			fmt.Println("==================打印", obj.Kind(), "类型的", obj, "数据开始==================")
			for key, value := range mapList {
				fmt.Println(key, "\t", value)
			}
			fmt.Println("==================打印", obj.Kind(), "类型的", obj, "数据结束==================")
			return
		}
		fmt.Println("无法打印：", mapList)
		return
	}
	if obj.Kind() == reflect.Struct {
		//s,e := i.(Person)//这里无需判断其类型是否是Person了，因为只要是struct类型的值，那么一律打印即可。
		//fmt.Println(e)//true
		//fmt.Println(s)//{刘阳 男 25 0 0 }

		//得到struct的值
		obj_v := reflect.ValueOf(ia)

		fmt.Println("==================打印", obj.Kind(), "类型的", obj, "数据开始==================")
		for i := 0; i < obj.NumField(); i++ {
			fmt.Println(obj.Field(i).Name, "\t", obj_v.Field(i))
		}
		fmt.Println("==================打印", obj.Kind(), "类型的", obj, "数据结束==================")
		return
	}

	fmt.Println("其他类型：", ia)
}

//打印参数的类型，得出变量的系统类型
func Print_T(i interface{}) { //struct，slice
	obj := reflect.TypeOf(i)
	fmt.Println("当前值的类型为：", obj.Kind().String())
}

//打印参数的类型，得出变量的实际类型(具体的实际类型，更细节)，如下打印类型的对比
func Print_TT(i interface{}) { //main.Fangchong，[]int
	fmt.Printf("当前值的类型为：%T\n", i)
}

/*
打印类型：

	//chan时
	var cc chan int
	fmt.Println(cc)//nil
	test.Print_T(cc)//chan
	fmt.Printf("%T", cc)//chan int

	ff := make(chan int)
	test.Print_T(ff)//chan
	fmt.Printf("%T", ff, ff)//chan int
	fmt.Printf("%v", ff, ff)//0xc0003981e0

	//结构体时
	ff := Fangchong{}
	test.Print_T(ff)//struct
	fmt.Printf("%T", ff)//main.Fangchong

	//map时
	ff := make(map[string]interface{})
	test.Print_T(ff)//map
	fmt.Printf("%T", ff)//map[string]interface {}

	//数组时
	ff := []byte{0,2}
	test.Print_T(ff)//slice
	fmt.Printf("%T", ff)//[]uint8

	//数组时，不指定元素个数被视为切片
	ff := []int{0,2}
	test.Print_T(ff)//slice
	fmt.Printf("%T", ff)//[]int

	//数组时
	ff := [5]int{2,1,4,2,1}
	test.Print_T(ff)//array
	fmt.Printf("%T", ff)//[5]int

	//int
	ff := 24
	test.Print_T(ff)//int
	fmt.Printf("%T", ff)//int

	//string
	ff := "haha"
	test.Print_T(ff)//string
	fmt.Printf("%T", ff)//string
*/

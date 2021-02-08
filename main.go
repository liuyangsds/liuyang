package main

import (
	"fmt"
	"liuyang/liuyang"
)

func main() {
	fmt.Println("======================================")

	liuyang.Base64_test()

	//liuyang.FileXuChuan("sunlulu/dibai.jpg", "nihao/a/b/c/d/e/saobi.png")

	//aa, bb := liuyang.FileFolderTotal("nihao/a", 0)
	//fmt.Println("文件夹累加数量", aa)//12
	//fmt.Println("文件累加数量" ,bb)//1

	//mm := make(map[string]interface{}, 0)
	//
	//mm["name "] = "刘阳"
	//mm["age"] = 25
	//mm["height"] = 175.05

	//
	//str := liuyang.MapToString(mm)
	//fmt.Println(str)
	//
	//aa := mm["age"]
	//bb := liuyang.ValueToString(aa)
	//fmt.Println(bb)

	//str := "liuyang"
	//aa := liuyang.MD5(str)
	//fmt.Println(aa)
	//
	//aa1 := liuyang.MD5(liuyang.MD5(str))
	//fmt.Println(aa1)

	//aa,err := liuyang.PassEncrypt("123456")
	//	//fmt.Println(err)
	//	//fmt.Println(aa)
	//	//
	//	//
	//	//bb := liuyang.PassVerify(aa, "123456")
	//	//fmt.Println("密码是否相同：", bb)

	//随机生成指定长度的字符串
	//str := liuyang.RandomString(5)
	//fmt.Println(str)
	//istrue := liuyang.Check_money_6("0")
	//fmt.Println(istrue)

	//aa := "caa"
	//bb := "caa "
	//ss := liuyang.StringCompareInt(aa, bb)
	//dd := strings.Compare(aa, bb)
	//
	//fmt.Println("刘阳写的两个字符串的大小为：", ss)
	//fmt.Println("系统内置两个字符串的大小为：", dd)

	arr := []int{3,2,0,8,9,4,5,7}
	newArr := liuyang.SliceCopyInt(arr)

	newArr = liuyang.SliceDeleteKeyInt(newArr, 4)
	fmt.Println(newArr)		//[3 2 0 8 4 5 7 7]
	fmt.Println(arr)		//[3 2 0 8 9 4 5 7]
	//fmt.Println(list)		//[3 2 0 8 4 5 7]
	//
	//UpdateArrarValue(list)
	//fmt.Println(list)		//[3 2 0 8 4 5 7]
	//fmt.Println("newArr中的元素个数为：", len(newArr))
	//说明，将newArr数组中的元素9删除后，其元素9虽然不见了，但是由于只是以覆盖操作而进行的假删除并未将其元素个数改变
	//即使删除一个元素9后，newArr中的元素个数依然是8个



	//mm := make(map[string]interface{}, 0)
	//mm["mobile"] = "13000000001"
	//mm["password"] = "123456"
	//
	////测试post提交
	//url := "http://localhost:8090/user/login"
	//
	//result, err := liuyang.HttpRequestPOST(url,mm)
	//fmt.Println("是否错误：",err)//<nil>
	//fmt.Println("返回值为：",result)//{"code":0,"data":{"info":{"uid":6660001}, ... }



	fmt.Println("======================================")



	//mm := make(map[string]interface{}, 0)
	//
	//mm["name "] = "刘阳"
	//mm["age"] = 25
	//mm["height"] = 175.05
	//
	//fmt.Println(mm)//map[age:25 height:175.05 name :刘阳]
	////删除map中的健值对
	//delete(mm, "age")
	//fmt.Println(mm)//map[height:175.05 name :刘阳]


}

//声明函数，参数为：切片
func UpdateArrarValue(arr []int)  {
	arr[0] = 188	//将参数数组的第一个元素设置为188
}

type Person struct {
	Mobile string `json:"mobile"`
	Password string `json:"password"`
}
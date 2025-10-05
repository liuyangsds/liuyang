package liuyang

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
)

// -----------------------------------------------去重相关-----------------------------------------------
// 数组检查重复元素，false不重复、true重复
// 样例1
// arr := []int{3, 7, 8, 9, 12, 5, 30, 2}
// ok := ArrayCheckRepeat(arr)
// fmt.Println(ok)
// 打印
// false
// 样例2
// arr := []int{3, 7, 8, 9, 2, 5, 3, 2}
// ok := ArrayCheckRepeat(arr)
// fmt.Println(ok)
// 打印
// true
func ArrayCheckRepeat[T cmp.Ordered](arr []T) bool {
	//临时变量，标记是否重复，默认不重复
	flag := false

	//定义map
	mm := make(map[T]int)
	//遍历切片
	for _, value := range arr {
		//fmt.Println(key, "---", value)
		//map指定下标的值累加1，这里可以放心执行，因为是map，即使没有对应下标也会创建。不像php那样先行判断有无指定下标，没有的话直接使用会报错的情况。
		mm[value]++
		//如果指定的map下标中的值大于1，说明已存在相同元素，标记为true
		if mm[value] > 1 {
			//有重复元素
			flag = true
			//停止遍历
			break
		}
	}

	return flag
}

// 数组移除重复元素
// arr := []int{3, 7, 8, 9, 2, 5, 3, 2}
// aa := ArrayRemoveRepeat(arr)
// fmt.Println(aa)
// 打印
// [3 7 8 9 2 5]
func ArrayRemoveRepeat[T cmp.Ordered](arr []T) []T {
	//定义map
	mm := make(map[T]int)
	//遍历切片
	for _, value := range arr {
		//fmt.Println(key, "---", value)
		//map指定下标的值累加1，这里可以放心执行，因为是map，即使没有对应下标也会创建。不像php那样先行判断有无指定下标，没有的话直接使用会报错的情况。
		mm[value] = 1
	}

	//临时变量，用于接收map中的key(也就是arr中的元素)
	tempArr := make([]T, 0)
	for key, _ := range mm {
		tempArr = append(tempArr, key)
	}

	return tempArr
}

//-----------------------------------------------去重相关-----------------------------------------------

// -----------------------------------------------查找相关-----------------------------------------------
// 数组搜索元素，返回元素索引，未找到时返回-1
// 样例1：
// arr := []int{3, 7, 8, 9, 2, 5, 3, 2}
// s := 8
// index := ArraySearch(arr, s)
// fmt.Println(index)
// 打印
// 2
// 样例2：
// s := 80
// ok := ArraySearch(arr, s)
// fmt.Println(ok)
// 打印
// -1
func ArraySearch[T cmp.Ordered](arr []T, elem T) int {

	for key, value := range arr {
		if value == elem {
			return key
		}
	}

	return NotFoundIndex
}

// 判断某元素是否在数组中，类似php中的函数in_array()。false不在、true在
// 样例1：
// arr := []int{3, 7, 8, 9, 2, 5, 3, 2}
// s := 8
// ok := InArray(arr, s)
// fmt.Println(ok)
// 打印
// true
// 样例2：
// s := 80
// ok := InArray(arr, s)
// fmt.Println(ok)
// 打印
// false
func InArray[T cmp.Ordered](arr []T, elem T) bool {

	index := ArraySearch(arr, elem)
	if index != NotFoundIndex {
		return true
	}

	return false
}

//-----------------------------------------------查找相关-----------------------------------------------

// -----------------------------------------------排序相关-----------------------------------------------
// 数组排序
// 标准冒泡排序
// 参数1：原数组
// 参数2：0正序、1倒序
// 样例
// arr := []int{3, 8, 9, 2, 7, 5, 1}
// ArraySort(arr, 0)
// fmt.Println(arr)
// 打印
// [1 2 3 5 7 8 9]
func ArraySort[T cmp.Ordered](arr []T, sort uint8) {
	if len(arr) < 2 || sort > 1 {
		return //这里是直接返回
	}

	var temp T

	if sort == 0 {
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr)-(1+i); j++ {
				if arr[j] > arr[j+1] {
					temp = arr[j]
					arr[j] = arr[j+1]
					arr[j+1] = temp
					//arr[j],arr[j+1] = arr[j+1],arr[j]
				}
			}
		}
	} else if sort == 1 {
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr)-(1+i); j++ {
				if arr[j] < arr[j+1] {
					temp = arr[j]
					arr[j] = arr[j+1]
					arr[j+1] = temp
					//arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}

	//return arr
}

// 英文字符串排序，使用冒泡排序，在比较两个字符串的大小时，使用刘阳封装的方法CompareStringInt进行比较。
func ArraySortString(arr []string, sort uint8) []string {
	if len(arr) < 2 || sort > 1 {
		return arr //这里是直接返回
	}

	if sort == 0 {
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr)-(1+i); j++ {
				if CompareString(arr[j], arr[j+1]) > 0 {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	} else if sort == 1 {
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr)-(1+i); j++ {
				if CompareString(arr[j], arr[j+1]) < 0 {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}
	return arr
}

//-----------------------------------------------排序相关-----------------------------------------------

// -----------------------------------------------删除相关-----------------------------------------------
// 删除一个数组元素，返回删除后的切片
// 样例1
// arr := []int{3, 7, 8, 9, 2, 5, 3, 2}
// dd := test_new.ArrayDelIndex(arr, 2)
// fmt.Println(dd)
// fmt.Println(arr)
// 打印
// [3 7 9 2 5 3 2]
// [3 7 9 2 5 3 2 2]
// 需使用原数组接收才行，如下
// 样例2
// arr := []int{3, 7, 8, 9, 2, 5, 3, 2}
// arr := test_new.ArrayDelIndex(arr, 2)
// fmt.Println(arr)
// 打印
// [3 7 9 2 5 3 2]
func ArrayDelIndex[T any](arr []T, index int) []T {

	arr = append(arr[:index], arr[index+1:]...)

	return arr
}

// 删除一个或多个数组元素，参数1，数组，参数2：下标，参数3：要删除的个数。返回删除后的切片
// 样例
// arr := []int{3, 7, 8, 9, 2, 5, 3, 2}
// arr = test_new.ArrayDelIndexLength(arr, 2, 3)
// fmt.Println(arr)
// 打印
// [3 7 5 3 2]
func ArrayDelIndexLength[T any](arr []T, index int, length int) []T {

	if length < 1 { //当要删除的个数小于1也就是等于0或负数时，就直接返回数组
		return arr
	}

	keyCount := len(arr)     //元素数量
	offset := index + length //当前key的下标和要删除的数量
	//当元素的总个数大于偏移量时，也就是当前要被删除的下标加上要删除的个数的和小于元素总个数，那么是可操作的。
	if keyCount > offset {
		arr = append(arr[:index], arr[offset:]...)
		return arr
	}

	//否则偏移量就是大于或等于元素总个数，那么,这里不能写成：s[keyCount-1]...因为这样写，会把该数组中的最后一个元素加入进来
	arr = append(arr[:index], arr[keyCount:]...) //当前key的位置之前的所有元素和该数组的总元素数及以后的元素，总元素数的下标及以后也就是没有元素

	return arr
}

// 收集要删除的索引值。参数：1要操作的切片、2要保留的索引。返回值：要删除的索引切片。该返回值主要配合liuyang.ArrayRemoveUint32()函数使用。
// 前题条件：要操作的切片中元素值不能为0，因为要保留的索引值就是以0为标记的。
// 样例
// arr := []int{3, 7, 8, 9, 2, 5, 3, 2}
// saveArr := []uint32{3, 5}
// delArr = test_new.ArrayCollectDelIndex(arr, saveArr)
// fmt.Println(delArr)
// 打印
// [0 1 2 4 6 7]
func ArrayCollectDelIndex[T int | int32 | uint32 | int64](arr []T, saveIndex []uint32) []T {
	//过滤非法参数，但却漏掉了saveIndex参数值，因为它是[]uint32的空值。但却不为nil
	//saveArr := []uint32{} //要保留的索引值在定义时设为某个类型的空值，但却不为nil。这就会让saveIndex == nil 的判断失效。
	//if arr == nil || saveIndex == nil {
	//	return nil
	//}

	//刘阳建议使用长度判断切片是否有值
	if len(arr) == 0 || len(saveIndex) == 0 {
		return nil
	}

	if len(saveIndex) > len(arr) {
		return nil
	}

	//拷贝临时数组，用于标记要保留的索引
	tempPotArr := make([]T, len(arr))
	copy(tempPotArr, arr)

	//用于接收要删除的索引值
	var delArr = make([]T, 0)

	//fmt.Println("拷贝的临时数组值：", tempPotArr)
	for i := 0; i < len(saveIndex); i++ {
		//如果要保留的索引值超出或等于要操作的切片总长度时，为非法
		if int(saveIndex[i]) >= len(tempPotArr) {
			return nil
		}
		//将要保留的索引(1、3)标记为0
		tempPotArr[saveIndex[i]] = 0
	}

	//遍历临时数组
	var i T = 0                     //声明i的类型
	var length = T(len(tempPotArr)) //获取临时数组长度
	for i = 0; i < length; i++ {
		//只要元素值大于0，就是要删除的索引值
		if tempPotArr[i] > 0 {
			//将要删除的索引值追加到临时删除索引收集数组中
			delArr = append(delArr, i)
		}
	}

	return delArr
}

// 收集要删除的索引值。参数：1要操作的切片、2要保留的索引。返回值：要删除的索引切片。该返回值主要配合liuyang.ArrayRemoveUint32()函数使用。
// 前题条件：要操作的切片中元素值不能为0，因为要保留的索引值就是以0为标记的。
// string类型要注意。定义切片时不要将元素值设置为""，如：[]string{"刘阳", "", "", "璐璐", "凤凤"}，这样会误删掉1、2索引。
// 建议定义切片时将元素值设置为" "，带一个空格。这样就不会发生误删除的情况了。
// 样例
// arr := []string{"a", "b", "c", "d", "e", "f"}
// saveArr := []uint32{3, 5}
// delArr := test_new.ArrayCollectDelIndexString(arr, saveArr)
// fmt.Println(delArr)
// 打印
// [0 1 2 4]
func ArrayCollectDelIndexString(arr []string, saveIndex []uint32) []uint32 {
	//过滤非法参数，但却漏掉了saveIndex参数值，因为它是[]uint32的空值。但却不为nil
	//saveArr := []uint32{} //要保留的索引值在定义时设为某个类型的空值，但却不为nil。这就会让saveIndex == nil 的判断失效。
	//if arr == nil || saveIndex == nil {
	//	return nil
	//}

	//刘阳建议使用长度判断切片是否有值
	if len(arr) == 0 || len(saveIndex) == 0 {
		return nil
	}

	if len(saveIndex) > len(arr) {
		return nil
	}

	//拷贝临时数组，用于标记要保留的索引
	tempPotArr := make([]string, len(arr))
	copy(tempPotArr, arr)

	//用于接收要删除的索引值
	var delArr = make([]uint32, 0)

	//fmt.Println("拷贝的临时数组值：", tempPotArr)
	for i := 0; i < len(saveIndex); i++ {
		//如果要保留的索引值超出或等于要操作的切片总长度时，为非法
		if int(saveIndex[i]) >= len(tempPotArr) {
			return nil
		}
		//将要保留的索引(1、3)标记为0
		tempPotArr[saveIndex[i]] = ""
	}

	//遍历临时数组
	var i uint32 = 0                     //声明uint32类型的i
	var length = uint32(len(tempPotArr)) //获取临时数组长度
	for i = 0; i < length; i++ {
		//只要元素值大于0，就是要删除的索引值
		if tempPotArr[i] != "" {
			//将要删除的索引值追加到临时删除索引收集数组中
			delArr = append(delArr, i)
		}
	}

	return delArr
}

// -----------------------------------------------删除相关-----------------------------------------------
// 批量删除切片索引。参数：1要操作的切片、2要删除的索引。返回值：已批量删除索引后的切片。uint32类型
// 样例
// arr := []int{3, 7, 8, 9, 2, 5, 3, 2}
// delIndex := []uint32{1, 3, 5}
// arr := test_new.ArrayRemove(arr, delIndex)
// fmt.Println(arr)
// 打印
// [3 8 2 3 2]
func ArrayRemove[T int | int32 | uint32 | int64](arr []T, delIndex []uint32) []T {
	//fmt.Println("原切片值：", arr)
	//如果要删除的索引数量大于原数组中的元素数量时，直接返回原数组
	if len(delIndex) > len(arr) {
		return arr
	}

	//原数组元素数量总计
	arrCount := len(arr)

	//先遍历先删除的索引值，将原数组中对应索引赋值为0进行标记，也就是待删除
	for _, value := range delIndex {
		//如果要删除的某个索引值大于等于原数组总长度时，说明不合法，则跳过
		if int(value) >= arrCount {
			continue
		}
		//将要删除的索引赋为0值
		arr[value] = 0
	}
	//fmt.Println("赋0后，切片值：", arr)

	//遍历删除值为0的元素
	for i := 0; i < len(arr); { //切记，这里的数组长度必须实时获取，不能用变量接收，不然会数组越界
		if arr[i] == 0 {
			//fmt.Println("判断等于0时------>当前索引：", i, "，当前值：", arr[i])
			arr = append(arr[:i], arr[i+1:]...)
			//fmt.Println("删除当前元素后的切片后，长度：", len(arr), "，值是：", arr)
		} else {
			//fmt.Println("否则大于0时------>当前索引：", i, "，当前值：", arr[i])
			i++
		}
	}

	return arr
}

// 批量删除切片索引。参数：1要操作的切片、2要删除的索引。返回值：已批量删除索引后的切片。string类型
// 样例
// arr := []string{"a", "b", "c", "d", "e", "f"}
// delIndex := []uint32{1, 3, 5}
// arr = test_new.ArrayRemoveString(arr, delIndex)
// fmt.Println(arr)
// 打印
// [a c e]
func ArrayRemoveString(arr []string, delIndex []uint32) []string {
	//fmt.Println("原切片值：", arr)
	//如果要删除的索引数量大于原数组中的元素数量时，直接返回原数组
	if len(delIndex) > len(arr) {
		return arr
	}

	//原数组元素数量总计
	arrCount := uint32(len(arr))

	//先遍历先删除的索引值，将原数组中对应索引赋值为0进行标记，也就是待删除
	for _, value := range delIndex {
		//如果要删除的某个索引值大于等于原数组总长度时，说明不合法，则跳过
		if value >= arrCount {
			continue
		}
		//将要删除的索引赋为0值
		arr[value] = ""
	}
	//fmt.Println("赋0后，切片值：", arr)

	//遍历删除值为0的元素
	for i := 0; i < len(arr); { //切记，这里的数组长度必须实时获取，不能用变量接收，不然会数组越界
		if arr[i] == "" {
			//fmt.Println("判断等于0时------>当前索引：", i, "，当前值：", arr[i])
			arr = append(arr[:i], arr[i+1:]...)
			//fmt.Println("删除当前元素后的切片后，长度：", len(arr), "，值是：", arr)
		} else {
			//fmt.Println("否则大于0时------>当前索引：", i, "，当前值：", arr[i])
			i++
		}
	}

	return arr
}

// -----------------------------------------------拷贝相关-----------------------------------------------
// 复制一个数组切片
// 样例
// arr := []string{"a", "b", "c", "d", "e", "f"}
// tempArr := test_new.ArrayCopy(arr)
// tempArr[1] = "666"
// fmt.Println("拷贝修改后：", tempArr)
// fmt.Println("原数据：", arr)
// 打印
// 拷贝修改后： [a 666 c d e f]
// 原数据： [a b c d e f]
func ArrayCopy[T any](arr []T) []T {
	var temp = make([]T, len(arr))
	//将原数据数组拷贝给临时数组。拷贝后，返回元素数量
	copy(temp, arr)

	return temp
}

// 复制一个切片范围，参数1，原数组，参数2，起始位置，参数3，复制的长度。返回复制后的新数组
// 样例1
// arr := []string{"a", "b", "c", "d", "e", "f"}
// tempArr := test_new.ArrayCopyRange(arr, 1, 3)
// tempArr[1] = "666"
// fmt.Println("拷贝修改后：", tempArr)
// fmt.Println("原数据：", arr)
// 打印
// 拷贝修改后： [b 666 d]
// 原数据： [a b c d e f]
// 样例2
// arr := []string{"a", "b", "c", "d", "e", "f"}
// tempArr := test_new.ArrayCopyRange(arr, 1, 30)
// tempArr[1] = "666"
// fmt.Println(len(tempArr), cap(tempArr))
// fmt.Println("拷贝修改后：", tempArr)
// fmt.Println("原数据：", arr)
// 5 5
// 拷贝修改后： [b 666 d e f]
// 原数据： [a b c d e f]
func ArrayCopyRange[T any](arr []T, start int, length int) []T {

	//首先要判断起始位置是否在合法范围内，如：[]int{2,3,5}//调用时用：arr,4,1，从下标4开始获取就属于非法操作
	if len(arr) < 1 || start < 0 || length < 1 { //如果数组长度小于1或下标小于0或想要获取的长度小于1则直接返回数组
		//fmt.Println("第1种情况")
		return arr
	}

	if start > (len(arr) - 1) { //如果下标位置数大于数组长度，则直接返回数组
		//fmt.Println("第2种情况")
		return arr
	}

	//差值
	offset := 0
	//判断length如果小于等于0或length大于(数组长度减起始位置)的差值时，也就是arr[2,3,5]有3个值时却想从下标1开始取6个数时：
	if length > (len(arr) - start) {
		//fmt.Println("这里是length小于等于0或length大于数组长度-起始位置的差值时")
		count := len(arr)
		//新数组长度差值 = 数组元素总个数减起始位置
		offset = count - start
		//将差值设置为临时数组的长度
		var temp = make([]T, offset)
		//将原数据数组拷贝给临时数组。其起始位置为start，结束位置则是原数组的总长度。
		copy(temp, arr[start:count]) //返回值为临时数组的长度

		return temp
	}

	//常规情况下
	//差值 = 起始位置 + 要复制的长度
	offset = start + length
	//将复制的长度设置为临时数组的长度
	var temp = make([]T, length)
	//将原数据数组拷贝给临时数组。起始位置为start，结束位置则是起始位置加上要复制的长度
	copy(temp, arr[start:offset])

	return temp
}

//-----------------------------------------------拷贝相关-----------------------------------------------

// -----------------------------------------------转字符串相关-----------------------------------------------
// 将所有类型的数组分割成以指定符号连接的字符串。该函数简单暴力，和下面ArrayImplode()函数的功能是一样的。
// 样例
// strArr := []string{"b", "d", "e", "aa"}
// bb := test_new.ArrayToString(strArr, ",")
// 打印
// b,d,e,aa
func ArrayToString[T any](arr T, sep string) string {
	temp_str := fmt.Sprint(arr)
	//fmt.Println(str) //[1 2 3 4 5 6 7 8 9]，字符串类型

	tt := strings.Trim(temp_str, "[]")
	//fmt.Println(tt) //1 2 3 4 5 6 7 8 9，去掉首尾字符：中括号

	//str := strings.Replace(tt, " ", ",", 5)
	//fmt.Println(str)//1,2,3,4,5,6 7 8 9

	//str1 := strings.Replace(tt, " ", ",", -1)
	//fmt.Println(str1) //1,2,3,4,5,6,7,8,9

	str := strings.Replace(tt, " ", sep, -1)

	return str
}

// 获取一个由int类型切片元素转成以指定符号连接的字符串。该方法学习自2010年长春好萝卜公司的同事冷志伟先生。
// 样例
// arr := []int{3, 7, 8, 9, 2, 5, 3, 2}
// ids := test_new.ArrayImplode(arr, ",")
// fmt.Println(ids)
// 打印
// 3,7,8,9,2,5,3,2
func ArrayImplode(arr []int, sep string) string {
	//临时变量
	var str string
	for i := 0; i < len(arr); i++ {
		//如果当前i的值小于元素总个数-1时，也就是最后一个key之前的值时，后面可以连接逗号
		if i < len(arr)-1 {
			str += strconv.Itoa(arr[i]) + sep
		} else {
			//否则i的值等于最后一个key时，后面就不要连接逗号了
			str += strconv.Itoa(arr[i])
		}
	}

	return str
}

//-----------------------------------------------转字符串相关-----------------------------------------------

//-----------------------------------------------统计相关-----------------------------------------------

// 数组中相同元素数量统计。返回：元素值与元素出现次数
// 样例
// arr := []int{3, 7, 8, 9, 2, 5, 3, 2}
// mm := test_new.ArraySameElementCount(arr)
// fmt.Println(mm)
// 打印
// map[2:2 3:2 5:1 7:1 8:1 9:1]
func ArraySameElementCount[T any](arr []T) map[any]int {
	//校验长度
	if len(arr) < 1 {
		return nil
	}
	//对象键值对法
	//该方法执行的速度比其他任何方法都快，就是占用的内存大一些
	tempMap := make(map[any]int, 0)

	//fmt.Println("初始，map的值：", tempMap, "，长度：", len(tempMap))

	for _, value := range arr {
		if _, ok := tempMap[value]; ok == true {
			tempMap[value]++
			//fmt.Println("是否有相同的key：", ok, value)
		} else {
			tempMap[value] = 1
		}
	}

	//fmt.Println("=======================遍历=======================")
	//for key, value := range tempMap {
	//	fmt.Println("map的key：", key, "------>value：", value)
	//}

	return tempMap
}

//-----------------------------------------------统计相关-----------------------------------------------

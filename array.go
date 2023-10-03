package liuyang

import (
	"fmt"
	"strconv"
	"strings"
)

//数组元素排序======================================

//标准冒泡排序，参数1为原数组，参数2如果为0则正序，为1则倒序。
func ArraySortByte(arr []byte, sort uint8) {
	if len(arr) < 2 || sort > 1 {
		return //这里是直接返回
	}

	if sort == 0 {
		var temp byte = 0
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
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}

	//return arr
}

//标准冒泡排序，参数1为原数组，参数2如果为0则正序，为1则倒序。
func ArraySortInt(arr []int, sort uint8) {
	if len(arr) < 2 || sort > 1 {
		return //这里是直接返回
	}

	if sort == 0 {
		temp := 0
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
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}

	//return arr
}

//标准冒泡排序，参数1为原数组，参数2如果为0则正序，为1则倒序。
func ArraySortUInt32(arr []uint32, sort uint8) {
	if len(arr) < 2 || sort > 1 {
		return //这里是直接返回
	}

	if sort == 0 {
		var temp uint32 = 0
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
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}

	//return arr
}

//英文字符串排序，使用冒泡排序，在比较两个字符串的大小时，使用刘阳封装的方法CompareStringInt进行比较。
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

//数组元素排序======================================

//删除数组切片======================================

//删除一个切片元素，返回byte类型切片
func ArrayDeleteKeyByte(s []byte, index int) []byte {

	s = append(s[:index], s[index+1:]...)

	return s
}

//删除一个切片元素，返回int类型切片
func ArrayDeleteKeyInt(s []int, index int) []int {

	//原数据：[3 2 0 8 9 4 5 7]
	//fmt.Println("切片的长度为：", len(s))//8
	//fmt.Println("前置元素：", s[:index])//[3 2 0 8]
	//fmt.Println("后置元素：", s[index + 1:])//[4 5 7]
	//将下标4，也就是为9的元素删除，其操作，只是覆盖而已
	s = append(s[:index], s[index+1:]...)

	return s
}

//删除一个切片元素，返回uint32类型切片
func ArrayDeleteKeyUInt32(s []uint32, index int) []uint32 {

	s = append(s[:index], s[index+1:]...)

	return s
}

//删除一个切片元素，返回uint64类型切片
func ArrayDeleteKeyUInt64(s []uint64, index int) []uint64 {

	s = append(s[:index], s[index+1:]...)

	return s
}

//删除一个切片元素，返回string类型切片
func ArrayDeleteKeyString(s []string, index int) []string {

	s = append(s[:index], s[index+1:]...)

	return s
}

//删除一个或多个切片元素，参数1，切片，参数2：下标，参数3：要删除的个数。返回byte类型切片
func ArrayDeleteKeyByteLength(s []byte, index int, length int) []byte {

	if length < 1 { //当要删除的个数小于1也就是等于0或负数时，就直接返回数组
		return s
	}

	keyCount := len(s)       //元素数量
	offset := index + length //当前key的下标和要删除的数量
	//当元素的总个数大于偏移量时，也就是当前要被删除的下标加上要删除的个数的和小于元素总个数，那么是可操作的。
	if keyCount > offset {
		s = append(s[:index], s[offset:]...)
		return s
	}

	//否则偏移量就是大于或等于元素总个数，那么,这里不能写成：s[keyCount-1]...因为这样写，会把该数组中的最后一个元素加入进来
	s = append(s[:index], s[keyCount:]...) //当前key的位置之前的所有元素和该数组的总元素数及以后的元素，总元素数的下标及以后也就是没有元素

	return s
}

//删除一个或多个切片元素，参数1，切片，参数2：下标，参数3：要删除的个数。返回int类型切片
func ArrayDeleteKeyIntLength(s []int, index int, length int) []int {

	if length < 1 { //当要删除的个数小于1也就是等于0或负数时，就直接返回数组
		return s
	}

	keyCount := len(s)       //元素数量
	offset := index + length //当前key的下标和要删除的数量
	//当元素的总个数大于偏移量时，也就是当前要被删除的下标加上要删除的个数的和小于元素总个数，那么是可操作的。
	if keyCount > offset {
		s = append(s[:index], s[offset:]...)
		return s
	}

	//否则偏移量就是大于或等于元素总个数，那么,这里不能写成：s[keyCount-1]...因为这样写，会把该数组中的最后一个元素加入进来
	s = append(s[:index], s[keyCount:]...) //当前key的位置之前的所有元素和该数组的总元素数及以后的元素，总元素数的下标及以后也就是没有元素

	return s
}

//删除一个或多个切片元素，参数1，切片，参数2：下标，参数3：要删除的个数。返回string类型切片
func ArrayDeleteKeyStringLength(s []string, index int, length int) []string {

	if length < 1 { //当要删除的个数小于1也就是等于0或负数时，就直接返回数组
		return s
	}

	keyCount := len(s)       //元素数量
	offset := index + length //当前key的下标和要删除的数量
	//当元素的总个数大于偏移量时，也就是当前要被删除的下标加上要删除的个数的和小于元素总个数，那么是可操作的。
	if keyCount > offset {
		s = append(s[:index], s[offset:]...)
		return s
	}

	//否则偏移量就是大于或等于元素总个数，那么,这里不能写成：s[keyCount-1]...因为这样写，会把该数组中的最后一个元素加入进来
	s = append(s[:index], s[keyCount:]...) //当前key的位置之前的所有元素和该数组的总元素数及以后的元素，总元素数的下标及以后也就是没有元素

	return s
}

//===========================================以保留元索引为参数得到要删除的索引=========================================================

//收集要删除的索引值。参数：1要操作的切片、2要保留的索引。返回值：要删除的索引切片。该返回值主要配合liuyang.ArrayRemoveUint32()函数使用。
//前题条件：要操作的切片中元素值不能为0，因为要保留的索引值就是以0为标记的。
func ArrayCollectDelIndexUint32(arr []uint32, saveIndex []uint32) []uint32 {
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
	tempPotArr := make([]uint32, len(arr))
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
		tempPotArr[saveIndex[i]] = 0
	}

	//遍历临时数组
	var i uint32 = 0                     //声明uint32类型的i
	var length = uint32(len(tempPotArr)) //获取临时数组长度
	for i = 0; i < length; i++ {
		//只要元素值大于0，就是要删除的索引值
		if tempPotArr[i] > 0 {
			//将要删除的索引值追加到临时删除索引收集数组中
			delArr = append(delArr, i)
		}
	}

	return delArr
}

//收集要删除的索引值。参数：1要操作的切片、2要保留的索引。返回值：要删除的索引切片。该返回值主要配合liuyang.ArrayRemoveUint32()函数使用。
//前题条件：要操作的切片中元素值不能为0，因为要保留的索引值就是以0为标记的。
func ArrayCollectDelIndexUint64(arr []uint64, saveIndex []uint32) []uint32 {
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
	tempPotArr := make([]uint64, len(arr))
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
		tempPotArr[saveIndex[i]] = 0
	}

	//遍历临时数组
	var i uint32 = 0                     //声明uint32类型的i
	var length = uint32(len(tempPotArr)) //获取临时数组长度
	for i = 0; i < length; i++ {
		//只要元素值大于0，就是要删除的索引值
		if tempPotArr[i] > 0 {
			//将要删除的索引值追加到临时删除索引收集数组中
			delArr = append(delArr, i)
		}
	}

	return delArr
}

//收集要删除的索引值。参数：1要操作的切片、2要保留的索引。返回值：要删除的索引切片。该返回值主要配合liuyang.ArrayRemoveUint32()函数使用。
//前题条件：要操作的切片中元素值不能为0，因为要保留的索引值就是以0为标记的。
func ArrayCollectDelIndexInt32(arr []int32, saveIndex []uint32) []uint32 {
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
	tempPotArr := make([]int32, len(arr))
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
		tempPotArr[saveIndex[i]] = 0
	}

	//遍历临时数组
	var i uint32 = 0                     //声明uint32类型的i
	var length = uint32(len(tempPotArr)) //获取临时数组长度
	for i = 0; i < length; i++ {
		//只要元素值大于0，就是要删除的索引值
		if tempPotArr[i] > 0 {
			//将要删除的索引值追加到临时删除索引收集数组中
			delArr = append(delArr, i)
		}
	}

	return delArr
}

//收集要删除的索引值。参数：1要操作的切片、2要保留的索引。返回值：要删除的索引切片。该返回值主要配合liuyang.ArrayRemoveUint32()函数使用。
//前题条件：要操作的切片中元素值不能为0，因为要保留的索引值就是以0为标记的。
func ArrayCollectDelIndexInt(arr []int, saveIndex []uint32) []uint32 {
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
	tempPotArr := make([]int, len(arr))
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
		tempPotArr[saveIndex[i]] = 0
	}

	//遍历临时数组
	var i uint32 = 0                     //声明uint32类型的i
	var length = uint32(len(tempPotArr)) //获取临时数组长度
	for i = 0; i < length; i++ {
		//只要元素值大于0，就是要删除的索引值
		if tempPotArr[i] > 0 {
			//将要删除的索引值追加到临时删除索引收集数组中
			delArr = append(delArr, i)
		}
	}

	return delArr
}

//收集要删除的索引值。参数：1要操作的切片、2要保留的索引。返回值：要删除的索引切片。该返回值主要配合liuyang.ArrayRemoveUint32()函数使用。
//前题条件：要操作的切片中元素值不能为0，因为要保留的索引值就是以0为标记的。
//string类型要注意。定义切片时不要将元素值设置为""，如：[]string{"刘阳", "", "", "璐璐", "凤凤"}，这样会误删掉1、2索引。
//建议定义切片时将元素值设置为" "，带一个空格。这样就不会发生误删除的情况了。
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

//===========================================批量删除切片索引=========================================================

//批量删除切片索引。参数：1要操作的切片、2要删除的索引。返回值：已批量删除索引后的切片。uint32类型
func ArrayRemoveUint32(arr []uint32, delIndex []uint32) []uint32 {
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

//批量删除切片索引。参数：1要操作的切片、2要删除的索引。返回值：已批量删除索引后的切片。uint64类型
func ArrayRemoveUint64(arr []uint64, delIndex []uint32) []uint64 {
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

//批量删除切片索引。参数：1要操作的切片、2要删除的索引。返回值：已批量删除索引后的切片。int类型
func ArrayRemoveInt(arr []int, delIndex []uint32) []int {
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

//批量删除切片索引。参数：1要操作的切片、2要删除的索引。返回值：已批量删除索引后的切片。int32类型
func ArrayRemoveInt32(arr []int32, delIndex []uint32) []int32 {
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

//批量删除切片索引。参数：1要操作的切片、2要删除的索引。返回值：已批量删除索引后的切片。string类型
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

//删除数组切片======================================

//复制数组切片======================================

//复制一个byte数组切片
func ArrayCopyByte(arr []byte) []byte {
	var temp = make([]byte, len(arr))
	//将原数据数组拷贝给临时数组。拷贝后，返回元素数量
	copy(temp, arr)

	return temp
}

//复制一个bool数组切片
func ArrayCopyBool(arr []bool) []bool {
	var temp = make([]bool, len(arr))
	//将原数据数组拷贝给临时数组。拷贝后，返回元素数量
	copy(temp, arr)

	return temp
}

//复制一个uint数组切片
func ArrayCopyUint(arr []uint) []uint {
	var temp = make([]uint, len(arr))
	//将原数据数组拷贝给临时数组。拷贝后，返回元素数量
	copy(temp, arr)

	return temp
}

//复制一个uint8数组切片
func ArrayCopyUint8(arr []uint8) []uint8 {
	var temp = make([]uint8, len(arr))
	//将原数据数组拷贝给临时数组。拷贝后，返回元素数量
	copy(temp, arr)

	return temp
}

//复制一个uint16数组切片
func ArrayCopyUint16(arr []uint16) []uint16 {
	var temp = make([]uint16, len(arr))
	//将原数据数组拷贝给临时数组。拷贝后，返回元素数量
	copy(temp, arr)

	return temp
}

//复制一个uint32数组切片
func ArrayCopyUint32(arr []uint32) []uint32 {
	var temp = make([]uint32, len(arr))
	//将原数据数组拷贝给临时数组。拷贝后，返回元素数量
	copy(temp, arr)

	return temp
}

//复制一个uint64数组切片
func ArrayCopyUint64(arr []uint64) []uint64 {
	var temp = make([]uint64, len(arr))
	//将原数据数组拷贝给临时数组。拷贝后，返回元素数量
	copy(temp, arr)

	return temp
}

//复制一个int数组切片
func ArrayCopyInt(arr []int) []int {
	var temp = make([]int, len(arr))
	//将原数据数组拷贝给临时数组。拷贝后，返回元素数量
	copy(temp, arr)

	return temp
}

//复制一个int8数组切片
func ArrayCopyInt8(arr []int8) []int8 {
	var temp = make([]int8, len(arr))
	//将原数据数组拷贝给临时数组。拷贝后，返回元素数量
	copy(temp, arr)

	return temp
}

//复制一个int16数组切片
func ArrayCopyInt16(arr []int16) []int16 {
	var temp = make([]int16, len(arr))
	//将原数据数组拷贝给临时数组。拷贝后，返回元素数量
	copy(temp, arr)

	return temp
}

//复制一个int32数组切片
func ArrayCopyInt32(arr []int32) []int32 {
	var temp = make([]int32, len(arr))
	//将原数据数组拷贝给临时数组。拷贝后，返回元素数量
	copy(temp, arr)

	return temp
}

//复制一个int64数组切片
func ArrayCopyInt64(arr []int64) []int64 {
	var temp = make([]int64, len(arr))
	//将原数据数组拷贝给临时数组。拷贝后，返回元素数量
	copy(temp, arr)

	return temp
}

//复制一个string数组切片
func ArrayCopyString(arr []string) []string {
	var temp = make([]string, len(arr))
	//将原数据数组拷贝给临时数组。拷贝后，返回元素数量
	copy(temp, arr)

	return temp
}

//复制一个切片范围，参数1，原数组，参数2，起始位置，参数3，复制的长度。返回复制后的新数组，int类型
func ArrayCopyRangeInt(arr []int, start int, length int) []int {

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
		var temp = make([]int, offset)
		//将原数据数组拷贝给临时数组。其起始位置为start，结束位置则是原数组的总长度。
		copy(temp, arr[start:count]) //返回值为临时数组的长度

		return temp
	}

	//常规情况下
	//差值 = 起始位置 + 要复制的长度
	offset = start + length
	//将复制的长度设置为临时数组的长度
	var temp = make([]int, length)
	//将原数据数组拷贝给临时数组。起始位置为start，结束位置则是起始位置加上要复制的长度
	copy(temp, arr[start:offset])

	return temp
}

//复制一个切片范围，参数1，原数组，参数2，起始位置，参数3，复制的长度。返回复制后的新数组，string类型
func ArrayCopyRangeString(arr []string, start int, length int) []string {

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
		var temp = make([]string, offset)
		//将原数据数组拷贝给临时数组。其起始位置为start，结束位置则是原数组的总长度。
		copy(temp, arr[start:count]) //返回值为临时数组的长度

		return temp
	}

	//常规情况下
	//差值 = 起始位置 + 要复制的长度
	offset = start + length
	//将复制的长度设置为临时数组的长度
	var temp = make([]string, length)
	//将原数据数组拷贝给临时数组。起始位置为start，结束位置则是起始位置加上要复制的长度
	copy(temp, arr[start:offset])

	return temp
}

//复制数组切片======================================
//将所有类型的数组分割成以指定符号连接的字符串。该函数简单暴力，和下面ArrayImplode()函数的功能是一样的。
func ArrayToString(arr interface{}, sep string) string {
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

//获取一个由int类型切片元素转成以指定符号连接的字符串。该方法学习自2010年长春好萝卜公司的同事冷志伟先生。
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

//获取一个由int32类型切片元素转成以指定符号连接的字符串。
func ArrayImplodeInt32(arr []int32, sep string) string {
	//临时变量
	var str string
	for i := 0; i < len(arr); i++ {
		//如果当前i的值小于元素总个数-1时，也就是最后一个key之前的值时，后面可以连接逗号
		if i < len(arr)-1 {
			str += Int32ToString(arr[i]) + sep
		} else {
			//否则i的值等于最后一个key时，后面就不要连接逗号了
			str += Int32ToString(arr[i])
		}
	}

	return str
}

//获取一个由uint32类型切片元素转成以指定符号连接的字符串。
func ArrayImplodeUInt32(arr []uint32, sep string) string {
	//临时变量
	var str string
	for i := 0; i < len(arr); i++ {
		//如果当前i的值小于元素总个数-1时，也就是最后一个key之前的值时，后面可以连接逗号
		if i < len(arr)-1 {
			str += UInt32ToString(arr[i]) + sep
		} else {
			//否则i的值等于最后一个key时，后面就不要连接逗号了
			str += UInt32ToString(arr[i])
		}
	}

	return str
}

//获取一个由int64类型切片元素转成以指定符号连接的字符串。
func ArrayImplodeInt64(arr []int64, sep string) string {
	//临时变量
	var str string
	for i := 0; i < len(arr); i++ {
		//如果当前i的值小于元素总个数-1时，也就是最后一个key之前的值时，后面可以连接逗号
		if i < len(arr)-1 {
			str += Int64ToString(arr[i]) + sep
		} else {
			//否则i的值等于最后一个key时，后面就不要连接逗号了
			str += Int64ToString(arr[i])
		}
	}

	return str
}

//获取一个由uint64类型切片元素转成以指定符号连接的字符串。
func ArrayImplodeUInt64(arr []uint64, sep string) string {
	//临时变量
	var str string
	for i := 0; i < len(arr); i++ {
		//如果当前i的值小于元素总个数-1时，也就是最后一个key之前的值时，后面可以连接逗号
		if i < len(arr)-1 {
			str += UInt64ToString(arr[i]) + sep
		} else {
			//否则i的值等于最后一个key时，后面就不要连接逗号了
			str += UInt64ToString(arr[i])
		}
	}

	return str
}

//切片中相同元素的数量统计(byte类型)。返回：元素值与元素出现次数
func ArraySameElementShowNum(arr []byte) map[byte]byte {
	//校验长度
	if len(arr) < 1 {
		return nil
	}
	//对象键值对法
	//该方法执行的速度比其他任何方法都快，就是占用的内存大一些
	tempMap := make(map[byte]byte, 0)

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

//切片中相同元素的数量统计(int类型)。返回：元素值与元素出现次数
func ArraySameElementShowNumInt(arr []int) map[int]int {
	//校验长度
	if len(arr) < 1 {
		return nil
	}
	//对象键值对法
	//该方法执行的速度比其他任何方法都快，就是占用的内存大一些
	tempMap := make(map[int]int, 0)

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

//切片元素去重(byte类型)。返回：去重后的切片
func ArrayElementSingle(arr []byte) []byte {
	//校验长度，如果长度为0或1时，直接返回原切片。即使长度为1时也不用去重。
	if len(arr) < 2 {
		return arr
	}

	//对象键值对法
	//该方法执行的速度比其他任何方法都快，就是占用的内存大一些
	tempMap := make(map[byte]byte, 0)

	//将所有元素设置为临时map的key
	for _, value := range arr {
		tempMap[value] = 1
	}

	//fmt.Println("=======================遍历=======================")

	//临时切片
	tempArr := make([]byte, 0)
	//遍历临时map，取其key
	for key, _ := range tempMap {
		//收集去重后的元素
		tempArr = append(tempArr, key)
	}

	return tempArr
}

//切片元素去重(int类型)。返回：去重后的切片
func ArrayElementSingleInt(arr []int) []int {
	//校验长度，如果长度为0或1时，直接返回原切片。即使长度为1时也不用去重。
	if len(arr) < 2 {
		return arr
	}

	//对象键值对法
	//该方法执行的速度比其他任何方法都快，就是占用的内存大一些
	tempMap := make(map[int]int, 0)

	//将所有元素设置为临时map的key
	for _, value := range arr {
		tempMap[value] = 1
	}

	//fmt.Println("=======================遍历=======================")

	//临时切片
	tempArr := make([]int, 0)
	//遍历临时map，取其key
	for key, _ := range tempMap {
		//收集去重后的元素
		tempArr = append(tempArr, key)
	}

	return tempArr
}

//检测元素是否在切片中(int类型)。返回：true存在、false不存在
func CheckInArray(item int, arr []int) bool {
	//临时变量，用于标记找寻状态
	var flag_ok = false
	//遍历切片，逐一比对元素
	for _, value := range arr {
		//如果切片中有相同元素时，及时停止遍历
		if value == item {
			flag_ok = true //标记已找到
			break          //及时停止
		}
	}

	return flag_ok
}

//检测元素是否在切片中(uint类型)。返回：true存在、false不存在
func CheckInArrayUint(item uint, arr []uint) bool {
	//临时变量，用于标记找寻状态
	var flag_ok = false
	//遍历切片，逐一比对元素
	for _, value := range arr {
		//如果切片中有相同元素时，及时停止遍历
		if value == item {
			flag_ok = true //标记已找到
			break          //及时停止
		}
	}

	return flag_ok
}

//检测元素是否在切片中(int8类型)。返回：true存在、false不存在
func CheckInArrayInt8(item int8, arr []int8) bool {
	//临时变量，用于标记找寻状态
	var flag_ok = false
	//遍历切片，逐一比对元素
	for _, value := range arr {
		//如果切片中有相同元素时，及时停止遍历
		if value == item {
			flag_ok = true //标记已找到
			break          //及时停止
		}
	}

	return flag_ok
}

//检测元素是否在切片中(uint8类型)。返回：true存在、false不存在
func CheckInArrayUint8(item uint8, arr []uint8) bool {
	//临时变量，用于标记找寻状态
	var flag_ok = false
	//遍历切片，逐一比对元素
	for _, value := range arr {
		//如果切片中有相同元素时，及时停止遍历
		if value == item {
			flag_ok = true //标记已找到
			break          //及时停止
		}
	}

	return flag_ok
}

//检测元素是否在切片中(int16类型)。返回：true存在、false不存在
func CheckInArrayInt16(item int16, arr []int16) bool {
	//临时变量，用于标记找寻状态
	var flag_ok = false
	//遍历切片，逐一比对元素
	for _, value := range arr {
		//如果切片中有相同元素时，及时停止遍历
		if value == item {
			flag_ok = true //标记已找到
			break          //及时停止
		}
	}

	return flag_ok
}

//检测元素是否在切片中(uint16类型)。返回：true存在、false不存在
func CheckInArrayUint16(item uint16, arr []uint16) bool {
	//临时变量，用于标记找寻状态
	var flag_ok = false
	//遍历切片，逐一比对元素
	for _, value := range arr {
		//如果切片中有相同元素时，及时停止遍历
		if value == item {
			flag_ok = true //标记已找到
			break          //及时停止
		}
	}

	return flag_ok
}

//检测元素是否在切片中(int32类型)。返回：true存在、false不存在
func CheckInArrayInt32(item int32, arr []int32) bool {
	//临时变量，用于标记找寻状态
	var flag_ok = false
	//遍历切片，逐一比对元素
	for _, value := range arr {
		//如果切片中有相同元素时，及时停止遍历
		if value == item {
			flag_ok = true //标记已找到
			break          //及时停止
		}
	}

	return flag_ok
}

//检测元素是否在切片中(uint32类型)。返回：true存在、false不存在
func CheckInArrayUint32(item uint32, arr []uint32) bool {
	//临时变量，用于标记找寻状态
	var flag_ok = false
	//遍历切片，逐一比对元素
	for _, value := range arr {
		//如果切片中有相同元素时，及时停止遍历
		if value == item {
			flag_ok = true //标记已找到
			break          //及时停止
		}
	}

	return flag_ok
}

//检测元素是否在切片中(int64类型)。返回：true存在、false不存在
func CheckInArrayInt64(item int64, arr []int64) bool {
	//临时变量，用于标记找寻状态
	var flag_ok = false
	//遍历切片，逐一比对元素
	for _, value := range arr {
		//如果切片中有相同元素时，及时停止遍历
		if value == item {
			flag_ok = true //标记已找到
			break          //及时停止
		}
	}

	return flag_ok
}

//检测元素是否在切片中(uint64类型)。返回：true存在、false不存在
func CheckInArrayUint64(item uint64, arr []uint64) bool {
	//临时变量，用于标记找寻状态
	var flag_ok = false
	//遍历切片，逐一比对元素
	for _, value := range arr {
		//如果切片中有相同元素时，及时停止遍历
		if value == item {
			flag_ok = true //标记已找到
			break          //及时停止
		}
	}

	return flag_ok
}

//检测元素是否在切片中(string类型)。返回：true存在、false不存在
func CheckInArrayString(item string, arr []string) bool {
	//临时变量，用于标记找寻状态
	var flag_ok = false
	//遍历切片，逐一比对元素
	for _, value := range arr {
		//如果切片中有相同元素时，及时停止遍历
		if value == item {
			flag_ok = true //标记已找到
			break          //及时停止
		}
	}

	return flag_ok
}

//注意：float类型不建议使用本方法，因为会涉及到丢精度问题
//检测元素是否在切片中(float32类型)。返回：true存在、false不存在
func CheckInArrayFloat32(item float32, arr []float32) bool {
	//临时变量，用于标记找寻状态
	var flag_ok = false
	//遍历切片，逐一比对元素
	for _, value := range arr {
		//如果切片中有相同元素时，及时停止遍历
		if value == item {
			flag_ok = true //标记已找到
			break          //及时停止
		}
	}

	return flag_ok
}

//注意：float类型不建议使用本方法，因为会涉及到丢精度问题
//检测元素是否在切片中(float64类型)。返回：true存在、false不存在
func CheckInArrayFloat64(item float64, arr []float64) bool {
	//临时变量，用于标记找寻状态
	var flag_ok = false
	//遍历切片，逐一比对元素
	for _, value := range arr {
		//如果切片中有相同元素时，及时停止遍历
		if value == item {
			flag_ok = true //标记已找到
			break          //及时停止
		}
	}

	return flag_ok
}

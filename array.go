package liuyang

//数组元素排序======================================

//标准冒泡排序，参数1为原数组，参数2如果为0则正序，为1则倒序。
func ArraySortInt(arr []int, sort uint8) {
	if len(arr) < 2 || sort < 0 || sort > 1 {
		return //这里是直接返回
	}

	if sort == 0 {
		temp := 0
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr) - (1 + i); j++ {
				if arr[j] > arr[j + 1] {
					temp = arr[j]
					arr[j] = arr[j+1]
					arr[j+1] = temp
					//arr[j],arr[j+1] = arr[j+1],arr[j]
				}
			}
		}
	} else if sort == 1 {
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr) - (1 + i); j++ {
				if arr[j] < arr[j + 1] {
					arr[j],arr[j+1] = arr[j+1],arr[j]
				}
			}
		}
	}

	//return arr
}

//英文字符串排序，使用冒泡排序，在比较两个字符串的大小时，使用刘阳封装的方法StringCompareInt进行比较。
func ArraySortString(arr []string, sort uint8) []string  {
	if len(arr) < 2 {
		return arr //这里是直接返回数组
	}

	if sort == 0 {
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr) - (1 + i); j++ {
				if StringCompareInt(arr[j],arr[j+1]) > 0 {
					arr[j],arr[j+1] = arr[j+1],arr[j]
				}
			}
		}
	} else if sort == 1 {
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr) - (1 + i); j++ {
				if StringCompareInt(arr[j],arr[j+1]) < 0 {
					arr[j],arr[j+1] = arr[j+1],arr[j]
				}
			}
		}
	}
	return arr
}
//数组元素排序======================================


//删除数组切片======================================

//删除一个切片元素，返回int类型切片
func ArrayDeleteKeyInt(s []int, index int) []int {

	//原数据：[3 2 0 8 9 4 5 7]
	//fmt.Println("切片的长度为：", len(s))//8
	//fmt.Println("前置元素：", s[:index])//[3 2 0 8]
	//fmt.Println("后置元素：", s[index + 1:])//[4 5 7]
	//将下标4，也就是为9的元素删除，其操作，只是覆盖而已
	s = append(s[:index], s[index + 1:]...)

	return s
}

//删除一个切片元素，返回uint32类型切片
func ArrayDeleteKeyUInt32(s []uint32, index int) []uint32 {

	s = append(s[:index], s[index + 1:]...)

	return s
}

//删除一个切片元素，返回uint64类型切片
func ArrayDeleteKeyUInt64(s []uint64, index int) []uint64 {

	s = append(s[:index], s[index + 1:]...)

	return s
}


//删除一个切片元素，返回string类型切片
func ArrayDeleteKeyString(s []string, index int) []string {

	s = append(s[:index], s[index + 1:]...)

	return s
}

//删除一个或多个切片元素，参数1，切片，参数2：下标，参数3：要删除的个数。返回int类型切片
func ArrayDeleteKeyIntLength(s []int, index int,length int) []int {

	if length < 1 {//当要删除的个数小于1也就是等于0或负数时，就直接返回数组
		return s
	}

	keyCount := len(s)//元素数量
	offset := index + length//当前key的下标和要删除的数量
	//当元素的总个数大于偏移量时，也就是当前要被删除的下标加上要删除的个数的和小于元素总个数，那么是可操作的。
	if (keyCount > offset) {
		s = append(s[:index],s[offset:]...)
		return s
	}

	//否则偏移量就是大于或等于元素总个数，那么,这里不能写成：s[keyCount-1]...因为这样写，会把该数组中的最后一个元素加入进来
	s = append(s[:index],s[keyCount:]...)//当前key的位置之前的所有元素和该数组的总元素数及以后的元素，总元素数的下标及以后也就是没有元素

	return s
}

//删除一个或多个切片元素，参数1，切片，参数2：下标，参数3：要删除的个数。返回string类型切片
func ArrayDeleteKeyStringLength(s []string, index int,length int) []string {

	if length < 1 {//当要删除的个数小于1也就是等于0或负数时，就直接返回数组
		return s
	}

	keyCount := len(s)//元素数量
	offset := index + length//当前key的下标和要删除的数量
	//当元素的总个数大于偏移量时，也就是当前要被删除的下标加上要删除的个数的和小于元素总个数，那么是可操作的。
	if (keyCount > offset) {
		s = append(s[:index],s[offset:]...)
		return s
	}

	//否则偏移量就是大于或等于元素总个数，那么,这里不能写成：s[keyCount-1]...因为这样写，会把该数组中的最后一个元素加入进来
	s = append(s[:index],s[keyCount:]...)//当前key的位置之前的所有元素和该数组的总元素数及以后的元素，总元素数的下标及以后也就是没有元素

	return s
}


//删除数组切片======================================




//复制数组切片======================================

//复制一个int数组切片
func ArrayCopyInt(arr []int) []int {
	var temp = make([]int,len(arr))
	//将原数据数组拷贝给临时数组。拷贝后，返回元素数量
	copy(temp,arr)

	return temp
}
//复制一个string数组切片
func ArrayCopyString(arr []string) []string {
	var temp = make([]string,len(arr))
	//将原数据数组拷贝给临时数组。拷贝后，返回元素数量
	copy(temp,arr)

	return temp
}

//复制一个切片范围，参数1，原数组，参数2，起始位置，参数3，复制的长度。返回复制后的新数组，int类型
func ArrayCopyRangeInt(arr []int, start int, length int) []int {

	//首先要判断起始位置是否在合法范围内，如：[]int{2,3,5}//调用时用：arr,4,1，从下标4开始获取就属于非法操作
	if len(arr) < 1 || start < 0 || length < 1 {//如果数组长度小于1或下标小于0或想要获取的长度小于1则直接返回数组
		//fmt.Println("第1种情况")
		return arr
	}

	if start > (len(arr) - 1) {//如果下标位置数大于数组长度，则直接返回数组
		//fmt.Println("第2种情况")
		return arr
	}

	//差值
	offset := 0
	//判断length如果小于等于0或length大于(数组长度减起始位置)的差值时，也就是arr[2,3,5]有3个值时却想从下标1开始取6个数时：
	if length > (len(arr) - start){
		//fmt.Println("这里是length小于等于0或length大于数组长度-起始位置的差值时")
		count := len(arr)
		//新数组长度差值 = 数组元素总个数减起始位置
		offset = count - start
		//将差值设置为临时数组的长度
		var temp = make([]int,offset)
		//将原数据数组拷贝给临时数组。其起始位置为start，结束位置则是原数组的总长度。
		copy(temp,arr[start:count])//返回值为临时数组的长度

		return temp
	}

	//常规情况下
	//差值 = 起始位置 + 要复制的长度
	offset = start + length
	//将复制的长度设置为临时数组的长度
	var temp = make([]int,length)
	//将原数据数组拷贝给临时数组。起始位置为start，结束位置则是起始位置加上要复制的长度
	copy(temp,arr[start:offset])

	return temp
}

//复制一个切片范围，参数1，原数组，参数2，起始位置，参数3，复制的长度。返回复制后的新数组，string类型
func ArrayCopyRangeString(arr []string, start int, length int) []string {

	//首先要判断起始位置是否在合法范围内，如：[]int{2,3,5}//调用时用：arr,4,1，从下标4开始获取就属于非法操作
	if len(arr) < 1 || start < 0 || length < 1 {//如果数组长度小于1或下标小于0或想要获取的长度小于1则直接返回数组
		//fmt.Println("第1种情况")
		return arr
	}

	if start > (len(arr) - 1) {//如果下标位置数大于数组长度，则直接返回数组
		//fmt.Println("第2种情况")
		return arr
	}

	//差值
	offset := 0
	//判断length如果小于等于0或length大于(数组长度减起始位置)的差值时，也就是arr[2,3,5]有3个值时却想从下标1开始取6个数时：
	if length > (len(arr) - start){
		//fmt.Println("这里是length小于等于0或length大于数组长度-起始位置的差值时")
		count := len(arr)
		//新数组长度差值 = 数组元素总个数减起始位置
		offset = count - start
		//将差值设置为临时数组的长度
		var temp = make([]string,offset)
		//将原数据数组拷贝给临时数组。其起始位置为start，结束位置则是原数组的总长度。
		copy(temp,arr[start:count])//返回值为临时数组的长度

		return temp
	}

	//常规情况下
	//差值 = 起始位置 + 要复制的长度
	offset = start + length
	//将复制的长度设置为临时数组的长度
	var temp = make([]string,length)
	//将原数据数组拷贝给临时数组。起始位置为start，结束位置则是起始位置加上要复制的长度
	copy(temp,arr[start:offset])

	return temp
}
//复制数组切片======================================

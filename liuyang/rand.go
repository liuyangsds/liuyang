package liuyang

import (
	"math/rand"
	"time"
)

//生成随机数======================================

//生成一个随机数，得到大于等于0并且小于参数值max的随机数
func RandomNumber(max int) int {
	if max <= 0 {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	randX := rand.Intn(max)

	//另一种写法：
	//首先生成一个纳米时间戳，作为随机的种子
	//nanoTime := time.Now().UnixNano()
	//返加一个像切片类型的数据，&{0 334 [-7237584939349062640 -4633729966338938091 -6634801558180662043....此处省略若干]}
	//sliceSource := rand.NewSource(nanoTime)
	//以随机种子重新创建了rand函数，此时返回了一个新的rand的指针
	//r := rand.New(sliceSource)

	return randX
}
//生成一个随机数，得到大于等于参数值min并且小于参数值max的随机数
func RandomNumberRange(min int, max int) int {
	if max <= 0 {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	//差值 = 最大数 - 最小数
	offset := max - min
	randX := rand.Intn(offset)

	//区间数 = 最小数 + 随机(差值)
	return randX + min
}

//int类型数组元素顺序打乱，以向临时数组添加新元素再删除原数组元素，反复操作，一般
func RandomArrayInt(arr []int) []int {
	//将传递过来的指针数组拷贝一个新数组，此时不应该在这里拷贝，而应该把是否要复制一份新数组的权力给调用者
	//copyArr := copyArrayInt(arr)
	rand.Seed(time.Now().UnixNano())
	var tempArr []int
	//注意，下面的10代表原数组中固定的元素值，也就是有多少个元素，就要循环多少次，一定要记住。
	num := len(arr)
	for i := 0; i < num; i++ {

		randX := rand.Intn(len(arr))
		//fmt.Print(randX,"\t")//4	8	2	2	1	2	0	1	1	0
		tempArr = append(tempArr,arr[randX])
		//删除生成的随机数所对应的值
		arr = SliceDeleteKeyInt(arr,randX)
		//fmt.Println(len(copyArr),"===>",randX)
	}
	//fmt.Println(tempArr)//[104 109 102 103 101 106 100 107 108 105]

	return tempArr
}

//string类型数组元素顺序打乱，以向临时数组添加新元素再删除原数组元素，反复操作，效率一般
func RandomArrayString(arr []string) []string {
	//将传递过来的指针数组拷贝一个新数组，此时不应该在这里拷贝，而应该把是否要复制一份新数组的权力给调用者
	//copyArr := copyArrayString(arr)
	rand.Seed(time.Now().UnixNano())
	var tempArr []string
	//注意，下面的10代表原数组中固定的元素值，也就是有多少个元素，就要循环多少次，一定要记住。
	num := len(arr)
	for i := 0; i < num; i++ {

		randX := rand.Intn(len(arr))
		//fmt.Print(randX,"\t")//4	8	2	2	1	2	0	1	1	0
		tempArr = append(tempArr,arr[randX])
		//删除生成的随机数所对应的值
		arr = SliceDeleteKeyString(arr,randX)
		//fmt.Println(len(copyArr),"===>",randX)
	}
	//fmt.Println(tempArr)//[104 109 102 103 101 106 100 107 108 105]

	return tempArr
}

//int类型数组元素打乱之最高境界，效率高，推荐
func RandomArrayIntSuper(arr []int) []int {
	//生成时间种子
	rand.Seed(time.Now().UnixNano())
	var temp int
	for i := 0; i < len(arr); i++ {
		//第1步，得到随机数，范围值一定要大于0
		randX := rand.Intn(i+1)
		//第2步，当前位置的值赋值临时变量，这样等于备份了一下，因为即将有随机位置的值赋过来并替换掉当前位置的值
		//temp变量之所以需要，是因为要打乱一个数组中的元素。这个数组元素有可能是200、300、555、666这样的值。
		//如果只是想要将0到10或0到100之内的数随机，那可以省略temp变量。
		temp = arr[i]
		//第3步，随机位置赋值给当前位置
		arr[i] = arr[randX]
		//第4步，将刚刚暂存的值赋给随机位置
		arr[randX] = temp
	}

	return arr
}
//string类型数组元素打乱之最高境界，效率高，推荐
func RandomArrayStringSuper(arr []string) []string {
	//生成时间种子
	rand.Seed(time.Now().UnixNano())
	var temp string
	for i := 0; i < len(arr); i++ {
		randX := rand.Intn(i+1)
		temp = arr[i]
		arr[i] = arr[randX]
		arr[randX] = temp
	}

	return arr
}

//人员位置随机就坐法：
//生成int类型数组，其元素为大于等于0并且小于参数值的随机数的数组
//如参数为5，则会得到[2 4 0 1 3]这样的数组
func RandomNumberArrayInt(max int) []int {
	//设置随机种子
	rand.Seed(time.Now().UnixNano())
	//声明临时切片数组
	arr := make([]int,max)

	for i := 0; i < max; i++ {
		//第1步，得到随机数，范围值一定要大于0
		randX := rand.Intn(i+1)
		//第2步，随机位置赋值给当前位置
		arr[i] = arr[randX]
		//第3步，当前值赋值给随机位置
		arr[randX] = i
	}

	return arr
}

//生成int类型数组，其元素值为大于等于参数1并且小于参数2的随机数的数组
//如参数为5，10，则会得到[6 5 8 7 9]这样的数组
func RandomNumberArrayIntRange(min int, max int) []int {
	//得到差值，计算数量
	offset := max - min
	//创建数组切片，元素个数为差值数量
	arr := make([]int,offset)
	//生成时间种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < offset; i++ {
		randX := rand.Intn(i+1)//i+1很重要，因为每次要获取小于这个数的值，i初始为0，总不能获取小于0的数。
		arr[i] = arr[randX]
		arr[randX] = i + min
	}

	return arr
}

//生成int类型数组，其元素值为大于等于参数1并且小于参数2、元素个数为参数3的随机数的数组
//如参数为5，10，3，则会得到[9 8 5]这样的数组
func RandomNumberArrayIntRangeNum(min int, max int,num int) []int {
	//先调用：获取一个指定区间不重复的随机数数组
	arr := RandomNumberArrayIntRange(min, max)
	////再调用：拷贝一个区间元素的数组
	//temp := ArrayCopyRangeInt(arr,0, num)
	//return temp

	//在新生成的数组中取元素个数的切片即可，如：
	return arr[0:num]
}
//生成随机数======================================

//生成随机字符串，参数为要生成的字符串长度
func RandomString(n int) string {
	arr := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	temp := make([]byte,n)

	rand.Seed(time.Now().UnixNano())//设置随机种子
	for key := range temp {
		//fmt.Println(key,"=====",value)
		temp[key] = arr[rand.Intn(len(arr))]
	}
	//fmt.Println(string(temp))//myzhmslxhr

	return string(temp)
}

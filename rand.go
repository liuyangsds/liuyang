package liuyang

import (
	"math/rand"
)

//生成随机数======================================
//生成一个随机数，得到大于等于0并且小于参数值max的随机数
func RandomNumber(max int) int {
	if max <= 0 {
		return max
	}
	//随机种子必须要有，并且要用公用的，不能写在方法里，不然不利于循环
	//rand.Seed(time.Now().UnixNano())
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

//生成一个随机数，得到大于等于参数值min并且小于参数值max的随机数(不包含最大值)
func RandomNumberRange(min int, max int) int {
	if max <= 0 {
		return 0
	}
	//随机种子必须要有，并且要用公用的，不能写在方法里，不然不利于循环
	//rand.Seed(time.Now().UnixNano())
	//差值 = 最大数 - 最小数
	offset := max - min
	randX := rand.Intn(offset)

	//区间数 = 最小数 + 随机(差值)
	return randX + min
}

//生成一个随机数，得到大于等于参数值min并且小于等于参数值max的随机数(包含最大值)
func RandomNumberRangeContain(min int, max int) int {
	if max <= 0 {
		return 0
	}
	//随机种子必须要有，并且要用公用的，不能写在方法里，不然不利于循环
	//rand.Seed(time.Now().UnixNano())
	//差值 = 最大数 - 最小数
	offset := max + 1 - min
	randX := rand.Intn(offset)

	//区间数 = 最小数 + 随机(差值)
	return randX + min
}

//int类型数组元素顺序打乱，以向临时数组添加新元素再删除原数组元素，反复操作，效率一般，参考学习
func RandomArrayIntNormal(arr []int) []int {
	//随机种子必须要有，只能用公用的，不能写在方法里，不然不利于循环
	//rand.Seed(time.Now().UnixNano())
	var tempArr []int
	//注意，下面的10代表原数组中固定的元素值，也就是有多少个元素，就要循环多少次，一定要记住。
	num := len(arr)
	for i := 0; i < num; i++ {

		randX := rand.Intn(len(arr))
		//fmt.Print(randX,"\t")//4	8	2	2	1	2	0	1	1	0
		tempArr = append(tempArr, arr[randX])
		//删除生成的随机数所对应的值
		arr = ArrayDeleteKeyInt(arr, randX)
		//fmt.Println(len(copyArr),"===>",randX)
	}
	//fmt.Println(tempArr)//[104 109 102 103 101 106 100 107 108 105]

	return tempArr
}

//string类型数组元素顺序打乱，以向临时数组添加新元素再删除原数组元素，反复操作，效率一般，参考学习
func RandomArrayStringNormal(arr []string) []string {
	//随机种子必须要有，只能用公用的，不能写在方法里，不然不利于循环
	//rand.Seed(time.Now().UnixNano())
	var tempArr []string
	//注意，下面的10代表原数组中固定的元素值，也就是有多少个元素，就要循环多少次，一定要记住。
	num := len(arr)
	for i := 0; i < num; i++ {

		randX := rand.Intn(len(arr))
		//fmt.Print(randX,"\t")//4	8	2	2	1	2	0	1	1	0
		tempArr = append(tempArr, arr[randX])
		//删除生成的随机数所对应的值
		arr = ArrayDeleteKeyString(arr, randX)
		//fmt.Println(len(copyArr),"===>",randX)
	}
	//fmt.Println(tempArr)//[104 109 102 103 101 106 100 107 108 105]

	return tempArr
}

//byte类型数组元素打乱之最高境界，效率高，推荐
func RandomArrayByte(arr []byte) []byte {
	//随机种子必须要有，只能用公用的，不能写在方法里，不然不利于循环
	//rand.Seed(time.Now().UnixNano())
	var temp byte
	for i := 0; i < len(arr); i++ {
		//第1步，得到随机数，范围值一定要大于0
		randX := rand.Intn(i + 1)
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

//uint8类型数组元素打乱之最高境界，效率高，推荐
func RandomArrayUInt8(arr []uint8) []uint8 {
	//随机种子必须要有，只能用公用的，不能写在方法里，不然不利于循环
	//rand.Seed(time.Now().UnixNano())
	var temp uint8
	for i := 0; i < len(arr); i++ {
		//第1步，得到随机数，范围值一定要大于0
		randX := rand.Intn(i + 1)
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

//int类型数组元素打乱之最高境界，效率高，推荐
func RandomArrayInt(arr []int) []int {
	//随机种子必须要有，只能用公用的，不能写在方法里，不然不利于循环
	//rand.Seed(time.Now().UnixNano())
	var temp int
	for i := 0; i < len(arr); i++ {
		//第1步，得到随机数，范围值一定要大于0
		randX := rand.Intn(i + 1)
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

//uint16类型数组元素打乱之最高境界，效率高，推荐
func RandomArrayUInt16(arr []uint16) []uint16 {
	//随机种子必须要有，只能用公用的，不能写在方法里，不然不利于循环
	//rand.Seed(time.Now().UnixNano())
	var temp uint16
	for i := 0; i < len(arr); i++ {
		//第1步，得到随机数，范围值一定要大于0
		randX := rand.Intn(i + 1)
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

//uint32类型数组元素打乱之最高境界，效率高，推荐
func RandomArrayUInt32(arr []uint32) []uint32 {
	//随机种子必须要有，只能用公用的，不能写在方法里，不然不利于循环
	//rand.Seed(time.Now().UnixNano())
	var temp uint32
	for i := 0; i < len(arr); i++ {
		//第1步，得到随机数，范围值一定要大于0
		randX := rand.Intn(i + 1)
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

//int64类型数组元素打乱之最高境界，效率高，推荐
func RandomArrayInt64(arr []int64) []int64 {
	//随机种子必须要有，只能用公用的，不能写在方法里，不然不利于循环
	//rand.Seed(time.Now().UnixNano())
	var temp int64
	for i := 0; i < len(arr); i++ {
		//第1步，得到随机数，范围值一定要大于0
		randX := rand.Intn(i + 1)
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

//uint64类型数组元素打乱之最高境界，效率高，推荐
func RandomArrayUInt64(arr []uint64) []uint64 {
	//随机种子必须要有，只能用公用的，不能写在方法里，不然不利于循环
	//rand.Seed(time.Now().UnixNano())
	var temp uint64
	for i := 0; i < len(arr); i++ {
		//第1步，得到随机数，范围值一定要大于0
		randX := rand.Intn(i + 1)
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
func RandomArrayString(arr []string) []string {
	//随机种子必须要有，只能用公用的，不能写在方法里，不然不利于循环
	//rand.Seed(time.Now().UnixNano())
	var temp string
	for i := 0; i < len(arr); i++ {
		randX := rand.Intn(i + 1)
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
	//随机种子必须要有，只能用公用的，不能写在方法里，不然不利于循环
	//rand.Seed(time.Now().UnixNano())
	//声明临时切片数组
	arr := make([]int, max)

	for i := 0; i < max; i++ {
		//第1步，得到随机数，范围值一定要大于0
		randX := rand.Intn(i + 1)
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
	arr := make([]int, offset)
	//随机种子必须要有，只能用公用的，不能写在方法里，不然不利于循环
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < offset; i++ {
		randX := rand.Intn(i + 1) //i+1很重要，因为每次要获取小于这个数的值，i初始为0，总不能获取小于0的数。
		arr[i] = arr[randX]
		arr[randX] = i + min
	}

	return arr
}

//生成int类型数组，其元素值为大于等于参数1并且小于参数2、元素个数为参数3的随机数的数组，但此方式会生成指定区间值的长度数组，
//如：1、100、5。会生成长度为99的数组，只是获取前5个而已。且生成的元素是不重复的，也就是真的有99个元素被随机打乱后，取前5个。
//如：20、60、5。会生成：[57 39 56 45 23]
//如：1、7、5。会生成：[3 2 6 1 5]
//当前方法生随机数组： [1 6 5 3 4]，此方式是生成1次区间值，所以不重复。
//调用专用生成色子数组： [2 2 5 4 5]，此方式是真的生成5次随机数，因为是独立各自生成的，所以区间值重复，所以生成色子就得用色子专用方法。
//综上所述，如果是生成色子就不能用此方法，因为色子会出现多个一样的点数，所以不建议使用此方法。
func RandomNumberArrayIntRangeNum(min int, max int, num int) []int {
	//先调用：获取一个指定区间不重复的随机数数组
	arr := RandomNumberArrayIntRange(min, max)
	////再调用：拷贝一个区间元素的数组
	//temp := ArrayCopyRangeInt(arr,0, num)
	//return temp

	//在新生成的数组中取元素个数的切片即可，如：
	return arr[0:num]
}

//生成类似色子的专用函数。比如5个色子(1,7,5)。参数：最小值、最大值、生成个数。(1到6之间的随机数，不包括7)
//和另一函数RandomNumberArrayIntRangeNum()比较如下：
//生成随机数组： [1 6 5 3 4]，此方式是生成1次区间值，所以不重复。
//生成色子数组： [2 2 5 4 5]，此方式是真的生成5次随机数，因为是独立各自生成的，所以区间值重复，所以生成色子就得用色子专用方法。
func RandomDice(min int, max int, num int) ([]uint32, uint32) {
	//声明色子数组，元素个数就是要生成的个数
	diceArr := make([]uint32, num)
	//临时变量，用于累加色子值
	var sum uint32 = 0
	for i := 0; i < num; i++ {
		randX := RandomNumberRange(min, max) //生成1到6之间的随机数，不包括7。
		diceData := uint32(randX)            //将生成的点数转换类型
		sum += diceData                      //记录生成随机数的累加值，最后要返回生成指定个数随机数的总和
		diceArr[i] = diceData                //将生成的随机数赋到对应索引上
	}

	return diceArr, sum
}

//生成类似色子豹子的专用函数。比如5个色子(1,7,5)。参数：最小值、最大值、生成个数。(1到6之间的随机数且是豹子，不包括7)
//生成色子豹子： [2 2 2 2 2]
//生成色子豹子： [6 6 6 6 6]
func RandomDiceBaoZi(min int, max int, num int) ([]uint32, uint32) {
	//声明色子数组，元素个数就是要生成的个数
	diceArr := make([]uint32, num)
	//随机一个区间数
	randX := RandomNumberRange(min, max)
	//临时变量，用于累加色子值
	var sum uint32 = 0
	//转成uint32
	diceData := uint32(randX)
	for i := 0; i < num; i++ {
		sum += diceData //记录生成随机数的累加值，最后要返回生成指定个数随机数的总和
		diceArr[i] = diceData
	}

	return diceArr, sum
}

//生成随机数======================================

//生成随机字符串，参数为要生成的字符串长度
func RandomString(n int) string {
	arr := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	temp := make([]byte, n)
	//随机种子必须要有，只能用公用的，不能写在方法里，不然不利于循环
	//rand.Seed(time.Now().UnixNano())
	for key := range temp {
		//fmt.Println(key,"=====",value)
		temp[key] = arr[rand.Intn(len(arr))]
	}
	//fmt.Println(string(temp))//myzhmslxhr

	return string(temp)
}

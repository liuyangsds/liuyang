package liuyang

import (
	"math"
)

//游戏坐标相关

//获取某个索引的对称坐标偏移位置。如：(0)-385、(1)-275、(2)-165、(3)-55、(4)55、(5)165、(6)275、(7)385
//参数：总个数、每个单位的直径、单位之间的间隔、每个单位所在位置(数据集遍历时的key，0起始)
func GetKeyOffsetPosition(total int, diameter float64, space float64, key int) float64 {
	if total <= 1 {
		return 0
	}

	//要想从0开始左右对称，必须要减1才行。
	//因为中间位置的0也算1个，总数7个，实际上6个即可，加上中间的0，共7个。
	total = total - 1

	length := diameter + space             //每个单位的直径(100) + 单位之间的间隔(10) = 实际每个单位的长度(110)
	countWeight := float64(total) * length //玩家总宽度：总人数(7) * 实际每个单位的长度(110) = 总宽度(770)
	offset := countWeight / 2              //总宽度想要从中间左右对称站立就得 除以 2，得到位置偏移：385
	offsetAdd := float64(key) * length     //key * 实际每个单位的长度 = 随着遍历的玩家数量增多，偏移的坐标值也要增加
	position := -offset + offsetAdd        //想要以左侧起始，就得设置为负，所以 -385 + 实际单位个数及长度 = 位置

	//fmt.Println(key, "*", length, "=", offsetAdd, "------>", -offset, "+", offsetAdd, "=", position)
	//常规情况：
	//GetKeyOffsetPosition(7, 100, 10, i)。i为0、1、2、3、4、5、6
	//0 * 110 = 0   ------> -385 + 0   = -385
	//1 * 110 = 110 ------> -385 + 110 = -275
	//2 * 110 = 220 ------> -385 + 220 = -165
	//3 * 110 = 330 ------> -385 + 330 = -55
	//4 * 110 = 440 ------> -385 + 440 = 55
	//5 * 110 = 550 ------> -385 + 550 = 165
	//6 * 110 = 660 ------> -385 + 660 = 275

	//GetKeyOffsetPosition(4, 100, 10, i)。i为0、1、2、3
	//0 * 110 = 0 	------> -220 + 0   = -220
	//1 * 110 = 110 ------> -220 + 110 = -110
	//2 * 110 = 220 ------> -220 + 220 = 0
	//3 * 110 = 330 ------> -220 + 330 = 110

	//GetKeyOffsetPosition(3, 100, 10, i)。i为0、1、2
	//0 * 110 = 0 	------> -165 + 0   = -165
	//1 * 110 = 110 ------> -165 + 110 = -55
	//2 * 110 = 220 ------> -165 + 220 = 55

	//GetKeyOffsetPosition(2, 100, 10, i)。i为0、1
	//0 * 110 = 0 	------> -110 + 0   = -110
	//1 * 110 = 110 ------> -110 + 110 = 0

	//=========================================================================================

	//对称情况：
	//GetKeyOffsetPosition(7, 100, 10, i)。i为0、1、2、3、4、5、6
	//0 * 110 = 0 	------> -330 + 0   = -330
	//1 * 110 = 110 ------> -330 + 110 = -220
	//2 * 110 = 220 ------> -330 + 220 = -110
	//3 * 110 = 330 ------> -330 + 330 = 0
	//4 * 110 = 440 ------> -330 + 440 = 110
	//5 * 110 = 550 ------> -330 + 550 = 220
	//6 * 110 = 660 ------> -330 + 660 = 330

	//GetKeyOffsetPosition(4, 100, 10, i)。i为0、1、2、3
	//0 * 110 = 0 	------> -165 + 0   = -165
	//1 * 110 = 110 ------> -165 + 110 = -55
	//2 * 110 = 220 ------> -165 + 220 = 55
	//3 * 110 = 330 ------> -165 + 330 = 165

	//GetKeyOffsetPosition(3, 100, 10, i)。i为0、1、2
	//0 * 110 = 0 ------> -110 + 0 = -110
	//1 * 110 = 110 ------> -110 + 110 = 0
	//2 * 110 = 220 ------> -110 + 220 = 110

	//GetKeyOffsetPosition(2, 100, 10, i)。i为0、1
	//0 * 110 = 0 ------> -55 + 0 = -55
	//1 * 110 = 110 ------> -55 + 110 = 55

	return position
}

//获取某个索引所在对称坐标位置。
//第一行：(0)-110,-110	(1)--->0,-110	(2)--->110,-110
//第二行：(3)-110,0		(4)--->0,0		(5)--->110,0
//第三行：(6)-110,110		(7)--->0,110	(8)--->110,110
//参数：总个数、每个单位的直径、单位之间的间隔、每个单位所在位置(数据集遍历时的key，0起始)、每行个数(小于1时取总个数)、是否完美站位(默认false)
func GetKeyCoorPosition(total int, diameter float64, space float64, key int, rowNum int, perfect bool) (float64, float64) {
	if total <= 1 {
		return 0, 0
	}

	if rowNum < 1 {
		rowNum = total
	}

	offset := float64(total) / float64(rowNum)
	colNum := int(math.Ceil(offset))

	//第1行：0、第2行：1、第3行：2
	rowKey := key / rowNum //注意，程序除0会报错，所以上面要判断为0的时候赋总数的值
	//第1列：0、第2列：1、第3列：2
	colKey := key % rowNum //注意，程序除0会报错，所以上面要判断为0的时候赋总数的值
	x := GetKeyOffsetPosition(rowNum, diameter, space, colKey)
	y := GetKeyOffsetPosition(colNum, diameter, space, rowKey)

	//如参数：GetKeyCoorPosition(9, 100, 10, i, 3, false)。i为0、1、2、3、4、5、6、7、8
	//第一行：-110,-110	--->0,-110	--->110,-110
	//第二行：-110,0		--->0,0		--->110,0
	//第三行：-110,110	--->0,110	--->110,110

	//如参数：GetKeyCoorPosition(2, 100, 10, i, 0, false)。i为0、1
	//-55,0	--->55,0

	//如果是追求完美时
	if perfect == true {
		//以总个数 模 每行个数 = 多余的个数
		surplus := total % rowNum
		//如果模每行个数大于0时，说明每行排完后，有多余的元素
		if surplus > 0 {
			//必须要有此判断，防止没到末尾就进行以下逻辑
			//注意：key是比总数小1的，总数7，key为6时就是最后一个值了。所以要判断大于等于才行。
			if key >= total-surplus {

				//fmt.Println("多余个数：", surplus, "，当前key：", key, "，colKey：", colKey)
				//多余1个时：
				//多余个数： 1 ，当前key： 6 ，colKey： 0
				//多余2个时：
				//多余个数： 2 ，当前key： 6 ，colKey： 0
				//多余个数： 2 ，当前key： 7 ，colKey： 1
				x = GetKeyOffsetPosition(surplus, diameter, space, colKey)
				y = GetKeyOffsetPosition(colNum, diameter, space, rowKey)

				//默认坐标站位样例：
				//总数7个时
				//第一行：(0)-110,-110	(1)--->0,-110	(2)--->110,-110
				//第二行：(3)-110,0		(4)--->0,0		(5)--->110,0
				//第三行：(6)-110,110

				//总数8个时
				//第一行：(0)-110,-110	(1)--->0,-110	(2)--->110,-110
				//第二行：(3)-110,0		(4)--->0,0		(5)--->110,0
				//第三行：(6)-110,110		(7)--->0,110

				//完美坐标站位样例：
				//总数7个时，多余的1个排在第三行的中间位置
				//第一行：(0)-110,-110	(1)--->0,-110	(2)--->110,-110
				//第二行：(3)-110,0		(4)--->0,0		(5)--->110,0
				//第三行：				(6)--->0,110

				//总数8个时，多余的2个排行第三行的中间位置
				//第一行：(0)-110,-110	(1)--->0,-110	(2)--->110,-110
				//第二行：(3)-110,0		(4)--->0,0		(5)--->110,0
				//第三行：		(6)-55,110		(7)--->55,110
			}
		}
	}

	return x, y
}

package liuyang

//获取坐标偏移地图中两个格子之间的距离。返回：以自身为起始的相距格子数量。
//如：0：与自身相距0格、1：与自身相距1格，也就是相邻的6个格子、2：与自身相距2格，也就是中间会隔一格(6个邻居)。
//参数：自身列、自身行、目标列、目标行、偏移坐标的类型：0尖朝上奇数行偏移、1尖朝上偶数行偏移、2边朝上奇数列偏移、3边朝上偶数列偏移、其他值视为0。
func GetCoordOffsetGridDistance(sX, sY, tX, tY, offsetType int) int {
	//如果自身坐标与目标坐标一样时，说明是同一坐标，直接返回0，代表0距离
	if sX == tX && sY == tY {
		//fmt.Printf("自身位置--->x：%v，y：%v--->目标位置--->x：%v，y：%v，两个位置为同一坐标，直接返回0距离\n", sX, sY, tX, tY)
		return 0
	}

	//1，先将两个偏移坐标转成立体坐标再进行运算

	//在转换坐标前，要区分当前供偏移坐标的类型，如果不在0到3之间时，则视为默认的0。
	var x1, y1, z1 int
	var x2, y2, z2 int

	//偏移坐标的类型：0奇数行偏移、1偶数行偏移、2奇数列偏移、3偶数列偏移
	switch offsetType {
	case 0:
		//偏移坐标转立方体坐标-尖朝上-奇数行偏移
		x1, y1, z1 = CoordOffsetToCubeTipOdd(sX, sY) //原点
		//偏移坐标转立方体坐标-尖朝上-奇数行偏移
		x2, y2, z2 = CoordOffsetToCubeTipOdd(tX, tY) //终点
	case 1:
		//偏移坐标转立方体坐标-尖朝上-偶数行偏移
		x1, y1, z1 = CoordOffsetToCubeTipEven(sX, sY) //原点
		//偏移坐标转立方体坐标-尖朝上-偶数行偏移
		x2, y2, z2 = CoordOffsetToCubeTipEven(tX, tY) //终点
	case 2:
		//偏移坐标转立方体坐标-边朝上-奇数列偏移
		x1, y1, z1 = CoordOffsetToCubeSideOdd(sX, sY) //原点
		//偏移坐标转立方体坐标-边朝上-奇数列偏移
		x2, y2, z2 = CoordOffsetToCubeSideOdd(tX, tY) //终点
	case 3:
		//偏移坐标转立方体坐标-边朝上-偶数列偏移
		x1, y1, z1 = CoordOffsetToCubeSideEven(sX, sY) //原点
		//偏移坐标转立方体坐标-边朝上-偶数列偏移
		x2, y2, z2 = CoordOffsetToCubeSideEven(tX, tY) //终点

	default:
		//偏移坐标转立方体坐标-尖朝上-奇数行偏移
		x1, y1, z1 = CoordOffsetToCubeTipOdd(sX, sY) //原点
		//偏移坐标转立方体坐标-尖朝上-奇数行偏移
		x2, y2, z2 = CoordOffsetToCubeTipOdd(tX, tY) //终点
	}

	//2，得到两个立体坐标的差值
	x := x1 - x2 //x轴相距格子数量
	y := y1 - y2 //y轴相距格子数量
	z := z1 - z2 //z轴相距格子数量

	//fmt.Printf("原点--->x：%v，y：%v--->转成的立体坐标--->x：%v，y：%v，z：%v\n", sX, sY, x1, y1, z1)
	//fmt.Printf("终点--->x：%v，y：%v--->转成的立体坐标--->x：%v，y：%v，z：%v\n", tX, tY, x2, y2, z2)

	//获取距离说明：
	//方式1：三个坐标中都取绝对值后，最大的值，就是距离。
	//方式2：三个坐标都取绝对值并相加后，除以2的值，就是距离。
	//这里，我们选用方式1获取距离

	//将负数转正
	if x < 0 {
		x *= -1
	}

	if y < 0 {
		y *= -1
	}

	if z < 0 {
		z *= -1
	}

	//获取三个数之间的最大值
	maxValue := GetThreeNubmerMax(x, y, z)

	//否则，返回临时变量上的最大值
	return maxValue
}

//==================================================================

//检测偏移坐标中(六边形地图)两个坐标点的距离是否在某一距离内(小于等于指定距离)
//参数：自身列、自身行、目标列、目标行、距离、偏移坐标的类型：0尖朝上奇数行偏移、1尖朝上偶数行偏移、2边朝上奇数列偏移、3边朝上偶数列偏移、其他值视为0。
func CheckOffsetCoordDistance(sX, sY, tX, tY, distance, offsetType int) bool {
	//获取坐标偏移地图中两个格子之间的距离。返回：以自身为起始的相距格子数量。
	tempDistance := GetCoordOffsetGridDistance(sX, sY, tX, tY, offsetType)
	//如果两个格子之间的距离小于等于参数中的距离时，说明符合条件
	if tempDistance <= distance {
		return true
	}

	return false
}

//检测偏移坐标中(六边形地图)两个坐标点的距离是否在某一距离区间内(大于等于最小距离并且小于等于最大距离)
//参数：自身列、自身行、目标列、目标行、距离最小值、距离最大值、偏移坐标的类型：0尖朝上奇数行偏移、1尖朝上偶数行偏移、2边朝上奇数列偏移、3边朝上偶数列偏移、其他值视为0。
func CheckOffsetCoordDistanceRange(sX, sY, tX, tY, distanceMin, distanceMax, offsetType int) bool {
	//获取坐标偏移地图中两个格子之间的距离。返回：以自身为起始的相距格子数量。
	tempDistance := GetCoordOffsetGridDistance(sX, sY, tX, tY, offsetType)
	//如果两个格子之间的距离小于等于参数中的距离时，说明符合条件
	if tempDistance >= distanceMin && tempDistance <= distanceMax {
		return true
	}

	return false
}

//检测偏移坐标中(六边形地图)两个坐标点的距离是否等于某一距离(等于指定距离)
//参数：自身列、自身行、目标列、目标行、距离、偏移坐标的类型：0尖朝上奇数行偏移、1尖朝上偶数行偏移、2边朝上奇数列偏移、3边朝上偶数列偏移、其他值视为0。
func CheckOffsetCoordDistanceEqual(sX, sY, tX, tY, distance, offsetType int) bool {
	//获取坐标偏移地图中两个格子之间的距离。返回：以自身为起始的相距格子数量。
	tempDistance := GetCoordOffsetGridDistance(sX, sY, tX, tY, offsetType)
	//如果两个格子之间的距离小于等于参数中的距离时，说明符合条件
	if tempDistance == distance {
		return true
	}

	return false
}

//==================================================================

//立方体偏移幅度(邻居)-数据集
var CubeOffsetArr = [][]int{
	{-1, +1, 0}, //左起第1个格子(x-1, y+1, z)
	{0, +1, -1}, //左起第2个格子(x, y+1, z-1)
	{+1, 0, -1}, //左起第3个格子(x+1, y, z-1)
	{+1, -1, 0}, //左起第4个格子(x+1, y-1, z)
	{0, -1, +1}, //左起第5个格子(x, y-1, z+1)
	{-1, 0, +1}, //左起第6个格子(x-1, y, z+1)
}

//获取立方体坐标的周围的邻居位置(六个格子)
func GetCubeNeighborArr(x, y, z int) [][]int {
	//格子位置数据集，位置元素最多6个。
	tempArr := make([][]int, 0)

	//遍历偶数行偏移幅度所需的数值集合
	for i := 0; i < len(CubeOffsetArr); i++ {
		//这里必须单独使用变量进行接收，不能使用x、y进行+=使用，不然值不固定会导致位置错误
		tempX := x + CubeOffsetArr[i][0]
		tempY := y + CubeOffsetArr[i][1]
		tempZ := z + CubeOffsetArr[i][2]

		//单个格子位置
		tempPosition := []int{tempX, tempY, tempZ}

		//将单个格子位置追加到位置数据集中
		tempArr = append(tempArr, tempPosition)
	}

	return tempArr
}

//偏移幅度(邻居)-尖朝上-奇数行偏移--->偶数行邻居-数据集。如：0、2、4、6、8行
var TipOddOffsetEvenRowArr = [][]int{
	{-1, 0},  //左起第1个格子(x-1, y)
	{-1, -1}, //左起第2个格子(x-1, y-1)
	{0, -1},  //左起第3个格子(x, y-1)
	{1, 0},   //左起第4个格子(x+1, y)
	{0, 1},   //左起第5个格子(x, y+1)
	{-1, 1},  //左起第6个格子(x-1, y+1)
}

//偏移幅度(邻居)-尖朝上-奇数行偏移--->奇数行邻居-数据集。如：1、3、5、7、9行
var TipOddOffsetOddRowArr = [][]int{
	{-1, 0}, //左起第1个格子(x-1, y)
	{0, -1}, //左起第2个格子(x, y-1)
	{1, -1}, //左起第3个格子(x+1, y-1)
	{1, 0},  //左起第4个格子(x+1, y)
	{1, 1},  //左起第5个格子(x+1, y+1)
	{0, 1},  //左起第6个格子(x, y+1)
}

//获取偏移坐标(尖朝上-奇数行偏移)周围的邻居位置(六个格子)
//参数：列位置、行位置、列数最大值、行数最大值
//返回：某位置周围合法的位置(除去<0或超出最大值或等于列最大值时的奇数行最后一个位置)
func GetOffsetTipOddNeighborArr(x, y, X_max, Y_max int) [][]int {
	//过滤非法坐标
	if x < 0 || x > X_max || y < 0 || y > Y_max {
		return nil
	}

	//如x最大值为8，y最大值为4，则参数8,3也是非法的，因为没有8,1、8,3、8,5这样的位置，所以也要过滤
	//过滤掉8,1、8,3、8,5。。。
	if x == X_max && y%2 == 1 {
		return nil
	}

	//格子位置数据集，位置元素最多6个。
	tempArr := make([][]int, 0)
	//如果y为偶数行时
	if y%2 == 0 {
		//遍历偶数行偏移幅度所需的数值集合
		for i := 0; i < len(TipOddOffsetEvenRowArr); i++ {
			//这里必须单独使用变量进行接收，不能使用x、y进行+=使用，不然值不固定会导致位置错误
			tempX := x + TipOddOffsetEvenRowArr[i][0]
			tempY := y + TipOddOffsetEvenRowArr[i][1]

			//过滤超出屏幕外的非法位置
			if tempX < 0 || tempY < 0 || tempX > X_max || tempY > Y_max {
				continue
			}

			//如x最大值为8，y最大值为4，则8,1、8,3、8,5位置也是非法的，因为没有这样的位置，所以过滤掉
			//过滤掉8,1、8,3、8,5。。。
			if tempX == X_max && tempY%2 == 1 {
				continue
			}

			//单个格子位置
			tempPosition := []int{tempX, tempY}

			//将单个格子位置追加到位置数据集中
			tempArr = append(tempArr, tempPosition)
		}
	} else {
		//否则遍历奇数行偏移幅度所需的数值集合
		for i := 0; i < len(TipOddOffsetOddRowArr); i++ {
			//这里必须单独使用变量进行接收，不能使用x、y进行+=使用，不然值不固定会导致位置错误
			tempX := x + TipOddOffsetOddRowArr[i][0]
			tempY := y + TipOddOffsetOddRowArr[i][1]

			//过滤超出屏幕外的非法位置
			if tempX < 0 || tempY < 0 || tempX > X_max || tempY > Y_max {
				continue
			}

			//如x最大值为8，y最大值为4，则8,1、8,3、8,5位置也是非法的，因为没有这样的位置，所以过滤掉
			//过滤掉8,1、8,3、8,5。。。
			if tempX == X_max && tempY%2 == 1 {
				continue
			}

			//单个格子位置
			tempPosition := []int{tempX, tempY}

			//将单个格子位置追加到位置数据集中
			tempArr = append(tempArr, tempPosition)
		}
	}

	return tempArr
}

//==========================================================================

//立方体坐标转偏移坐标-尖朝上-奇数行偏移
func CoordCubeToOffsetTipOdd(x, z int) (int, int) {
	//立体坐标转偏移坐标
	col := x + (z-(z&1))/2 //列
	row := z               //行
	//fmt.Printf("col：%d，row：%d\n", col, row)

	return col, row
}

//偏移坐标转立方体坐标-尖朝上-奇数行偏移
func CoordOffsetToCubeTipOdd(col, row int) (int, int, int) {
	//偏移坐标转立体坐标
	x := col - (row-(row&1))/2
	z := row
	y := -x - z
	//fmt.Printf("x：%d，y：%d，z：%d\n", tempX, tempY, tempZ)

	return x, y, z
}

//立方体坐标转偏移坐标-尖朝上-偶数行偏移
func CoordCubeToOffsetTipEven(x, z int) (int, int) {
	//立体坐标转偏移坐标
	col := x + (z+(z&1))/2 //列
	row := z               //行
	//fmt.Printf("col：%d，row：%d\n", col, row)

	return col, row
}

//偏移坐标转立方体坐标-尖朝上-偶数行偏移
func CoordOffsetToCubeTipEven(col, row int) (int, int, int) {
	//偏移坐标转立体坐标
	x := col - (row+(row&1))/2
	z := row
	y := -x - z
	//fmt.Printf("x：%d，y：%d，z：%d\n", tempX, tempY, tempZ)

	return x, y, z
}

//立方体坐标转偏移坐标-边朝上-奇数列偏移
func CoordCubeToOffsetSideOdd(x, z int) (int, int) {
	//立体坐标转偏移坐标
	col := x               //列
	row := z + (x-(x&1))/2 //行
	//fmt.Printf("col：%d，row：%d\n", col, row)

	return col, row
}

//偏移坐标转立方体坐标-边朝上-奇数列偏移
func CoordOffsetToCubeSideOdd(col, row int) (int, int, int) {
	//偏移坐标转立体坐标
	x := col
	z := row - (col-(col&1))/2
	y := -x - z
	//fmt.Printf("x：%d，y：%d，z：%d\n", tempX, tempY, tempZ)

	return x, y, z
}

//立方体坐标转偏移坐标-边朝上-偶数列偏移
func CoordCubeToOffsetSideEven(x, z int) (int, int) {
	//立体坐标转偏移坐标
	col := x               //列
	row := z + (x+(x&1))/2 //行
	//fmt.Printf("col：%d，row：%d\n", col, row)

	return col, row
}

//偏移坐标转立方体坐标-边朝上-偶数列偏移
func CoordOffsetToCubeSideEven(col, row int) (int, int, int) {
	//偏移坐标转立体坐标
	x := col
	z := row - (col+(col&1))/2
	y := -x - z
	//fmt.Printf("x：%d，y：%d，z：%d\n", tempX, tempY, tempZ)

	return x, y, z
}

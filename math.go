package liuyang

import (
	"math"
)

//以角度值获取弧度值，参数：角度
func GetRadian(angle float64) float64 {
	//弧度值 = 角度 * (pi/180)
	radian := angle * (math.Pi / 180) //角度转弧度

	return radian
}

//以弧度值获取角度值，参数：弧度
func GetAngle(radian float64) float64 {
	//角度值 = 弧度 * (180/pi)
	angle := radian * (180 / math.Pi) //弧度转角度

	return angle
}

//以弧度值获取x与y的向量值，参数：弧度
func GetVector(radian float64) (float64, float64) {
	//以弧度值获取余弦值和正弦值
	stepX := math.Cos(radian) //得到x方向的每步速度
	stepY := math.Sin(radian) //得到y方向的每步速度

	return stepX, stepY
}

//获取两点之间的直线距离，利用勾股定理，求两个坐标点直线距离。参数：a坐标点的x,y和b坐标点的x,y。返回两点的直线距离。
func GetTwoPointDistance(aX, aY, bX, bY float64) float64 {
	//这里不要取绝对值，因为3 - -2 的实际长度为5，但取绝对值后会是3-2=1，所以不要取绝值对。
	a := aX - bX //得到直角三角形a边的边长
	b := aY - bY //得到直角三角形b边的边长

	a2 := math.Pow(a, 2) //得到a的平方
	b2 := math.Pow(b, 2) //得到b的平方

	c2 := a2 + b2 //a方 + b方 = c方
	//fmt.Println("a的2次方是：", a2, "b的2次方是：", b2, "c的2次方是：", c2)
	c := math.Sqrt(c2) //求c的平方根
	//fmt.Println("c的2次方被求根后的值是：", c)

	return c
}

//获取某位置坐标距离0,0坐标的角度值。返回(0-359之间的float值)
//参数说明：t为target简写字母(目标)，s为self简写字母(自己)
func GetPositionAngel(tX, tY, sX, sY float64) float64 {
	//1，以目标(敌方)坐标 - 自己坐标 = 某点坐标距离0,0坐标的差值坐标
	offsetX := tX - sX
	offsetY := tY - sY
	//2，再以差值坐标获取距离0,0坐标的角度。注意：先y轴再x轴，得到弧度值
	radian := math.Atan2(offsetY, offsetX)
	//3，将弧度值转角度值
	angle := GetAngle(radian) //会得到从左上到左下一圈的-180到180之间的角度
	//fmt.Println("获取某位置坐标(x：", offsetX, "，y：", offsetY, ")距离0,0坐标的角度值：", angle)
	//由于Atan2函数会根据x，y的正负号确定象限，所以得到的弧度值转角度后，会得到-180到180之间的角度，所以要判断一下，为负数时+360才对。
	//这样就会得到：0-359之间的float值
	if angle < 0 {
		angle = angle + 360
	}

	return angle
}

//以圆心点坐标、角度、半径获取圆边某点坐标。参数角度为正数或按象线位置传负数都可以。
func GetRoundEdgePosition(x0, y0, angle, r float64) (float64, float64) {
	radian := GetRadian(angle) //角度值转为弧度值
	//fmt.Println("弧度：", radian, "，转角度：", tempAngle)
	//打印：弧度： 3.6651914291880923 ，转角度： 210.00000000000003
	x1 := x0 + r*math.Cos(radian) //圆心点x坐标 + 半径 * 余弦值(也就是x轴的向量)
	y1 := y0 + r*math.Sin(radian) //圆心点y坐标 + 半径 * 正弦值(也就是y轴的向量)
	//fmt.Println("得到的圆边某点坐标------>x：", x1, "，y：", y1)

	return x1, y1
}

//以出招朝向角度获取左右两侧面向的角度(0-359之间的正整数)
func GetDirectionSideTwoAngle(dirAngle int) (int, int, int) {
	//0(右)、90(下)、180(左)、270(上)
	d_angle := GetPositiveAngle(dirAngle)
	l_angle := d_angle - 90  //朝向角度时的左侧面向的角度
	r_angle := dirAngle + 90 //朝向角度时的右侧面向的角度

	l_angle = GetPositiveAngle(l_angle)
	r_angle = GetPositiveAngle(r_angle)

	return d_angle, l_angle, r_angle
}

//获取符合判断条件的角度值。参数：朝向角度、敌人距0,0的角度(-90至450之间的整数)
func GetJudgeAngle(dirAngle, enemyAngle int) (int, int) {
	//0(右)、90(下)、180(左)、270(上)
	//直接模掉360度
	//d_angle := dirAngle % 360
	//e_angle := enemyAngle % 360

	//获取正数角度值。返回：0-359之间的正整数
	d_angle := GetPositiveAngle(dirAngle)
	e_angle := GetPositiveAngle(enemyAngle)

	if d_angle >= 0 && d_angle < 90 { //如果朝向右下时
		//问题：
		//得到敌方坐标 9 4 和 5 5 的差值坐标： 4 -1 与0,0坐标的角度为： 345.96375653207355
		//出招朝向角度： 1 敌方所在角度： 345.96375653207355 ，区间差值角度： -29 31
		if e_angle >= 270 && e_angle < 360 {
			e_angle = e_angle - 360 //敌方所在角度由于在0度之上，需要减掉360度变成-5.5度才能符合判断需要
		}
		//解决：
		//得到敌方坐标 9 4 和 5 5 的差值坐标： 4 -1 与0,0坐标的角度为： 345.96375653207355
		//出招朝向角度： 1 敌方所在角度： -14.036243467926454 ，区间差值角度： -29 31
		//条件1达成------>出招朝向角度： 1 敌方所在角度： -14.036243467926454 符合区间差值角度： -29 31
	} else if d_angle >= 270 && d_angle < 360 { //如果朝向右上时
		//问题：
		//得到敌方坐标 9 6 和 5 5 的差值坐标： 4 1 与0,0坐标的角度为： 14.036243467926479
		//出招朝向角度： 359 敌方所在角度： 14.036243467926479 ，区间差值角度： 329 389
		if e_angle >= 0 && e_angle < 90 {
			e_angle = e_angle + 360 //敌方所在角度由于模掉了360度，需要加回来变成374度才能符合判断需要
		}
		//得到敌方坐标 9 6 和 5 5 的差值坐标： 4 1 与0,0坐标的角度为： 14.036243467926479
		//出招朝向角度： 359 敌方所在角度： 374.03624346792645 ，区间差值角度： 329 389
		//条件1达成------>出招朝向角度： 359 敌方所在角度： 374.03624346792645 符合区间差值角度： 329 389
	}

	return d_angle, e_angle
}

//以出招朝向角度获取敌方最近距离面向的角度。参参：出招朝向角度、敌方距离0,0的角度。
//返回：出招朝向角度、敌方距离0,0的角度、敌方距离出招角度最近的角度(0-359之间的正整数)
func GetDirectionSideEnemyAngle(dirAngle, enemyAngle int) (int, int, int) {
	//0(右)、90(下)、180(左)、270(上)
	//定义临时变量
	var near_angle = 0 //敌方距离出招角度最近的角度

	//获取符合判断条件的角度值
	d_angle, e_angle := GetJudgeAngle(dirAngle, enemyAngle)

	//fmt.Println("获取符合判断条件的角度值：", d_angle, e_angle)

	//如果出招朝向角度大于敌方角度时，也就是敌人在出招朝向的左侧时，则+90度
	if d_angle > e_angle {
		near_angle = d_angle + 90
	} else if d_angle < e_angle { //如果出招朝向角度小于敌方角度时，也就是敌人在出招朝向的右侧时，则-90度
		near_angle = d_angle - 90
	} else if d_angle == e_angle { //如果出招朝向角度等于敌方角度时，也就是敌人与出招朝向为同一条线上时，则取对向角度，也就是+180度
		near_angle = d_angle + 180
	}

	//将朝向角度置为0-359之间
	d_angle = GetPositiveAngle(d_angle)

	//将敌方角度置为0-359之间
	e_angle = GetPositiveAngle(e_angle)

	//将敌方距离出招朝向的最近角度置为0-359之间
	near_angle = GetPositiveAngle(near_angle)

	return d_angle, e_angle, near_angle
}

//获取正整数角度值。返回：0-359之间的正整数
func GetPositiveAngle(angle int) int {
	//将角度置为0-359之间
	if angle <= -360 || angle >= 360 { //必须要判断大于等于或小于等于才行
		angle = angle % 360
	}

	//这里必须要单独再判断一下，因为负数求模后还是负数，此时就+360变为正数
	if angle < 0 {
		angle = angle + 360
	}

	return angle
}

//样例
//判定值： 60 ------>随机数： 99
//false
//判定值： 60 ------>随机数： 25
//true
//判定值： 60 ------>随机数： 88
//false
//判定值： 60 ------>随机数： 92
//false
//判定值： 60 ------>随机数： 9
//true
//获取机率触发的可能性。参数：机率值0-100，返回：机率达成返true、未达成返false。
func GetOdds(n int) bool {
	//如果机率值小于等于0时，则百分之百返回false。
	if n <= 0 {
		return false
	}

	//如果机率值大于等于100时，则百分之百返回true。
	if n >= 100 {
		return true
	}

	//随机0-99之间的数，不会超过100。
	randX := RandomNumber(100)
	//fmt.Println("判定值：", n, "------>随机数：", randX)
	if randX < n {
		return true
	}

	return false
}

//获取物体对象碰到另一物体对象后的反弹向量(参考传奇占位，A英雄碰到B英雄后的反弹向量，也就是某个方向每步要移动的值)。返回：反弹角度、x方向每步移动值、y方向每步移动值。不进行移动。
//参数说明：s为self简写字母(自己)、t为target简写字母(目标)
//判断说明：以左上0,0为起始点
func GetObjectReboundVector(sX, sY, tX, tY float64) (float64, float64) {
	//获取物体对象碰到另一物体对象后的反弹角度与弧度
	_, radian := GetObjectReboundAngle(sX, sY, tX, tY)
	//得到x与y方向的每步速度
	stepX, stepY := GetVector(radian)

	return stepX, stepY
}

//获取物体对象碰到边界后的反弹向量(也就是某个方向每步要移动的值)。返回：反弹角度、x方向每步移动值、y方向每步移动值。不进行移动。
//参数说明：边界标识->上1、左2、下3、右4
//判断说明：以左上0,0为起始点
func GetBorderReboundVector(border_flag uint32) (float64, float64) {
	//过滤，必须要单独过滤，不然即使下面的函数会返角度0和弧度0，也会导至cos(0)=1,sin(0)=0，也就是stepX会是1，也就是整个函数会返1,0。而不是想要的0,0
	if border_flag < border_Up || border_flag > border_Right {
		return 0, 0
	}

	//获取物体碰到边框后反弹时的随机角度与弧度
	_, radian := GetBorderReboundAngle(border_flag)

	stepX, stepY := GetVector(radian) //得到x与y方向的每步速度

	return stepX, stepY
}

//获取物体对象碰到另一物体对象后的反弹角度与弧度(参考传奇占位，A英雄碰到B英雄后的反弹)。返回：反弹角度、弧度
//参数说明：s为self简写字母(自己)、t为target简写字母(目标)
//判断说明：坐标以左上0,0为起始点、角度以右起顺时针0到360
func GetObjectReboundAngle(sX, sY, tX, tY float64) (int, float64) {
	//临时变量
	var randX = 0 //接收随机数

	//如果当前球在右下方时，弹向右下角
	if sX > tX && sY > tY {
		//如果当前球在右下方时，弹向右下角
		randX = RandomNumberRange(0, 90)

	} else if sX > tX && sY < tY {
		//如果当前球在右上方时，弹向右上角
		randX = RandomNumberRange(270, 360)

	} else if sX < tX && sY < tY {
		//如果当前球在左上方时，弹向左上角
		randX = RandomNumberRange(180, 270)

	} else if sX < tX && sY > tY {
		//如果当前球在左下方时，弹向左下角
		randX = RandomNumberRange(90, 180)

	} else if sX < tX && sY == tY {
		//如果当前球与另一球在同一水平线时，且当前球在正左侧时，弹向左侧面
		randX = RandomNumberRange(90, 270)

	} else if sX > tX && sY == tY {
		//如果当前球与另一球在同一水平线时，且当前球在正右侧时，弹向右侧面
		tempX := RandomNumberRange(270, 450)
		randX = tempX % 360 //超出360时必须取模值

	} else if sX == tX && sY < tY {
		//如果当前球与另一球在同一竖平线时，且当前球在正上侧时，弹向上侧面
		randX = RandomNumberRange(180, 360)

	} else if sX == tX && sY > tY {
		//如果当前球与另一球在同一竖平线时，且当前球在正下侧时，弹向下侧面
		randX = RandomNumberRange(0, 180)

	} else if sX == tX && sY == tY {
		//如果当前球与另一球在同一坐标点上时，则弹向0-360
		randX = RandomNumberRange(0, 360)
	}

	//以角度值获取弧度值
	radian := GetRadian(float64(randX))

	return randX, radian
}

//获取物体碰到边框后的反弹角度与弧度。参数：碰到的边界标识。返回：角度值、弧度值
//判断说明：坐标以左上0,0为起始点、角度以右起顺时针0到360
func GetBorderReboundAngle(border_flag uint32) (int, float64) {
	//过滤，必须
	if border_flag < border_Up || border_flag > border_Right {
		return 0, 0
	}

	//临时变量
	var randX = 0 //接收随机数

	if border_flag == border_Up { //如果碰到上边界
		randX = RandomNumberRange(10, 170)
		//fmt.Println("当前角度：", angle, "，碰到上边界，返回10-170的新角度为：", newAngle)
	} else if border_flag == border_Down {
		randX = RandomNumberRange(190, 350)
		//fmt.Println("当前角度：", angle, "，碰到下边界，返回190-350的新角度为：", newAngle)
	} else if border_flag == border_Left {
		tempX := RandomNumberRange(280, 440)
		randX = tempX % 360
		//fmt.Println("当前角度：", angle, "，碰到左边界，返回280-80的新角度为：", newAngle)
	} else if border_flag == border_Right {
		randX = RandomNumberRange(100, 260)
		//fmt.Println("当前角度：", angle, "，碰到右边界，返回100-260的新角度为：", newAngle)
	}

	//以角度值获取弧度值
	radian := GetRadian(float64(randX))

	return randX, radian
}

//获取圆的两侧切线坐标--->点p到圆的右侧切线坐标、点p到圆的左侧切线坐标、点p与圆的切线->与点p与圆心直线的夹角度数、是否获取到切线坐标与角度值
//说明：点p如果在圆内时是得不到切线坐标的，所以会返：0,0,0,0,0,false
//说明：点p如果在圆外时是可以得到切线坐标的，所以会返：坐标1、坐标2、角度值、true
//	:param px: 圆外点P横坐标
//	:param py: 圆外点P纵坐标
//	:param cx: 圆心横坐标
//	:param cy: 圆心纵坐标
//	:param r:  圆半径
//	:return:   过点P的俩条切线坐标
//作者：josepa https: //www.bilibili.com/read/cv13951079 出处：bilibili
func GetCircleTangentCoor(px, py, cx, cy, r float64) (float64, float64, float64, float64, float64, bool) {
	//求点p到圆心的距离
	distance := math.Sqrt((px-cx)*(px-cx) + (py-cy)*(py-cy))
	//fmt.Println("两点直线距离：", distance)

	//aa := GetTwoPointDistance(px, py, cx, cy)
	//fmt.Println("刘阳函数获取的两点直线距离：", aa)

	//点p所在圆内时，这里必须要校验，不然下面求点p到切点的距离会是NaN
	if distance <= r {
		//fmt.Println("输入的数值不在范围内--->点p所在圆内")
		return 0, 0, 0, 0, 0, false
	}

	//否则点p所在圆外时，继续下面的逻辑

	//点p到切点的距离
	length := math.Sqrt(distance*distance - r*r)
	//fmt.Println("点p到切点的距离：", length) //点p到切点的距离： NaN

	//点p到圆心的单位向量
	ux := (cx - px) / distance
	uy := (cy - py) / distance
	//fmt.Println("ux：", ux)
	//fmt.Println("uy：", uy)
	//ux： 0.7071067811865475
	//uy： 0.7071067811865475

	//计算切线与圆心连线的夹角(弧度值)
	angle := math.Asin(r / distance) //返回的是弧度值
	//fmt.Println("计算切线与圆心连线的夹角(弧度值)：", angle)

	//将弧度值转成角度值
	a_angle := GetAngle(angle)
	//fmt.Println("将弧度值转成角度值：", a_angle)
	//计算切线与圆心连线的夹角(弧度值)： 1.084101371720123
	//将弧度值转成角度值： 62.11443316390628

	//向正反两个方向旋转单位向量
	q1x := ux*math.Cos(angle) - uy*math.Sin(angle)
	q1y := ux*math.Sin(angle) + uy*math.Cos(angle)
	q2x := ux*math.Cos(-angle) - uy*math.Sin(-angle)
	q2y := ux*math.Sin(-angle) + uy*math.Cos(-angle)

	//得到新坐标
	q1x = q1x*length + px
	q1y = q1y*length + py
	q2x = q2x*length + px
	q2y = q2y*length + py

	return q1x, q1y, q2x, q2y, a_angle, true
}

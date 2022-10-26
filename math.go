package liuyang

import "math"

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
//参数说明：e为enemy简写字母(敌人)，s为self简写字母(自己)
func GetPositionAngel(eX, eY, sX, sY float64) float64 {
	//1，以敌方坐标 - 自己坐标 = 某点坐标距离0,0坐标的差值坐标
	offsetX := eX - sX
	offsetY := eY - sY
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

//以圆心点坐标、角度、半径获取圆边某点坐标位置。参数角度为正数或按象线位置传负数都可以。
func GetRoundEdgePosition(x0, y0, angle, r float64) (float64, float64) {
	radian := GetRadian(angle) //角度值转为弧度值
	//tempAngle := liuyang.GetAngle(radian)
	//fmt.Println("弧度：", radian, "，转角度：", tempAngle)
	//打印：弧度： 3.6651914291880923 ，转角度： 210.00000000000003
	x1 := x0 + r*math.Cos(radian) //圆心点x坐标 + 半径 * 余弦(角度转换后的弧度值)
	y1 := y0 + r*math.Sin(radian) //圆心点y坐标 + 半径 * 正弦(角度转换后的弧度值)
	//fmt.Println("得到的圆边某点坐标------>x：", x1, "，y：", y1)

	return x1, y1
}

//以出招朝向角度获取左右两侧面向的角度(0-359之间的正整数)
func GetDirectionSideTwoAngle(dirAngle int32) (int32, int32, int32) {
	//0(右)、90(下)、180(左)、270(上)
	d_angle := GetPositiveAngle(dirAngle)
	l_angle := d_angle - 90  //朝向角度时的左侧面向的角度
	r_angle := dirAngle + 90 //朝向角度时的右侧面向的角度

	l_angle = GetPositiveAngle(l_angle)
	r_angle = GetPositiveAngle(r_angle)

	return d_angle, l_angle, r_angle
}

//获取符合判断条件的角度值。参数：朝向角度、敌人距0,0的角度(-90至450之间的整数)
func GetJudgeAngle(dirAngle, enemyAngle int32) (int32, int32) {
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
func GetDirectionSideEnemyAngle(dirAngle, enemyAngle int32) (int32, int32, int32) {
	//0(右)、90(下)、180(左)、270(上)
	//定义临时变量
	var near_angle int32 = 0 //敌方距离出招角度最近的角度

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
func GetPositiveAngle(angle int32) int32 {
	//将角度置为0-359之间
	if angle >= 360 || angle <= 360 { //必须要判断大于等于或小于等于才行
		angle = angle % 360
	}

	//这里必须要单独再判断一下，因为负数求模后还是负数，此时就+360变为正数
	if angle < 0 {
		angle = angle + 360
	}

	return angle
}

package liuyang

//技能出招相关函数

//边框标识定义
const bound_Up = 1    //上
const bound_Left = 2  //左
const bound_Down = 3  //下
const bound_Right = 4 //右

//=======================================================弹球碰到边框后的随机反弹角度值与弧度值=======================================================
//获取物体碰到边框后反弹时的随机角度与弧度。参数：碰到的边界标识。返回：角度值、弧度值
func GetRandomAngleRadian(bound uint32) (float64, float64) {
	var newAngle float64 = 0

	if bound == bound_Up { //如果碰到上边界
		randX := RandomNumberRange(10, 170)
		newAngle = float64(randX)
		//fmt.Println("当前角度：", angle, "，碰到上边界，返回10-170的新角度为：", newAngle)
	} else if bound == bound_Down {
		randX := RandomNumberRange(190, 350)
		newAngle = float64(randX)
		//fmt.Println("当前角度：", angle, "，碰到下边界，返回190-350的新角度为：", newAngle)
	} else if bound == bound_Left {
		randX := RandomNumberRange(280, 440)
		tempN := randX % 360
		newAngle = float64(tempN)
		//fmt.Println("当前角度：", angle, "，碰到左边界，返回280-80的新角度为：", newAngle)
	} else if bound == bound_Right {
		randX := RandomNumberRange(100, 260)
		newAngle = float64(randX)
		//fmt.Println("当前角度：", angle, "，碰到右边界，返回100-260的新角度为：", newAngle)
	}

	//角度 * (pi / 180) = 弧度
	radian := GetRadian(newAngle)

	return newAngle, radian
}

//样例：
//	圆形物体对象定义时，一定要记得x与y的坐标是从左上0,0起点开始，加上半径后才是圆心点
//	radian := liuyang.GetRadian(angle)//以角度获取弧度
//	prey := &Prey{
//		IndexID:    self.PreyIndexID, //猎物索引ID，不重复，从1开始累加，使用后记得++进行自增1
//		ID:         id,               //猎物表中的id
//		Radius:     radius,           //猎物(气泡)半径
//		X:          x + radius,       //左上起始点 + 半径 = x轴圆心点，必须。因为两圆碰撞时计算的就是两圆心点的直线距离
//		Y:          y + radius,       //左上起始点 + 半径 = y轴圆心点，必须。因为两圆碰撞时计算的就是两圆心点的直线距离
//		Speed:      speed,            //速度
//		Angle:      angle,            //角度
//		Radian:     radian,           //弧度
//		StepX:      math.Cos(radian), //x方向每步速度(余弦值)
//		StepY:      math.Sin(radian), //y方向每步速度(正弦值)
//		FixedBlood: blood,            //固定血量
//		Blood:      blood,            //剩余血量
//		MScore:     score,            //奖励积分数(小积分)
//		IsYard:     false,            //是否已进入战场，默认false
//		IsHit:      false,            //是否被击中，击中后会变红一下。默认false
//		IsSweat:    false,            //是否流汗，血量低于30%出现特殊效果：流汗。默认false
//		IsKill:     false,            //是否被击杀，默认false
//	}

//判断两个圆是否发生碰撞
//参数说明：s为self简写字母(自己)、e为enemy简写字母(敌人)、x轴、y轴、r半径长度
func IsCircleCollision(sX, sY, sR, eX, eY, eR float64) bool {
	//1，首先得到两个圆的圆心点的坐标。然让计算两个的圆心点之间的直线距离。
	//2，如果两圆心点的直线距离小于等于两圆半径之和时，说明两圆已发生碰撞
	//3，否则两圆心点的直线距离大于两圆半径之和时，说明两圆未发生碰撞
	//获取默认两点的直线距离，如果两圆的大小相同，则不会有问题。如果两圆大小不同时，就会发生圆重叠或距离很远时发生碰撞的问题。
	//因为默认是以左上角为起始点，所以想要获取两圆的圆心点的话，就得以自身坐标+自身半径得到自身圆心点坐标，再去判断两个圆心点的直线距离是否小于等于两圆半径之和才行。
	//如果x、y在定义时加上了半径的话，这里就可以直接这样传值了。因为此时的x、y就是圆心点。
	distance := GetTwoPointDistance(sX, sY, eX, eY)

	//fmt.Println("直线距离：", distance)

	radiusSum := sR + eR //两圆半径之和
	//如果两个圆心点的直线距离小于等于两个圆的半径之和时，则认为发生碰撞
	if distance <= radiusSum {
		return true
	}

	return false
}

//=======================================================检测敌人位置是否在自己出招范围内=======================================================

//扇面技能出招
//1，以敌方坐标 - 自己坐标 = 某点坐标距离0,0坐标的差值坐标
//2，以差值坐标获取到距离0,0坐标的角度值
//3，判断该角度值是否 >= 自己出招角度-出招扇面宽度 && 该角度值是否 <= 自己出招角度+出招扇面宽度
//4，只有符合条件的，才能去判断自己坐标与敌方坐标的两点直线距离是否小于等于自己出招的长度
//5，之所以最后才判断两点直线距离是否小于自己出招(半径)距离，是因为测试的时候方便，直接就可以循环中自增半径长度的方式进行测试。

//检测是否命中敌方单位-扇形技能
//参数说明：e为enemy简写字母(敌人)、s为self简写字母(自己)、扇面夹角度数、扇面长度(以自己为圆心的半径)、扇面朝向角度、敌方底座半径(人物脚下蓝圈)
func IsHit_Fan(eX, eY, sX, sY, sk_a, sk_l, dirAngle float64) bool {
	//1，以敌方坐标 - 自己坐标 = 某点坐标距离0,0坐标的差值坐标，再以差值坐标获取到敌方位置距离0,0坐标的角度值
	e_angle := GetPositionAngel(eX, eY, sX, sY)
	//fmt.Println("得到敌方坐标", eX, eY, "和", sX, sY, "的差值坐标：", eX-sX, eY-sY, "与0,0坐标的角度为：", e_angle)
	i_e_angle := Float64ToInt32(e_angle)

	var offsetWeight = sk_a / 2 //扇面夹角宽度的一半

	i_dirAngle := Float64ToInt32(dirAngle)

	//获取符合判断条件的角度值
	i_d_angle, i_e_angle := GetJudgeAngle(i_dirAngle, i_e_angle)

	var dir_angle = float64(i_d_angle)
	e_angle = float64(i_e_angle)

	//fmt.Println("出招朝向角度：", i_d_angle, "敌方所在角度：", e_angle, "，区间差值角度：", dir_angle-offsetWeight, dir_angle+offsetWeight)

	//3，判断该角度值是否 >= 自己出招角度-出招扇面宽度 && 该角度值是否 <= 自己出招角度+出招扇面宽度
	if e_angle >= dir_angle-offsetWeight && e_angle <= dir_angle+offsetWeight {
		//4，只有符合条件的，才能去判断自己坐标与敌方坐标的两点直线距离是否小于等于自己出招的长度
		//fmt.Println("条件1达成------>出招朝向角度：", dir_angle, "敌方所在角度：", e_angle, "符合区间差值角度：", dir_angle-offsetWeight, dir_angle+offsetWeight)
		//获取两点之间的直线距离
		distance := GetTwoPointDistance(sX, sY, eX, eY)
		if distance <= sk_l {
			//fmt.Println("自己坐标与敌方坐标的两点直线距离：", distance, "，小于等于自己出招的长度：", sk_l)
			//===============================执行扇面覆盖敌人条件成立后的逻辑===============================
			//
			//
			//
			//===============================执行扇面覆盖敌人条件成立后的逻辑===============================

			return true
		}
	}

	//注意：
	//1，如果上面以敌方中心点为准的判断如果不成立时，还要检测一下敌方距离出招角度最近的坐标点，判断该坐标点是否符合条件才行。就是技能区域粘到敌方边缘位置的情况，
	//2，新的判断两点之间的直线距离就要以敌方距离出招角度最近的坐标为准了。
	//但是，这里千万不要使用递规方式再次调用当前函数，否则会导致无限递规的发生。
	//原因如下：
	//如果出招角度未击中敌方时，也就是顺序到达此处调用本函数方法，接着又会在本函数中继续执行到这里，这就会导致连return都没有机会执行，造成无限递规。
	//步骤：角度不在扇面中->调用本函数->角度不在扇面中->调用本函数->无限递规(因为本函数中的return false根本没机会执行)
	//所以，在使用递规时，一定要注意返回的条件。
	//另外需要注意的是：递规中的参数最好是可变的，不然参数不变的情况，就会容易忽略返回的条件，进而造成无限递规。

	return false
}

//长方形技能出招
//待完成：
//1，获取技能出招时朝向的角度的两侧面向的角度值。
//2，封装函数：以圆心点、半径、角度获取圆边某点坐标
//3，判断疾光电影所朝方向时以自己为圆心，以技能招式的宽度为半径得到技能招式的左右圆边上的两点。
//然后判断覆盖面上的所有人，得到技能招式圆边左点与当前敌人的角度和技能招式圆边右点与当前敌人的角度。
//然后判断，某人与左侧点角度值 大于等于 技能出招朝向角度 并且 某人与右侧点角度值 小于等于 技能出招朝向角度
//4，之所以最后才判断两点直线距离是否小于自己出招(半径)距离，是因为测试的时候方便，直接就可以循环中自增半径长度的方式进行测试。

//检测是否命中敌方单位-长方形技能
//参数说明：e为enemy简写字母(敌人)、s为self简写字母(自己)、长方技能宽度、长方技能长度、长方技能朝向角度
//出招宽度，宽度除以2才是以自己为圆心的以左右为两点的圆心半径，这样才能得到左右两点坐标
//条件1：敌方位置在出招技能宽度范围内(左右两点的角度内)
//条件2：敌方位置在出招技能长度范围内(半径长度内)
func IsHit_Rectangle(eX, eY, sX, sY, sk_w, sk_l, dirAngle float64) bool {
	i_dirAngle := Float64ToInt32(dirAngle)
	d_angle, l_angle, r_angle := GetDirectionSideTwoAngle(i_dirAngle) //以出招朝向角度获取左右两侧面向的角度

	var sk_r = sk_w / 2 //扇面夹角宽度的一半

	//GetRoundEdgePosition()函数中，角度按大于0的角度或按象线位置的负数角度传值都是一样的结果。
	//var aa = dirAngle - 90
	//var bb = dirAngle + 90
	//l_x1, l_y1 := GetRoundEdgePosition(sX, sY, aa, sk_r) //得到出招时的圆边左点位置
	//r_x1, r_y1 := GetRoundEdgePosition(sX, sY, bb, sk_r) //得到出招时的圆边右点位置
	//fmt.Println("未处理出招长方形时的朝向角度：", dirAngle, "，左右两侧角度：", aa, bb, "，两侧坐标点：", l_x1, l_y1, r_x1, r_y1)

	//以圆心点坐标、角度、半径获取圆边某点坐标位置
	l_x, l_y := GetRoundEdgePosition(sX, sY, float64(l_angle), sk_r) //得到出招时的圆边左点位置
	r_x, r_y := GetRoundEdgePosition(sX, sY, float64(r_angle), sk_r) //得到出招时的圆边右点位置

	//fmt.Println("处理后出招长方形时的朝向角度：", d_angle, "，左右两侧角度：", l_angle, r_angle, "，两侧坐标点：", l_x, l_y, r_x, r_y)

	//2，以敌方坐标和出招朝向时的两侧坐标点计算出差值坐标对应0,0坐标的角度
	e_l_angle := GetPositionAngel(eX, eY, l_x, l_y) //以敌方坐标和出招时的左侧坐标得到差值坐标对应0,0坐标的角度
	e_r_angle := GetPositionAngel(eX, eY, r_x, r_y) //以敌方坐标和出招时的右侧坐标得到差值坐标对应0,0坐标的角度

	//fmt.Println("以敌方坐标和出招朝向时的两侧坐标点计算出差值坐标对应0,0坐标的角度：", e_l_angle, e_r_angle)

	//===================================注意，为避免359、0、1这样的求模过渡
	//朝向角度大于等于0度并且小于90度时，要将右侧角度值+360度，将其变为大于左侧180度，这样才能保证大于左边的起始270度值
	//朝向角度： 0 ，获取左右角度为--->left： 270 --->right： 90
	//朝向角度： 1 ，获取左右角度为--->left： 271 --->right： 91
	//朝向角度： 2 ，获取左右角度为--->left： 272 --->right： 92

	//朝向角度： 89 ，获取左右角度为--->left： 359 --->right： 179
	//朝向角度： 90 ，获取左右角度为--->left： 0 --->right： 180
	//朝向角度： 91 ，获取左右角度为--->left： 1 --->right： 181

	//朝向角度大于等于270度并且小于360度时，要将左侧角度值-360度，将其变为小于右侧180度，这样才能保集小于右边的起始0度值
	//朝向角度： 269 ，获取左右角度为--->left： 179 --->right： 359
	//朝向角度： 270 ，获取左右角度为--->left： 180 --->right： 0	#这里开始应该判断一下，将右侧角度按左侧角度+180后才是正确的
	//朝向角度： 271 ，获取左右角度为--->left： 181 --->right： 1

	//朝向角度： 358 ，获取左右角度为--->left： 268 --->right： 88
	//朝向角度： 359 ，获取左右角度为--->left： 269 --->right： 89
	//朝向角度： 360 ，获取左右角度为--->left： 270 --->right： 90

	//由于左侧坐标点与敌方坐标点的角度一直要小于右侧坐标点与敌方坐标点的角度的，但是朝向右侧时，会有357、358、359、0、1、2这样的满360会求模的地方
	//就会导致左侧坐标点与敌方坐标点的角度大于右侧坐标点与敌方坐标点的角度，那么，此时就要判断一下：
	//如果出招朝向角度为右下时，则右侧坐标点与敌方坐标点的角度需要-360，以便小于0度的判断需要。
	//如果出招朝向角度为右上时，则左侧坐标点与敌方坐标点的角度需要+360，以便大于360度的判断需要。
	//前题条件必须是以下两个条件都成立时才可以操作。
	//1，右侧坐标点与敌方坐标点的角度为270到360(不含360)度之间
	//2，左侧坐标点与敌方坐标点的角度为0到90(不含90)度之间时。
	if d_angle >= 0 && d_angle < 90 {
		if e_r_angle >= 270 && e_r_angle < 360 && e_l_angle >= 0 && e_l_angle < 90 {
			e_r_angle = e_r_angle - 360
		}
	} else if d_angle >= 270 && d_angle < 360 {
		if e_r_angle >= 270 && e_r_angle < 360 && e_l_angle >= 0 && e_l_angle < 90 {
			e_l_angle = e_l_angle + 360
		}
	}

	//fmt.Println("朝向角度：", d_angle, "，左右两侧角度：", e_l_angle, e_r_angle)

	var dir_angle = float64(d_angle)

	//只有左侧角度-朝向角度小于等于90度，并且朝向角度-右侧角度也小于等于90度时，才符合出招技能的基本判断标准。
	if e_l_angle-dir_angle <= 90 && dir_angle-e_r_angle <= 90 {
		//fmt.Println("条件1达成------>左侧角度-朝向角度小于等于90度，并且朝向角度-右侧角度也小于等于90度")
		//如果左侧坐标点与敌方坐标点的角度大于等于出招角度并且右侧坐标点与敌方坐标点的角度小于等于出招角度时。则条件2达成。
		if e_l_angle >= dir_angle && e_r_angle <= dir_angle {
			//fmt.Println("条件2达成------>已在出招朝向角度范围内")
			//获取两点之间的直线距离
			distance := GetTwoPointDistance(sX, sY, eX, eY)
			//如果直线距离小于等于技能长度(出招半径)时，则条件3达成。
			if distance <= sk_l {
				//fmt.Println("条件3达成------>自己坐标与敌方坐标的两点直线距离：", distance, "，小于等于自己出招的长度：", sk_l)
				//===============================执行长方形技能覆盖敌人条件成立后的逻辑===============================
				//
				//
				//
				//===============================执行长方形技能覆盖敌人条件成立后的逻辑===============================

				return true
			}

		}
	}

	//注意：
	//1，如果上面以敌方中心点为准的判断如果不成立时，还要检测一下敌方距离出招角度最近的坐标点，判断该坐标点是否符合条件才行。就是技能区域粘到敌方边缘位置的情况，
	//2，新的判断两点之间的直线距离就要以敌方距离出招角度最近的坐标为准了。

	return false
}

//=======================================================检测敌人位置是否在自己出招范围内=======================================================

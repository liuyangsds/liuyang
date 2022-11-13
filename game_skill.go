package liuyang

//游戏技能相关

//=======================================================检测敌人位置是否在自己出招范围内=======================================================

//检测是否命中目标单位-扇形技能
//参数说明：s为self简写字母(自己)、扇面夹角度数、扇面长度(以自己为圆心的半径)、扇面朝向角度、t为target简写字母(目标)、目标鸡蛋壳半径。
//需要与检测圆与直线是否发生碰撞函数配合使用。
func CheckHit_Fan(sX, sY, sk_a, sk_l, dirAngle, tX, tY, tR float64) bool {
	//1，以目标(敌方)坐标 - 自己坐标 = 某点坐标距离0,0坐标的差值坐标，再以差值坐标获取到敌方位置距离0,0坐标的角度值
	e_angle := GetPositionAngel(tX, tY, sX, sY)
	//fmt.Println("得到目标(敌方)坐标", tX, tY, "和", sX, sY, "的差值坐标：", tX-sX, tY-sY, "与0,0坐标的角度为：", e_angle)
	n_e_angle := Float64ToInt32(e_angle)

	var angleOffset = sk_a / 2 //扇面夹角宽度的一半

	n_dirAngle := Float64ToInt32(dirAngle)

	//获取符合判断条件的角度值。参数：朝向角度、敌人距0,0的角度(-90至450之间的整数)
	i_d_angle, i_e_angle := GetJudgeAngle(n_dirAngle, n_e_angle)

	var dir_angle = float64(i_d_angle)
	leftLineAngle := dir_angle - angleOffset  //得到左侧直线角度
	rightLineAngle := dir_angle + angleOffset //得到右侧直线角度

	e_angle = float64(i_e_angle)

	//fmt.Println("出招朝向角度：", i_d_angle, "敌方所在角度：", e_angle, "，区间差值角度：", dir_angle-offsetWeight, dir_angle+offsetWeight)

	//3，判断该角度值是否 >= 自己出招角度-出招扇面宽度 && 该角度值是否 <= 自己出招角度+出招扇面宽度
	if e_angle >= leftLineAngle && e_angle <= rightLineAngle {
		//4，只有符合条件的，才能去判断自己坐标与敌方坐标的两点直线距离是否小于等于自己出招的长度
		//fmt.Println("条件1达成------>出招朝向角度：", dir_angle, "敌方所在角度：", e_angle, "符合区间差值角度：", dir_angle-offsetWeight, dir_angle+offsetWeight)
		//获取两点之间的直线距离
		distance := GetTwoPointDistance(sX, sY, tX, tY)
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

	//第一次的注意：
	//此处现在只作提醒用。
	//1，如果上面以敌方中心点为准的判断如果不成立时，还要检测一下敌方距离出招角度最近的坐标点，判断该坐标点是否符合条件才行。就是技能区域粘到敌方边缘位置的情况，
	//2，新的判断两点之间的直线距离就要以敌方距离出招角度最近的坐标为准了。
	//但是，这里千万不要使用递规方式再次调用当前函数，否则会导致无限递规的发生。
	//原因如下：
	//如果出招角度未击中敌方时，也就是顺序到达此处调用本函数方法，接着又会在本函数中继续执行到这里，这就会导致连return都没有机会执行，造成无限递规。
	//步骤：角度不在扇面中->调用本函数->角度不在扇面中->调用本函数->无限递规(因为本函数中的return false根本没机会执行)
	//所以，在使用递规时，一定要注意返回的条件。
	//另外需要注意的是：递规中的参数最好是可变的，不然参数不变的情况，就会容易忽略返回的条件，进而造成无限递规。

	//第二次的注意：
	//如果扇面没有覆盖到目标点时，还要判断一下扇面两侧的直线是否粘边到目标(鸡蛋壳)边缘。
	//fmt.Println("左右两侧直线朝向角度：", leftLineAngle, rightLineAngle)
	//以自身圆心点坐标、角度、半径获取圆边某点坐标位置
	leftEndX, leftEndY := GetRoundEdgePosition(sX, sY, leftLineAngle, sk_l)
	rightEndX, rightEndY := GetRoundEdgePosition(sX, sY, rightLineAngle, sk_l)
	//fmt.Println("扇面左侧直线终点坐标：", leftEndX, leftEndY)
	//fmt.Println("扇面右侧直线终点坐标：", rightEndX, rightEndY)

	//检测圆形物体与直线物体之间是否发生碰撞。参数中直线宽度为0，只有长方形技能时该参数才会大于0。
	leftHit := CheckCircleLineCollide(tX, tY, tR, sX, sY, leftEndX, leftEndY, 0)
	rightHit := CheckCircleLineCollide(tX, tY, tR, sX, sY, rightEndX, rightEndY, 0)
	//fmt.Println("左侧是否击中：", leftHit)
	//fmt.Println("右侧是否击中：", rightHit)

	//如果扇面的左侧直线与目标发生碰撞时
	if leftHit == true {
		return true
	}

	//如果扇面的右侧直线与目标发生碰撞时
	if rightHit == true {
		return true
	}

	return false
}

//检测是否命中目标单位-长方形技能
//参数说明：s为self简写字母(自己)、长方形的长度(以自己为圆心的半径)、长方形的宽度、长方形朝向角度、t为target简写字母(目标)、目标鸡蛋壳半径。
func CheckHit_Rectangle(sX, sY, sk_l, sk_w, dirAngle, tX, tY, tR float64) bool {
	//以自身圆心点坐标、朝向角度、技能长度(半径)获取圆边某点坐标
	endX, endY := GetRoundEdgePosition(sX, sY, dirAngle, sk_l)
	//fmt.Println("长方形终点坐标：", endX, endY)

	//得到长方形技能宽度的一半(半径)
	sk_w_r := sk_w / 2

	//检测圆形物体与直线物体之间是否发生碰撞。参数中直线宽度为0，只有长方形技能时该参数才会大于0。
	isCollide := CheckCircleLineCollide(tX, tY, tR, sX, sY, endX, endY, sk_w_r)
	if isCollide == true {
		return true
	}

	return false
}

//检测是否命中目标单位-长方形技能X3
//参数说明：s为self简写字母(自己)、长方形的长度(以自己为圆心的半径)、长方形的宽度、长方形朝向角度、左右两侧长方形的间隔角度、t为target简写字母(目标)、目标鸡蛋壳半径。
func CheckHit_Rectangle_3(sX, sY, sk_l, sk_w, dirAngle, spaceAngle, tX, tY, tR float64) bool {
	//检测中间长方形的朝向角度是否命中目标
	isCollide := CheckHit_Rectangle(sX, sY, sk_l, sk_w, dirAngle, tX, tY, tR)
	if isCollide == true {
		return true
	}

	//获得左右两侧长方形朝向角度
	leftDirAngle := dirAngle - spaceAngle  //以中间朝向角度 - 间隔角度 = 左侧长方形朝向角度
	rightDirAngle := dirAngle + spaceAngle //以中间朝向角度 + 间隔角度 = 右侧长方形朝向角度

	//检测左侧长方形的朝向角度是否命中目标
	leftIsCollide := CheckHit_Rectangle(sX, sY, sk_l, sk_w, leftDirAngle, tX, tY, tR)
	if leftIsCollide == true {
		return true
	}

	//检测右侧长方形的朝向角度是否命中目标
	rightIsCollide := CheckHit_Rectangle(sX, sY, sk_l, sk_w, rightDirAngle, tX, tY, tR)
	if rightIsCollide == true {
		return true
	}

	return false
}

//=======================================================检测敌人位置是否在自己出招范围内=======================================================

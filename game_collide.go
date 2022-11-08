package liuyang

import (
	"math"
)

//游戏碰撞相关

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

//检测两个圆形物体之间是否发生碰撞
//参数说明：s为self简写字母(自己)：x坐标、y坐标、r自己半径长度、t为target简写字母(目标)：x坐标、y坐标、r目标半径长度
func CheckCircleCollide(sX, sY, sR, tX, tY, tR float64) bool {
	//1，首先得到两个圆的圆心点的坐标。然让计算两个的圆心点之间的直线距离。
	//2，如果两圆心点的直线距离小于等于两圆半径之和时，说明两圆已发生碰撞
	//3，否则两圆心点的直线距离大于两圆半径之和时，说明两圆未发生碰撞
	//获取默认两点的直线距离，如果两圆的大小相同，则不会有问题。如果两圆大小不同时，就会发生圆重叠或距离很远时发生碰撞的问题。
	//因为默认是以左上角为起始点，所以想要获取两圆的圆心点的话，就得以自身坐标+自身半径得到自身圆心点坐标，再去判断两个圆心点的直线距离是否小于等于两圆半径之和才行。
	//如果x、y在定义时加上了半径的话，这里就可以直接这样传值了。因为此时的x、y就是圆心点。
	distance := GetTwoPointDistance(sX, sY, tX, tY)

	//fmt.Println("直线距离：", distance)

	radiusSum := sR + tR //两圆半径之和
	//如果两个圆心点的直线距离小于等于两个圆的半径之和时，则认为发生碰撞
	if distance <= radiusSum {
		return true
	}

	return false
}

//获取物体对象碰到另一物体对象后的反弹向量(参考传奇占位，A英雄碰到B英雄后的反弹向量，也就是某个方向每步要移动的值)。返回：x方向每步移动值、y方向每步移动值。不进行移动。
//参数说明：s为self简写字母(自己)、t为target简写字母(目标)
//判断说明：以左上0,0为起始点
func GetObjectReboundVector(sX, sY, tX, tY float64) (float64, float64) {
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
	//以弧度值获取余弦值和正弦值
	stepX := math.Cos(radian) //得到x方向的每步速度
	stepY := math.Sin(radian) //得到y方向的每步速度

	return stepX, stepY
}

//获取物体对象碰到边界后的反弹向量(也就是某个方向每步要移动的值)。返回：x方向每步移动值、y方向每步移动值。不进行移动。
//参数说明：边界标识->上1、左2、下3、右4
//判断说明：以左上0,0为起始点
func GetBorderReboundVector(bound_flag uint32) (float64, float64) {
	//过滤，必须要单独过滤，不然即使下面的函数会返角度0和弧度0，也会导至cos(0)=1,sin(0)=0，也就是stepX会是1，也就是整个函数会返1,0。而不是想要的0,0
	if bound_flag < bound_Up || bound_flag > bound_Right {
		return 0, 0
	}

	//获取边界标识的反弹弧度
	_, radian := GetRandomAngleRadian(bound_flag)

	stepX := math.Cos(radian) //得到x方向的每步速度
	stepY := math.Sin(radian) //得到y方向的每步速度

	return stepX, stepY
}

//检测是否碰到边界。返回：碰到的边界位置：上1、左2、下3、右4，是否碰到边界：true碰到、false未碰到
//参数说明：物体的x坐标、物体的y坐标、物体的半径、屏幕的宽、屏幕的高
//判断说明：以左上0,0为起始点
func CheckCollideBorder(x, y, r, w, h float64) (uint32, bool) {
	//临时变量
	var isCollide = false     //是否碰撞，默认false未碰撞
	var bound_flag uint32 = 0 //边界标识

	//检测边界碰撞。注意：只在气泡与边界发生碰撞时才重新获取角度、弧度、正弦、余弦。
	if x-r <= 0 { //如果 气泡坐标x+x方向速度-半径值 小于等于 0 时，说明小球已到达左侧边界

		bound_flag = bound_Left //碰撞的边界为：左
		isCollide = true        //发生碰撞

	} else if x+r >= w { //如果 气泡坐标x+x方向速度-半径值 大于等于 屏幕宽度 时，说明小球已到达右侧边界

		bound_flag = bound_Right //碰撞的边界为：右
		isCollide = true         //发生碰撞

	} else if y-r <= 0 { //如果 气泡坐标y+y方向速度-半径值 小于等于 0 时，说明小球已到达上侧边界

		bound_flag = bound_Up //碰撞的边界为：上
		isCollide = true      //发生碰撞

	} else if y+r >= h { //如果 气泡坐标y+y方向速度+半径值 大于等于 屏幕高度 时，说明小球已到达下侧边界

		bound_flag = bound_Down //碰撞的边界为：下
		isCollide = true        //发生碰撞

	}

	return bound_flag, isCollide
}

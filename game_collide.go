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
func GetBorderReboundVector(border_flag uint32) (float64, float64) {
	//过滤，必须要单独过滤，不然即使下面的函数会返角度0和弧度0，也会导至cos(0)=1,sin(0)=0，也就是stepX会是1，也就是整个函数会返1,0。而不是想要的0,0
	if border_flag < border_Up || border_flag > border_Right {
		return 0, 0
	}

	//获取边界标识的反弹弧度
	_, radian := GetRandomAngleRadian(border_flag)

	stepX := math.Cos(radian) //得到x方向的每步速度
	stepY := math.Sin(radian) //得到y方向的每步速度

	return stepX, stepY
}

//检测圆形物体是否碰到边界。返回：碰到的边界位置：上1、左2、下3、右4，是否碰到边界：true碰到、false未碰到
//参数说明：物体的x坐标、物体的y坐标、物体的半径、屏幕的宽、屏幕的高
//判断说明：以左上0,0为起始点
func CheckCollideBorder(x, y, r, w, h float64) (uint32, bool) {
	//临时变量
	var isCollide = false      //是否碰撞，默认false未碰撞
	var border_flag uint32 = 0 //边界标识

	//检测边界碰撞。注意：只在气泡与边界发生碰撞时才重新获取角度、弧度、正弦、余弦。
	if x-r <= 0 { //如果 气泡坐标x+x方向速度-半径值 小于等于 0 时，说明小球已到达左侧边界

		border_flag = border_Left //碰撞的边界为：左
		isCollide = true          //发生碰撞

	} else if x+r >= w { //如果 气泡坐标x+x方向速度-半径值 大于等于 屏幕宽度 时，说明小球已到达右侧边界

		border_flag = border_Right //碰撞的边界为：右
		isCollide = true           //发生碰撞

	} else if y-r <= 0 { //如果 气泡坐标y+y方向速度-半径值 小于等于 0 时，说明小球已到达上侧边界

		border_flag = border_Up //碰撞的边界为：上
		isCollide = true        //发生碰撞

	} else if y+r >= h { //如果 气泡坐标y+y方向速度+半径值 大于等于 屏幕高度 时，说明小球已到达下侧边界

		border_flag = border_Down //碰撞的边界为：下
		isCollide = true          //发生碰撞

	}

	return border_flag, isCollide
}

//检测圆形物体与矩形物体之间是否发生碰撞
//参数：圆心点x、圆心点y、圆的半径、矩形起始点x、矩形起始点y、矩形宽度、矩形高度
func CheckCircleRectCollide(circleX, circleY, radius, rectX, rectY, rectW, rectH float64) bool {
	//分别判断矩形4个顶点与圆心的距离是否<=圆半径；如果<=，说明碰撞成功

	//矩形起始点x - 圆点点x = a边长度
	//矩形起始点y - 圆心点y = b边长度
	//a*a + b*b <= c*c 证明矩形左上点与圆发生碰撞
	if ((rectX-circleX)*(rectX-circleX) + (rectY-circleY)*(rectY-circleY)) <= radius*radius {
		//fmt.Println("证明矩形左上点与圆发生碰撞")
		return true
	}

	//矩形起始点x + 矩形宽度 - 圆点点x = a边长度
	//矩形起始点y - 圆心点y = b边长度
	//a*a + b*b <= c*c 证明矩形右上点与圆发生碰撞
	if ((rectX+rectW-circleX)*(rectX+rectW-circleX) + (rectY-circleY)*(rectY-circleY)) <= radius*radius {
		//fmt.Println("证明矩形右上点与圆发生碰撞")
		return true
	}

	//矩形起始点x - 圆点点x = a边长度
	//矩形起始点y + 矩形高度 - 圆心点y = b边长度
	//a*a + b*b <= c*c 证明矩形左下点与圆发生碰撞
	if ((rectX-circleX)*(rectX-circleX) + (rectY+rectH-circleY)*(rectY+rectH-circleY)) <= radius*radius {
		//fmt.Println("证明矩形左下点与圆发生碰撞")
		return true
	}

	//矩形起始点x + 矩形宽度 - 圆点点x = a边长度
	//矩形起始点y + 矩形高度 - 圆心点y = b边长度
	//a*a + b*b <= c*c 证明矩形右下点与圆发生碰撞
	if ((rectX+rectW-circleX)*(rectX+rectW-circleX) + (rectY+rectH-circleY)*(rectY+rectH-circleY)) <= radius*radius {
		//fmt.Println("证明矩形右下点与圆发生碰撞")
		return true
	}

	//判断当圆心的Y坐标进入矩形内时X的位置，如果X在(rectX-radius)到(rectX+rectW+radius)这个范围内，则碰撞成功
	var minDisX float64 = 0
	//如果圆心点y 大于等于 矩形起始点y 并且圆心点y 小于等于 矩形起始点y + 矩形长度 时，
	//也就是圆在长方形的高度范围内时，则判断圆心点x 是否在矩形宽度之内。
	if circleY >= rectY && circleY <= rectY+rectH {
		//fmt.Println("圆在长方形的高度范围内时，则判断圆心点x 是否在矩形宽度之内。")
		if circleX < rectX {
			//如果圆心点x 小于 矩形起始点x 时，也就是圆在矩形左侧时，则得到左侧圆的圆心点x距离右侧矩形左侧边(矩形起始点x)的距离
			minDisX = rectX - circleX
			//fmt.Println("如果圆心点x 小于 矩形起始点x 时，也就是圆在矩形左侧时")
		} else if circleX > rectX+rectW {
			//如果圆心点x 大于 矩形起始点x + 矩形宽度 时，也就是圆在矩形右侧时，则得到右侧圆的圆心点x距离左侧矩形右侧边(矩形起始点x+矩形宽度)的距离
			minDisX = circleX - rectX - rectW
			//fmt.Println("如果圆心点x 大于 矩形起始点x + 矩形宽度 时，也就是圆在矩形右侧时")
		} else {
			//否则圆心点x 正好在矩形的左侧边与右侧边之间的位置，说明圆与矩形发生碰撞
			//fmt.Println("否则圆心点x 正好在矩形的左侧边与右侧边之间的位置，说明圆与矩形发生碰撞")
			return true
		}

		//如果圆心点x距离矩形的左侧边或右侧边的距离 小于等于 圆的半径 时，证明圆与矩形发生碰撞
		if minDisX <= radius {
			//fmt.Println("如果圆心点x距离矩形的左侧边或右侧边的距离 小于等于 圆的半径 时")
			return true
		}
	}

	//判断当圆心的X坐标进入矩形内时Y的位置，如果X在(rectY-radius)到(rectY+rectH+radius)这个范围内，则碰撞成功
	var minDisY float64 = 0
	//如果圆心点x 大于等于 矩形起始点x 并且圆心点x 小于等于 矩形起始点x + 矩形宽度 时，
	//也就是圆在长方形的宽度范围内时，则判断圆心点y 是否在矩形高度之内。
	if circleX >= rectX && circleX <= rectX+rectW {
		//fmt.Println("圆在长方形的宽度范围内时，则判断圆心点y 是否在矩形高度之内。")
		if circleY < rectY {
			//如果圆心点y 小于 矩形起始点y 时，也就是圆在矩形上侧时，则得到上侧圆的圆心点y距离下侧矩形上侧边(矩形起始点y)的距离
			minDisY = rectY - circleY
			//fmt.Println("如果圆心点y 小于 矩形起始点y 时，也就是圆在矩形上侧时")
		} else if circleY > rectY+rectH {
			//如果圆心点y 大于 矩形起始点y + 矩形高度 时，也就是圆在矩形下侧时，则得到下侧圆的圆心点y距离上侧矩形下侧边(矩形起始点y+矩形高度)的距离
			minDisY = circleY - rectY - rectH
			//fmt.Println("如果圆心点y 大于 矩形起始点y + 矩形高度 时，也就是圆在矩形下侧时")
		} else {
			//否则圆心点y 正好在矩形的上侧边与下侧边之间的位置，说明圆与矩形发生碰撞
			//fmt.Println("否则圆心点y 正好在矩形的上侧边与下侧边之间的位置，说明圆与矩形发生碰撞")
			return true
		}

		//如果圆心点y距离矩形的上侧边或下侧边的距离 小于等于 圆的半径 时，证明圆与矩形发生碰撞
		if minDisY <= radius {
			//fmt.Println("如果圆心点y距离矩形的上侧边或下侧边的距离 小于等于 圆的半径 时，证明圆与矩形发生碰撞")
			return true
		}
	}

	return false
}

//检测圆形物体与直线物体之间是否发生碰撞
//参数：圆心点x、圆心点y、圆的半径、直线起始点x、直线起始点y、直线终点x、直线终点y、直线半径(宽度的一半)
func CheckCircleLineCollide(circleX, circleY, circleRadius, startX, startY, endX, endY, lineRadius float64) bool {
	//圆心点x 与 线段起始点x 的差值
	var c_s_x = circleX - startX //圆心点x - 线段起始点x
	//圆心点y 与 线段起始点y 的差值
	var c_s_y = circleY - startY //圆心点y - 线段起始点y

	//线段终点x 与 线段起始点x 的差值
	var e_s_x = endX - startX //线段终点x - 线段起始点x
	//线段终点y 与 线段起始点y 的差值
	var e_s_y = endY - startY //线段终点y - 线段起始点y

	//线段长度：线段起始点与线段终点的两点直线距离
	var lineLength = math.Sqrt(e_s_x*e_s_x + e_s_y*e_s_y)
	//fmt.Println("线段长度：", lineLength)

	// v2.normalize()
	m_cosA := e_s_x / lineLength //角a的邻边比斜边=b/c = cosA(余弦)
	m_sinA := e_s_y / lineLength //角a的对边比斜边=a/c = sinA(正弦)

	// u = v1.dot(v2)
	// u is the vector projection length of vector v1 onto vector v2.
	//u是向量v1到向量v2的向量投影长度
	var projectLength = c_s_x*m_cosA + c_s_y*m_sinA //(圆心点x 与 线段起始点x 的差值) 乘余弦 + (圆心点y 与 线段起始点y 的差值) 乘正弦
	//fmt.Println("投影长度是：", projectLength)

	// determine the nearest point on the lineseg
	//确定线段上最近的点
	var pX float64 = 0      //投影点坐标x
	var pY float64 = 0      //投影点坐标y
	if projectLength <= 0 { //如果投影长度小于等于0时
		// p is on the left of p1, so p1 is the nearest point on lineseg
		//p在p1的左边，所以p1是线段上最近的点
		//此判断主要是为了防止圆在线段起始点以左很远的情况
		pX = startX
		pY = startY
	} else if projectLength >= lineLength { //如果投影长度大于等于线段长度时
		// p is on the right of p2, so p2 is the nearest point on lineseg
		//p在p2的右边，所以p2是线段上最近的点
		//此判断主要是为了防止圆在线段终点以右很远的情况
		pX = endX
		pY = endY
	} else { //否则投影长度大于0并且小于线段长度时

		pX = startX + projectLength*m_cosA //线段起始点x + 投影长度 乘 余弦 = 投影点坐标x
		pY = startY + projectLength*m_sinA //线段起始点y + 投影长度 乘 正弦 = 投影点坐标y

		//fmt.Println("得到投影点的坐标是：", pX, pY)
	}

	//return (circleX-x0)*(circleX-x0)+(circleY-y0)*(circleY-y0) <= radius*radius

	//根据勾股定理，求得投影点与圆心点的距离
	tempValue := (circleX-pX)*(circleX-pX) + (circleY-pY)*(circleY-pY)
	distance := math.Sqrt(tempValue)

	//如果投点与与圆心点的直线距离小于等于圆的半径+线段半径(宽度的一半)时，证明发生碰撞
	if distance <= circleRadius+lineRadius {
		return true
	}

	return false
}

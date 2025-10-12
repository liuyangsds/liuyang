# 项目开发实用小工具



天下风云出我辈，一入江湖岁月催......



刘阳——从一个小卡了咪成长到团队负责人，使用过：c、c++、java、php、swift、c#、lua、python、go等，实战项目中总会将多次使用的功能进行封装，这时一个公用函数就产生了，随着时间的推移，公用函数越来越多，为便于区分，特将函数分门别类的创建在对应文件名的文件中，我将它们统称为小工具。于是，项目开发实用小工具诞生了。



起初，此库的文件名为tools，但这个名称太通用了，总会和别人项目中已有工具包相冲突。为避免冲突，遂更名为：liuyang



## 下载与使用

### go version < 1.20 使用1.6.3

### go version >= 1.20 使用 1.7.0+

下载：
go get -u github.com/liuyangsds/liuyang

或指定版本：
go get -u github.com/liuyangsds/liuyang@v1.6.3

使用：
require (
    github.com/liuyangsds/liuyang v1.6.3
)



## 文件名与函数名对应如下：

### array.go

数组检查重复元素
func ArrayCheckRepeat([]T) bool

数组移除重复元素
func ArrayRemoveRepeat([]T) []T

数组查找元素
func ArraySearch([]T, elem) int

判断某元素是否在数组中
func InArray([]T) bool

数组排序
func ArraySort([]T, sort)

删除一个数组元素
func ArrayDelIndex([]T, index) []T

删除一个或多个数组元素
func ArrayDelIndexLength([]T, index, length) []T

收集要删除的索引值
func ArrayCollectDelIndex([]T, saveIndex) []T

批量删除切片索引
func ArrayRemove([]T, delIndex) []T

复制一个数组切片
func ArrayCopy([]T) []T

复制一个切片范围
func ArrayCopyRange([]T, start, length) []T

获取一个由int类型切片元素转成以指定符号连接的字符串
func ArrayImplode([]int, sep) string

将所有类型的数组分割成以指定符号连接的字符串
func ArrayToString(T, sep) string

数组中相同元素数量统计
func ArraySameElementCount([]T) map



### base.go

进制转换，10转2
func Base10To2(n) string

进制转换，10转8
func Base10To8(n) string

进制转换，10转16
func Base10To16(n) string

进制转换，2816转10--->可根据0b、0o、0x前缀自动判断想要的进制计算
func Base2816To10(s) int64

进制转换，2转10
func Base2To10(s) int64

进制转换，8转10
func Base8To10(s) int64

进制转换，16转10
func Base16To10(s) int64

将10进制转2进制后得到字符串(传统计算方式)
func Base10To2Custom(n) string

将10进制转8进制后得到字符串(传统计算方式)
func Base10To8Custom(n) string

将10进制数值转成16进制字符串(传统计算方式)
func Base10To16Custom(n) string

将2进制字符串转成10进制数值(传统计算方式)
func Base2To10Custom(s) int64

将8进制字符串转成10进制数值(传统计算方式)
func Base8To10Custom(s) int64

将16进制字符串转成10进制数值(传统计算方式)
func Base16To10Custom(s) int64

将10进制数值转成62进制字符串(传统计算方式)
func Base10To62Custom(n) string

将62进制字符串转成10进制数值(传统计算方式)
func Base62To10Custom(s) int64



### base64.go

将字符串进行base64编码
func Base64Encode(s) string

将base64字符串进行解码
func Base64Decode(s) string



### fixed_point.go

将普通数值转为定点数
func FixedPointEncode(f,digit) int

将定点数还原为普通数值
func FixedPointDecode(n, digit) string



### urlcode.go

将字符串进行urlencode编码
func UrlEncode(s) string

将urlencode字符串进行解码
func UrlDecode(s) string



### md5.go

将字符串进行md5加密
func Md5(s) string



### sha1.go

将字符串进行sha1加密
func Sha1(s) string



### position.go

获取两个经纬度之间的直线距离(单位：公里)
func GetLngLatDistance(Lat1, Lng1, Lat2, Lng2) float64

获取某点范围内的最小经纬度和最大经纬度
func GetPointMinMaxLngLat(lat, lng, radius) (float64, float64, float64, float64)



### regexp.go

检查用户名
func CheckUsername(s) bool

检查密码
func CheckPassword(s) bool

检查手机号
func CheckMobile(s) bool

检查邮箱
func CheckEmail(s) bool

只能输入汉字、字母、数字、下划线(1至16个)
func CheckTitle(s) bool

检查身份证码码
func CheckIDCard(s) bool

检查座机号码
func CheckTelephone(s) bool

检查网址是否合法
func CheckUrl(s) bool

检查金额
func CheckMoney(s) bool

检查(正负)金额
func CheckSMoney(s) bool

检查小数
func CheckDecimal(s) bool

检查(正负)小数
func CheckSDecimal(s) bool

检查小数(6位)
func CheckDecimal6(s) bool

检查(正负)小数(6位)
func CheckSDecimal6(s) bool

检查纯数字
func CheckNumber(s) bool

检查(正负)纯数字
func CheckSNumber(s) bool

检查用户id
func CheckUid(s) bool

检查版本号
func CheckVersion(s) bool

检查坐标
func CheckPosition(s) bool

检查字符串是否为单个或以英文逗号连接的多个字符串
func CheckStrStr(s) bool

检查字符串是否为多个整型数字字符串
func CheckNumberNumber(s) bool

检查字符串是否为json字符串
func CheckJson(s) bool

检查json字符串是否已加上转义符
func CheckJsonEscape(s) bool



### http.go

http请求,get方式
func HttpRequestGet(url) (string, error)

http请求,post方式
func HttpRequestPost(url, map) (string, error)

http请求,post方式->json数据传参
func HttpRequestPostJson(url, []byte) (string, error)



### math.go

以角度值获取弧度值，参数：角度
func GetRadian(angle float64) float64

以弧度值获取角度值，参数：弧度
func GetAngle(radian float64) float64

以弧度值获取x与y的向量值，参数：弧度
func GetVector(radian float64) (float64, float64)

获取两点之间的直线距离
func GetTwoPointDistance(aX, aY, bX, bY float64) float64

获取目标坐标距自己坐标的角度值(0-359)和对应的弧度值
参数说明：t为target简写字母(目标)，s为self简写字母(自己)
func GetPositionAngel(tX, tY, sX, sY float64) (float64, float64)

获取反正切函数得到的角度值(0-359)和对应的弧度值
func GetAtan2AngleRadian(y, x float64) (float64, float64)

以圆心点坐标、角度、半径获取圆边某点坐标
func GetRoundEdgePosition(x0, y0, angle, r float64) (float64, float64)

以出招朝向角度获取左右两侧面向的角度(0-359之间的正整数)
func GetDirectionSideTwoAngle(dirAngle int) (int, int, int)

以出招朝向角度获取敌方最近距离面向的角度
返回：出招朝向角度、敌方距离0,0的角度、敌方距离出招角度最近的角度(0-359之间的正整数)
func GetDirectionSideEnemyAngle(dirAngle, enemyAngle int) (int, int, int)

获取正整数角度值。返回：0-359之间的正整数
func GetPositiveAngle(angle int) int

获取机率触发的可能性。参数：机率值0-100，返回：机率达成返true、未达成返false
func GetOdds(n int) bool

获取物体对象碰到另一物体对象后的随机向量(参考传奇占位，A英雄碰到B英雄后的反弹向量，也就是某个方向每步要移动的值)。返回：随机角度、弧度、x方向每步移动值、y方向每步移动值
参数说明：s为self简写字母(自己)、t为target简写字母(目标)
判断说明：以左上0,0为起始点
func GetObjectRandomVector(sX, sY, tX, tY float64) (int, float64, float64, float64)

获取物体对象碰到边界后的随机向量(也就是某个方向每步要移动的值)。返回：随机角度、弧度、x方向每步移动值、y方向每步移动值
参数说明：边界标识->上1、左2、下3、右4
判断说明：以左上0,0为起始点
GetBorderRandomVector(border_flag uint32) (int, float64, float64, float64)

获取物体对象碰到障碍物后的回弹向量(也就是某个方向每步要移动的值)。返回：回弹角度、弧度、x方向每步移动值、y方向每步移动值
func GetBorderReboundVector(angle int) (int, float64, float64, float64)

获取物体对象碰到边界后的反射向量(也就是某个方向每步要移动的值)。返回： 反射角度、弧度、x方向每步移动值、y方向每步移动值
参数说明：边界标识->上1、左2、下3、右4、正弦值(y向量)、余弦值(x向量)
func GetBorderReflectVector(border_flag uint32, stepY, stepX float64) (float64, float64, float64, float64)

获取物体对象碰到另一物体对象后的随机角度(参考传奇占位，A英雄碰到B英雄后的反弹)。返回：角度值
参数说明：s为self简写字母(自己)、t为target简写字母(目标)
判断说明：坐标以左上0,0为起始点、角度以右起顺时针0到360
func GetObjectRandomAngle(sX, sY, tX, tY float64)

获取物体碰到边框后的随机角度。参数：碰到的边界标识。返回：角度值
判断说明：坐标以左上0,0为起始点、角度以右起顺时针0到360
func GetBorderRandomAngle(border_flag uint32)

获取回弹角度。不用考虑碰到的目标是圆形物体还是墙壁，都是直接回弹
func GetReboundAngle(angle int) int

获取物体对象碰到边界后的反射角度
func GetReflectAngle(border_flag uint32, stepY, stepX float64) (float64, float64)

获取圆的两侧切线坐标--->点p到圆的右侧切线坐标、点p到圆的左侧切线坐标、点p与圆的切线->与点p与圆心直线的夹角度数、是否获取到切线坐标与角度值
说明：点p如果在圆内时是得不到切线坐标的，所以会返：0,0,0,0,0,false
说明：点p如果在圆外时是可以得到切线坐标的，所以会返：坐标1、坐标2、角度值、true
:param px: 圆外点P横坐标
:param py: 圆外点P纵坐标
:param cx: 圆心横坐标
:param cy: 圆心纵坐标
:param r:  圆半径
:return:   过点P的俩条切线坐标
func GetCircleTangentCoor(px, py, cx, cy, r float64) (float64, float64, float64, float64, float64, bool)

检测目标角度是否在朝向角度与偏移角度区间内。返回：左侧偏移后的角度、右侧偏移后的角度、是否在左侧到右侧角度内：true是、false否
参数：朝向角度、左右浮动(偏移)角度、目标角度
样例：朝向135度，偏移区间30度、目标120度--->左侧角度105(135-30)度、右侧角度165(135+30)度--->在区间角度内
样例：朝向356度，偏移区间30度、目标12度--->左侧角度326(356-30)度、右侧角度26((356+30)%360)度--->在区间角度内
样例：朝向12度，偏移区间30度、目标356度--->左侧角度342((12-30)+360)度、右侧角度42(12+30)度--->在区间角度内
func CheckTargetAtOffsetAngle(dirAngle int, offsetAngle int, targetAngle int) (int, int, bool)

获取两个数之间的最大值
func GetTwoNubmerMax(a, b T) int

获取三个数之间的最大值
func GetThreeNubmerMax(a, b, c T) int



### game_collide.go

检测两个圆形物体之间是否发生碰撞
参数说明：s为self简写字母(自己)：x坐标、y坐标、r自己半径长度、t为target简写字母(目标)：x坐标、y坐标、r目标半径长度
func CheckCircleCollide(sX, sY, sR, tX, tY, tR float64) bool

检测圆形物体是否碰到边界。返回：碰到的边界位置：上1、左2、下3、右4，是否碰到边界：true碰到、false未碰到
参数说明：物体的x坐标、物体的y坐标、物体的半径、屏幕的宽、屏幕的高
判断说明：以左上0,0为起始点
func CheckCollideBorder(x, y, r, w, h float64) (uint32, bool)

检测圆形物体与矩形物体之间是否发生碰撞
参数：圆心点x、圆心点y、圆的半径、矩形起始点x、矩形起始点y、矩形宽度、矩形高度
func CheckCircleRectCollide(circleX, circleY, radius, rectX, rectY, rectW, rectH float64) bool

检测圆形物体与直线物体之间是否发生碰撞
参数：圆心点x、圆心点y、圆的半径、直线起始点x、直线起始点y、直线终点x、直线终点y、直线半径(宽度的一半)
func CheckCircleLineCollide(circleX, circleY, circleRadius, startX, startY, endX, endY, lineRadius float64) bool



### game_coord.go

获取某个索引的对称坐标偏移位置。如：(0)-385、(1)-275、(2)-165、(3)-55、(4)55、(5)165、(6)275、(7)385
参数：总个数、每个单位的直径、单位之间的间隔、每个单位所在位置(数据集遍历时的key，0起始)
func GetKeyOffsetPosition(total int, diameter float64, space float64, key int) float64

//获取某个索引所在对称坐标位置。
第一行：(0)-110,-110	(1)--->0,-110	(2)--->110,-110
第二行：(3)-110,0		(4)--->0,0		(5)--->110,0
第三行：(6)-110,110		(7)--->0,110	(8)--->110,110
参数：总个数、每个单位的直径、单位之间的间隔、每个单位所在位置(数据集遍历时的key，0起始)、每行个数(小于1时取总个数)、是否完美站位(默认false)
func GetKeyCoordPosition(total int, diameter float64, space float64, key int, rowNum int, perfect bool) (float64, float64)



### game_coordOffset.go

获取坐标偏移地图中两个格子之间的距离。返回：以自身为起始的相距格子数量。
如：0：与自身相距0格、1：与自身相距1格，也就是相邻的6个格子、2：与自身相距2格，也就是中间会隔一格(6个邻居)。
参数：自身列、自身行、目标列、目标行、偏移坐标的类型：0尖朝上奇数行偏移、1尖朝上偶数行偏移、2边朝上奇数列偏移、3边朝上偶数列偏移、其他值视为0。
func GetCoordOffsetGridDistance(sX, sY, tX, tY, offsetType int) int

检测偏移坐标中(六边形地图)两个坐标点的距离是否在某一距离内(小于等于指定距离)
参数：自身列、自身行、目标列、目标行、距离、偏移坐标的类型：0尖朝上奇数行偏移、1尖朝上偶数行偏移、2边朝上奇数列偏移、3边朝上偶数列偏移、其他值视为0。
func CheckOffsetCoordDistance(sX, sY, tX, tY, distance, offsetType int) bool

检测偏移坐标中(六边形地图)两个坐标点的距离是否在某一距离区间内(大于等于最小距离并且小于等于最大距离)
参数：自身列、自身行、目标列、目标行、距离最小值、距离最大值、偏移坐标的类型：0尖朝上奇数行偏移、1尖朝上偶数行偏移、2边朝上奇数列偏移、3边朝上偶数列偏移、其他值视为0。
func CheckOffsetCoordDistanceRange(sX, sY, tX, tY, distanceMin, distanceMax, offsetType int) bool

检测偏移坐标中(六边形地图)两个坐标点的距离是否等于某一距离(等于指定距离)
参数：自身列、自身行、目标列、目标行、距离、偏移坐标的类型：0尖朝上奇数行偏移、1尖朝上偶数行偏移、2边朝上奇数列偏移、3边朝上偶数列偏移、其他值视为0。
func CheckOffsetCoordDistanceEqual(sX, sY, tX, tY, distance, offsetType int) bool

获取立方体坐标的周围的邻居位置(六个格子)
顺序：左起开始
func GetCubeNeighborArr(x, y, z int) [][]int

获取偏移坐标(尖朝上-奇数行偏移)周围的邻居位置(六个格子)
参数：列位置、行位置、列数最大值、行数最大值
返回：某位置周围合法的位置(除去<0或超出最大值或等于列最大值时的奇数行最后一个位置)
顺序：左起开始
func GetOffsetTipOddNeighborArr(x, y, X_max, Y_max int) [][]int

获取偏移坐标(尖朝上-奇数行偏移)对应位置的索引。返回值：索引、是否找到对应索引
参数：偏移坐标中的某位置(中心坐标)x、某位置(中心坐标)y、最大列、最大行、要获取对应索引的目标位置x、目标位置y
func GetOffsetTipOddPositionIndex(x, y, X_max, Y_max, targetX, targetY int) (int, bool)

立方体坐标转偏移坐标-尖朝上-奇数行偏移
func CoordCubeToOffsetTipOdd(x, z int) (int, int)

偏移坐标转立方体坐标-尖朝上-奇数行偏移
func CoordOffsetToCubeTipOdd(col, row int) (int, int, int)

立方体坐标转偏移坐标-尖朝上-偶数行偏移
func CoordCubeToOffsetTipEven(x, z int) (int, int)

偏移坐标转立方体坐标-尖朝上-偶数行偏移
func CoordOffsetToCubeTipEven(col, row int) (int, int, int)

立方体坐标转偏移坐标-边朝上-奇数列偏移
func CoordCubeToOffsetSideOdd(x, z int) (int, int)

偏移坐标转立方体坐标-边朝上-奇数列偏移
func CoordOffsetToCubeSideOdd(col, row int) (int, int, int)

立方体坐标转偏移坐标-边朝上-偶数列偏移
func CoordCubeToOffsetSideEven(x, z int) (int, int)

偏移坐标转立方体坐标-边朝上-偶数列偏移
func CoordOffsetToCubeSideEven(col, row int) (int, int, int)

获取偏移坐标(尖朝上-奇数行偏移)周围的邻居位置(六个格子)所对应的角度
func GetOffsetTipOddNeighborAngle(sX, sY, tX, tY, X_max, Y_max int) int



### game_skill.go

检测是否命中目标单位-扇形技能
参数说明：s为self简写字母(自己)、扇面夹角度数、扇面长度(以自己为圆心的半径)、扇面朝向角度、t为target简写字母(目标)、目标鸡蛋壳半径。
需要与检测圆与直线是否发生碰撞函数配合使用。
func CheckHit_Fan(sX, sY, sk_a, sk_l, dirAngle, tX, tY, tR float64) bool

检测是否命中目标单位-长方形技能
参数说明：s为self简写字母(自己)、长方形的长度(以自己为圆心的半径)、长方形的宽度、长方形朝向角度、t为target简写字母(目标)、目标鸡蛋壳半径。
func CheckHit_Rectangle(sX, sY, sk_l, sk_w, dirAngle, tX, tY, tR float64) bool

检测是否命中目标单位-长方形技能X3
参数说明：s为self简写字母(自己)、长方形的长度(以自己为圆心的半径)、长方形的宽度、长方形朝向角度、左右两侧长方形的间隔角度、t为target简写字母(目标)、目标鸡蛋壳半径。
func CheckHit_Rectangle_3(sX, sY, sk_l, sk_w, dirAngle, spaceAngle, tX, tY, tR float64) bool

package liuyang

import "math"

// 地球半径，单位：公里
const earthRadius = 6371 //地球平均半径，单位为公里
//const earthRadius = 6378.137; //地球赤道半径，单位为公里

// 获取两个经纬度之间的直线距离(单位：公里)
// 样例1
// distance := test_new.LngLatDistance(39.9042, 116.4074, 31.2304, 121.4737)//北京到上海的直线距离
// 打印
// 1067.310170927128
// 样例2
// distance := test_new.LngLatDistance(39.9, 116.4, 40.7, -74.0)//北京到纽约的直线距离
// 打印
// 10991.120256055912
func GetLngLatDistance(Lat1, Lng1, Lat2, Lng2 float64) float64 {
	// 将角度转换为弧度
	lat1 := degreesToRadians(Lat1)
	lng1 := degreesToRadians(Lng1)
	lat2 := degreesToRadians(Lat2)
	lng2 := degreesToRadians(Lng2)

	// 计算纬度和经度的差值
	dLat := lat2 - lat1
	dLng := lng2 - lng1

	// Haversine公式
	aa := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLng/2)*math.Sin(dLng/2)
	cc := 2 * math.Atan2(math.Sqrt(aa), math.Sqrt(1-aa))

	// 计算距离（地球半径 * 弧度）
	distance := earthRadius * cc

	//返回单位：公里
	return distance
}

// 将角度转换为弧度
func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// 获取某点范围内的最小经纬度和最大经纬度(如同坐标点一样的左上角和右下角两个点)
// 参数1：纬度、参数2：经度
// 本函数存在的意义为：方便查询mysql，因为hotel表中有经度和纬度字段，想要获取某个酒店附近其他酒店时，
// 样例
// lat := 45.77188995878113  //纬度
// lng := 131.02478238034496 //经度
// minLat, minLng, maxLat, maxLng := test_new.GetPointMinMaxLngLat(lat, lng, 1)
// 打印
// 45.76289674272194
// 131.0118891901371
// 45.78088317484032
// 131.0376755705528
func GetPointMinMaxLngLat(lat, lng, radius float64) (float64, float64, float64, float64) {
	//地球赤道半径6378.137千米，极半径6356.752千米，平均半径约6371千米，赤道周长大约为40076千米。地球上71%为海洋，29%为陆地。
	latOffset := 180 / math.Pi * radius / earthRadius
	lngOffset := latOffset / math.Cos(lat*math.Pi/180)

	minLat := lat - latOffset //最小纬度
	minLng := lng - lngOffset //最小经度
	maxLat := lat + latOffset //最大纬度
	maxLng := lng + lngOffset //最大经度

	return minLat, minLng, maxLat, maxLng
}

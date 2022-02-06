package coordsconvert

import "math"

const (
	x_PI = 3.14159265358979324 * 3000.0 / 180.0
)

func BD09ToGCJ02(x, y float64) (float64, float64) {
	tmp_x := x - 0.0065
	tmp_y := y - 0.006
	z := math.Sqrt(tmp_x*tmp_x+tmp_y*tmp_y) - 0.00002*math.Sin(tmp_y*x_PI)
	theta := math.Atan2(tmp_y, tmp_x) - 0.000003*math.Cos(tmp_x*x_PI)
	ret_x := z * math.Cos(theta)
	ret_y := z * math.Sin(theta)
	return ret_x, ret_y
}

func GCJ02ToBD09(x, y float64) (float64, float64) {
	z := math.Sqrt(x*x+y*y) + 0.00002*math.Sin(y*x_PI)
	theta := math.Atan2(y, x) + 0.000003*math.Cos(x*x_PI)
	ret_x := z*math.Cos(theta) + 0.0065
	ret_y := z*math.Sin(theta) + 0.006
	return ret_x, ret_y
}

func WGS84ToBD09(x, y float64) (float64, float64) {
	return GCJ02ToBD09(WGS84ToGCJ02(x, y))
}

func BD09ToWGS84(x, y float64) (float64, float64) {
	return GCJ02ToWGS84(BD09ToGCJ02(x, y))
}

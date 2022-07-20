package godecimal

import (
	"fmt"
	"github.com/dtapps/go-library/utils/gostring"
	"math"
	"strconv"
)

// Decimal 四舍五入
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// Round 四舍五入
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

// RoundYInt64 四舍五入
func RoundYInt64(y int64, n int) float64 {
	return Round(float64(y/100), n)
}

// RoundYString 四舍五入
func RoundYString(y string, n int) float64 {
	return Round(gostring.ToFloat64(y)/100, n)
}

// Multiply 相乘
func Multiply(y, x float64) float64 {
	return Round(y*x, 2)
}

// PddCouponAmount 优惠券金额
func PddCouponAmount(y int64) float64 {
	return Round(float64(y)/100, 2)
}

// PddCouponProportion 拼多多佣金比率
func PddCouponProportion(y int64) float64 {
	return Round(float64(y)/10, 2)
}

// PddGoodsOriginalPrice 拼多多商品原价
func PddGoodsOriginalPrice(y int64) float64 {
	return Round(float64(y)/100, 2)
}

// PddGoodsPrice 拼多多商品券后价
func PddGoodsPrice(y, x float64) float64 {
	return Round(y-x, 2)
}

// PddCommission 拼多多佣金
func PddCommission(y, x float64) float64 {
	return Round((y*x)/100, 2)
}

// TbCouponAmount 淘宝优惠券金额
func TbCouponAmount(y int64) float64 {
	return Round(float64(y), 2)
}

// TbCouponProportion 淘宝佣金比率
func TbCouponProportion(y string) float64 {
	return Round(gostring.ToFloat64(y)/100, 2)
}

// TbGoodsOriginalPrice 淘宝商品原价
func TbGoodsOriginalPrice(y string) float64 {
	return Round(gostring.ToFloat64(y), 2)
}

// TbGoodsPrice 淘宝商品券后价
func TbGoodsPrice(y, x float64) float64 {
	return Round(y-x, 2)
}

// TbCommission 淘宝佣金
func TbCommission(y, x float64) float64 {
	return Round((y*x)/100, 2)
}

// WmCouponAmount 小商店优惠券金额
func WmCouponAmount(y string) float64 {
	return Round(gostring.ToFloat64(y)/100, 2)
}

// WmCouponProportion 小商店佣金比率
func WmCouponProportion(y int64) float64 {
	return Round(float64(y)/100, 2)
}

// WmCommission 小商店佣金
func WmCommission(y int64) float64 {
	return Round(float64(y)/100, 2)
}

// WmGoodsOriginalPrice 小商店商品原价
func WmGoodsOriginalPrice(y int64) float64 {
	return Round(float64(y)/100, 2)
}

// WmGoodsPrice 小商店商品券后价
func WmGoodsPrice(y int64) float64 {
	return Round(float64(y)/100, 2)
}

// JdCommission 京东佣金
func JdCommission(y, x float64) float64 {
	return Round((y*x)/100, 2)
}

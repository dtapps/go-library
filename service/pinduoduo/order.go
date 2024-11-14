package pinduoduo

import (
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
)

var (
	OrderStatusType = []int64{0, 1, 2, 3, 4, 5, 10}
	OrderStatusDesc = []string{"已支付", "已成团", "确认收货", "审核失败(不可提现)", "已经结算", "已处罚"}
)

// GetOrderStatusDesc 订单状态
func GetOrderStatusDesc(Type int64) (desc string) {
	for i, v := range OrderStatusType {
		if v == Type {
			desc = OrderStatusDesc[i]
			break
		}
	}
	return
}

var orderSubsidyMap = map[int64]string{
	0: "非补贴订单",
	1: "千万补贴",
	2: "社群补贴",
	3: "多多星选",
	4: "品牌优选",
	5: "千万神券",
	7: "佣金翻倍补贴",
	8: "拼团享多多",
}

// GetOrderSubsidyDesc 订单补贴类型
func GetOrderSubsidyDesc(orderSubsidy int64) (desc string) {
	if desc, ok := orderSubsidyMap[orderSubsidy]; ok {
		return desc
	}
	return "" // 或者返回一个默认值或错误信息
}

var orderTypeMap = map[int64]string{
	0:   "单品",
	1:   "红包",
	2:   "领券页推荐",
	3:   "主题",
	4:   "手机商城",
	6:   "拼团后推荐",
	7:   "今日爆款",
	8:   "品牌清仓",
	9:   "1.9包邮",
	10:  "全店关联",
	11:  "PC商城",
	13:  "大转盘锁佣",
	16:  "支付新用户锁佣",
	18:  "CPA拉新锁佣",
	22:  "果园",
	25:  "跨店",
	51:  "商详推荐",
	52:  "店铺券",
	54:  "大转盘",
	55:  "大转盘",
	56:  "大转盘",
	57:  "挽回推荐",
	61:  "频道活动",
	64:  "跨店关联",
	68:  "活动推广",
	69:  "拼团后推荐",
	72:  "新人红包",
	74:  "拼团后推荐",
	76:  "家装",
	77:  "刮刮卡",
	78:  "大转盘",
	80:  "直播",
	83:  "挽回推荐",
	84:  "直播全店关联",
	85:  "拼多多品类榜单",
	88:  "频道活动",
	89:  "频道活动",
	90:  "拼团后推荐",
	91:  "浏览关联",
	93:  "砸金蛋",
	94:  "充值中心",
	95:  "频道活动",
	101: "品牌黑卡",
	103: "百亿补贴频道",
	104: "内购清单频道",
	105: "超级红包",
	200: "拼团模式",
	201: "拼团模式关联",
}

// GetOrderTypeDesc 根据订单类型获取描述
func GetOrderTypeDesc(orderType int64) string {
	if desc, ok := orderTypeMap[orderType]; ok {
		return desc
	}
	return "" // 或者返回一个默认值或错误信息
}

var (
	OrderBandanRiskConsultType = []int64{-1, 0, 1}
	OrderBandanRiskConsultDesc = []string{"未出结果", "不是代购订单", "是代购订单"}
)

// GetOrderBandanRiskConsultDesc 预判断是否为代购订单
func GetOrderBandanRiskConsultDesc(Type int64) (desc string) {
	for i, v := range OrderBandanRiskConsultType {
		if v == Type {
			desc = OrderBandanRiskConsultDesc[i]
			break
		}
	}
	return
}

var (
	OrderIsDirectType = []int64{0, 1}
	OrderIsDirectDesc = []string{"否", "是"}
)

// GetOrderIsDirectDesc 下单场景类型
func GetOrderIsDirectDesc(Type int64) (desc string) {
	for i, v := range OrderIsDirectType {
		if v == Type {
			desc = OrderIsDirectDesc[i]
			break
		}
	}
	return
}

var (
	OrderPriceCompareStatusType = []int64{0, 1}
	OrderPriceCompareStatusDesc = []string{"正常", "比价"}
)

// GetOrderPriceCompareStatusDesc 下单场景类型
func GetOrderPriceCompareStatusDesc(Type int64) (desc string) {
	for i, v := range OrderPriceCompareStatusType {
		if v == Type {
			desc = OrderPriceCompareStatusDesc[i]
			break
		}
	}
	return
}

func GetOrderSalesTipParseInt64(salesTip string) int64 {
	// 去掉字符串中的 "+"
	salesTip = strings.Replace(salesTip, "+", "", -1)

	// 如果包含 "万"（包括 "万+" 或 "万"）
	if strings.Contains(salesTip, "万") {
		// 去掉 "万"
		salesTip = strings.Replace(salesTip, "万", "", -1)

		// 如果有小数部分，直接转换为 decimal 类型并乘以 10000
		if val, err := decimal.NewFromString(salesTip); err == nil {
			// 如果原始字符串含有 "+"，意味着是一个 "万+" 的情况，稍微提高估值
			if strings.Contains(salesTip, "+") {
				// 向上调整为 1.1 倍，或者其他比例，根据业务逻辑来调整
				return val.Mul(decimal.NewFromInt(10000)).Mul(decimal.NewFromFloat(1.1)).IntPart()
			}
			return val.Mul(decimal.NewFromInt(10000)).IntPart()
		}
	}

	// 如果是普通数字（没有 "万"）
	if val, err := strconv.ParseInt(salesTip, 10, 64); err == nil {
		return val
	}

	// 无法解析时返回 0
	return 0
}

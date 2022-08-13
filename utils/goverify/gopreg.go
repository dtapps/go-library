package goverify

import (
	"regexp"
)

const (
	Mobile   = "mobile"   // 中国移动
	Unicom   = "unicom"   // 中国联通
	Telecom  = "telecom"  // 中国电信
	Broadnet = "broadnet" // 中国广电
	Virtual  = "virtual"  // 虚拟
)

// ChinaMobile 验证手机号码
// https://baike.baidu.com/item/%E7%94%B5%E8%AF%9D%E5%8F%B7%E7%A0%81/1417271
// https://zh.wikipedia.org/wiki/%E4%B8%AD%E5%9B%BD%E5%A4%A7%E9%99%86%E7%A7%BB%E5%8A%A8%E7%BB%88%E7%AB%AF%E9%80%9A%E4%BF%A1%E5%8F%B7%E7%A0%81#cite_note-2
// https://www.qqzeng.com/tongji.html
func ChinaMobile(number string) (status bool, operator string) {
	status = ChinaMobileNumber(number) // 中国移动运营商
	if status {
		return status, Mobile
	}
	status = ChinaUnicomNumber(number) // 中国联通运营商
	if status {
		return status, Unicom
	}
	status = ChinaTelecomNumber(number) // 中国电信运营商
	if status {
		return status, Telecom
	}
	status = ChinaBroadnetNumber(number) // 中国广电运营商
	if status {
		return status, Broadnet
	}
	status = ChinaVirtualNumber(number) // 虚拟运营商
	if status {
		return status, Virtual
	}
	return
}

// ChinaMobileNumber 验证中国移动手机号码
// https://www.qqzeng.com/tongji.html
// 移动：134 135 136 137 138 139 147 148 150 151 152 157 158 159 172 178 182 183 184 187 188 195 197 198
func ChinaMobileNumber(number string) bool {
	regular := "^[1](([3][4-9])|([4][7-8])|([5][0-2,7-9])|([7][2,8])|([8][2-4,7-8])|([9][5,7-8]))[0-9]{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(number)
}

// ChinaUnicomNumber 验证中国联通手机号码
// https://www.qqzeng.com/tongji.html
// 联通：130 131 132 145 146 155 156 166 167 171 175 176 185 186 196
func ChinaUnicomNumber(number string) bool {
	regular := "^[1](([3][0-2])|([4][5-6])|([5][5-6])|([6][6-7])|([7][1,5-6])|([8][5-6])|([9][6]))[0-9]{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(number)
}

// ChinaTelecomNumber 验证中国电信手机号码
// https://www.qqzeng.com/tongji.html
// 电信：133 149 153 173 174 177 180 181 189 190 191 193 199
func ChinaTelecomNumber(number string) bool {
	regular := "^[1](([3][3])|([4][9])|([5][3])|([7][3-4,7])||([8][0-1,9])|([9][0-1,3,9]))[0-9]{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(number)
}

// ChinaBroadnetNumber 验证中国广电手机号码
// https://www.qqzeng.com/tongji.html
// 广电：192
func ChinaBroadnetNumber(number string) bool {
	regular := "^[1](([9][2]))[0-9]{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(number)
}

// ChinaVirtualNumber 验证虚拟运营商手机号码
// https://www.qqzeng.com/tongji.html
// 移动/联通/电信: 162 165 167 170 171
func ChinaVirtualNumber(number string) bool {
	regular := "^[1](([6][2,5,7])|([7][0-1]))[0-9]{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(number)
}

// IdCard 验证身份证号码
func IdCard(idCard string) bool {
	regular := "^[1-9]\\d{7}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}$|^[1-9]\\d{5}[1-9]\\d{3}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}([0-9]|X)$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(idCard)
}

// Email 验证邮箱号码
func Email(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

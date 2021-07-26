package v20210726

import (
	"regexp"
)

// VerifyMobile 验证手机号码
// 移动：134 135 136 137 138 139 147 150 151 152 157 158 159 178 182 183 184 187 188 198
// 联通：130 131 132 145 155 156 166 171 175 176 185 186
// 电信：133 149 153 173 177 180 181 189 199
// 虚拟运营商: 170 195
func VerifyMobile(mobile string) bool {
	regular := "^[1](([3][0-9])|([4][5-9])|([5][0-3,5-9])|([6][5,6])|([7][0-8])|([8][0-9])|([9][1,5,8-9]))[0-9]{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

// VerifyIdCard 验证身份证号码
func VerifyIdCard(idCard string) bool {
	regular := "^[1-9]\\d{7}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}$|^[1-9]\\d{5}[1-9]\\d{3}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}([0-9]|X)$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(idCard)
}

// VerifyEmail 验证邮箱号码
func VerifyEmail(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

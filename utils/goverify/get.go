package goverify

// 运营商名称
var typeName = map[string]string{
	Mobile:   "移动",
	Unicom:   "联通",
	Telecom:  "电信",
	Broadnet: "广电",
	Virtual:  "虚拟",
}

// GetTypeName 获取运营商类型名称
func GetTypeName(Type string) string {
	return typeName[Type]
}

// GetType 获取运营商类型
func GetType(name string) (Type string) {
	switch name {
	case "移动", "移动运营商", "中国移动", "中国移动运营商":
		return Mobile
	case "联通", "联通运营商", "中国联通", "中国联通运营商":
		return Unicom
	case "电信", "电信运营商", "中国电信", "中国电信运营商":
		return Telecom
	case "广电", "广电运营商", "中国广电", "中国广电运营商":
		return Broadnet
	case "虚拟", "虚拟运营商":
		return Virtual
	}
	return Type
}

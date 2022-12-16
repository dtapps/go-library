package gostring

func GetDefault(key, defVal any) any {
	if key != nil {
		return key
	}
	return defVal
}

func GetStringDefault(key, defVal string) string {
	if key != "" {
		return key
	}
	return defVal
}

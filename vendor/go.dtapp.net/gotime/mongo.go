package gotime

// Bson mongoDB
func (p Pro) Bson() string {
	return p.Now().String()
}

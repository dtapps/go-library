package duuid

import (
	"github.com/google/uuid"
)

// GenUUID 获取唯一ID
func GenUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}

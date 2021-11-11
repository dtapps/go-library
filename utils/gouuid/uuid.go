package gouuid

import (
	"github.com/google/uuid"
)

// GetUuId 获取唯一ID
func GetUuId() string {
	u, _ := uuid.NewRandom()
	return u.String()
}

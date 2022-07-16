package dorm

import "github.com/upper/db/v4"

// UpperClient
// https://upper.io/
type UpperClient struct {
	Db *db.Session // 驱动
}

package dorm

import "github.com/upper/db/v4"

// UpperClient
// https://upper.io/
type UpperClient struct {
	db *db.Session // 驱动
}

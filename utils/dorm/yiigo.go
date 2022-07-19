package dorm

import (
	"github.com/jmoiron/sqlx"
	"github.com/shenghui0779/yiigo"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConfigYiiGoClient struct {
	Dns  string
	Addr string
}

// YiiGoClient
// https://github.com/shenghui0779/yiigo
type YiiGoClient struct {
	Db        *sqlx.DB
	MDb       *mongo.Client
	RedisPool *yiigo.RedisConn
	config    *ConfigYiiGoClient
}

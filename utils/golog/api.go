package golog

import (
	"go.dtapp.net/library/utils/goip"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"log"
	"os"
	"runtime"
	"strings"
	"unicode/utf8"
)

type ApiConfig struct {
	Db        *gorm.DB // 驱动
	TableName string   // 表名
}

// Api 接口
type Api struct {
	db        *gorm.DB // pgsql数据库
	tableName string   // 日志表名
	insideIp  string   // 内网ip
	hostname  string   // 主机名
	goVersion string   // go版本
}

// NewApi 创建接口实例化
func NewApi(config *ApiConfig) *Api {
	app := &Api{}
	if config.Db == nil {
		panic("驱动不正常")
	}
	if config.TableName == "" {
		panic("表名不能为空")
	}
	hostname, _ := os.Hostname()

	app.db = config.Db
	app.tableName = config.TableName
	app.hostname = hostname
	app.insideIp = goip.GetInsideIp()
	app.goVersion = strings.TrimPrefix(runtime.Version(), "go")

	app.AutoMigrate()
	return app
}

// ApiPostgresqlLog 结构体
type ApiPostgresqlLog struct {
	LogId                 uint           `gorm:"primaryKey" json:"log_id"`                   //【记录】编号
	RequestTime           TimeString     `gorm:"index" json:"request_time"`                  //【请求】时间
	RequestUri            string         `gorm:"type:text" json:"request_uri"`               //【请求】链接
	RequestUrl            string         `gorm:"type:text" json:"request_url"`               //【请求】链接
	RequestApi            string         `gorm:"type:text;index" json:"request_api"`         //【请求】接口
	RequestMethod         string         `gorm:"type:text;index" json:"request_method"`      //【请求】方式
	RequestParams         datatypes.JSON `gorm:"type:jsonb" json:"request_params"`           //【请求】参数
	RequestHeader         datatypes.JSON `gorm:"type:jsonb" json:"request_header"`           //【请求】头部
	ResponseHeader        datatypes.JSON `gorm:"type:jsonb" json:"response_header"`          //【返回】头部
	ResponseStatusCode    int            `gorm:"type:bigint" json:"response_status_code"`    //【返回】状态码
	ResponseBody          datatypes.JSON `gorm:"type:jsonb" json:"response_body"`            //【返回】内容
	ResponseContentLength int64          `gorm:"type:bigint" json:"response_content_length"` //【返回】大小
	ResponseTime          TimeString     `gorm:"index" json:"response_time"`                 //【返回】时间
	SystemHostName        string         `gorm:"type:text" json:"system_host_name"`          //【系统】主机名
	SystemInsideIp        string         `gorm:"type:text" json:"system_inside_ip"`          //【系统】内网ip
	GoVersion             string         `gorm:"type:text" json:"go_version"`                //【程序】Go版本
}

// AutoMigrate 自动迁移
func (p *Api) AutoMigrate() {
	err := p.db.Table(p.tableName).AutoMigrate(&ApiPostgresqlLog{})
	if err != nil {
		panic("创建表失败：" + err.Error())
	}
}

// Record 记录日志
func (p *Api) Record(content ApiPostgresqlLog) *gorm.DB {
	if utf8.ValidString(string(content.ResponseBody)) == false {
		log.Println("内容格式无法记录")
		return p.db
	}
	content.SystemHostName = p.hostname
	if content.SystemInsideIp == "" {
		content.SystemInsideIp = p.insideIp
	}
	content.GoVersion = p.goVersion
	resp := p.db.Table(p.tableName).Create(&content)
	if resp.RowsAffected == 0 {
		log.Println("Api：", resp.Error)
	}
	return resp
}

// Query 查询
func (p *Api) Query() *gorm.DB {
	return p.db.Table(p.tableName)
}

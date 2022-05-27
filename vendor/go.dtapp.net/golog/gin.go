package golog

import (
	"go.dtapp.net/goip"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"log"
	"os"
	"runtime"
	"strings"
)

type GinConfig struct {
	Db        *gorm.DB // 驱动
	TableName string   // 表名
}

// Gin 框架
type Gin struct {
	db        *gorm.DB // pgsql数据库
	tableName string   // 日志表名
	insideIp  string   // 内网ip
	hostname  string   // 主机名
	goVersion string   // go版本
}

// NewGin 创建框架实例化
func NewGin(config *GinConfig) *Gin {
	app := &Gin{}
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

// GinPostgresqlLog 结构体
type GinPostgresqlLog struct {
	LogId             uint           `gorm:"primaryKey" json:"log_id"`              //【记录】编号
	TraceId           string         `gorm:"type:text" json:"trace_id"`             //【系统】链编号
	RequestTime       TimeString     `gorm:"index" json:"request_time"`             //【请求】时间
	RequestUri        string         `gorm:"type:text" json:"request_uri"`          //【请求】请求链接 域名+路径+参数
	RequestUrl        string         `gorm:"type:text" json:"request_url"`          //【请求】请求链接 域名+路径
	RequestApi        string         `gorm:"type:text;index" json:"request_api"`    //【请求】请求接口 路径
	RequestMethod     string         `gorm:"type:text;index" json:"request_method"` //【请求】请求方式
	RequestProto      string         `gorm:"type:text" json:"request_proto"`        //【请求】请求协议
	RequestUa         string         `gorm:"type:text" json:"request_ua"`           //【请求】请求UA
	RequestReferer    string         `gorm:"type:text" json:"request_referer"`      //【请求】请求referer
	RequestBody       datatypes.JSON `gorm:"type:jsonb" json:"request_body"`        //【请求】请求主体
	RequestUrlQuery   datatypes.JSON `gorm:"type:jsonb" json:"request_url_query"`   //【请求】请求URL参数
	RequestIp         string         `gorm:"type:text" json:"request_ip"`           //【请求】请求客户端Ip
	RequestIpCountry  string         `gorm:"type:text" json:"request_ip_country"`   //【请求】请求客户端城市
	RequestIpRegion   string         `gorm:"type:text" json:"request_ip_region"`    //【请求】请求客户端区域
	RequestIpProvince string         `gorm:"type:text" json:"request_ip_province"`  //【请求】请求客户端省份
	RequestIpCity     string         `gorm:"type:text" json:"request_ip_city"`      //【请求】请求客户端城市
	RequestIpIsp      string         `gorm:"type:text" json:"request_ip_isp"`       //【请求】请求客户端运营商
	RequestHeader     datatypes.JSON `gorm:"type:jsonb" json:"request_header"`      //【请求】请求头
	ResponseTime      TimeString     `gorm:"index" json:"response_time"`            //【返回】时间
	ResponseCode      int            `gorm:"type:bigint" json:"response_code"`      //【返回】状态码
	ResponseMsg       string         `gorm:"type:text" json:"response_msg"`         //【返回】描述
	ResponseData      datatypes.JSON `gorm:"type:jsonb" json:"response_data"`       //【返回】数据
	CostTime          int64          `gorm:"type:bigint" json:"cost_time"`          //【系统】花费时间
	SystemHostName    string         `gorm:"type:text" json:"system_host_name"`     //【系统】主机名
	SystemInsideIp    string         `gorm:"type:text" json:"system_inside_ip"`     //【系统】内网ip
	GoVersion         string         `gorm:"type:text" json:"go_version"`           //【程序】Go版本
}

// AutoMigrate 自动迁移
func (p *Gin) AutoMigrate() {
	err := p.db.Table(p.tableName).AutoMigrate(&GinPostgresqlLog{})
	if err != nil {
		panic("创建表失败：" + err.Error())
	}
}

// Record 记录日志
func (p *Gin) Record(content GinPostgresqlLog) *gorm.DB {
	content.SystemHostName = p.hostname
	if content.SystemInsideIp == "" {
		content.SystemInsideIp = p.insideIp
	}
	content.GoVersion = p.goVersion
	resp := p.db.Table(p.tableName).Create(&content)
	if resp.RowsAffected == 0 {
		log.Println("Gin：", resp.Error)
	}
	return resp
}

// Query 查询
func (p *Gin) Query() *gorm.DB {
	return p.db.Table(p.tableName)
}

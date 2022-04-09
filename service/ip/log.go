package ip

import (
	"github.com/dtapps/go-library/service/ip/ip2region"
	v4 "github.com/dtapps/go-library/service/ip/v4"
	v6 "github.com/dtapps/go-library/service/ip/v6"
)

func (app *App) InitLog() {
	if app.Pgsql.Db == nil {
		return
	}

	err := app.Pgsql.Db.AutoMigrate(&postgresqlIpv4{}, &postgresqlIpv6{})
	if err != nil {
		panic(err.Error())
	}
}

// Ipv4 数据库
type postgresqlIpv4 struct {
	Id       int64
	Ip       string `gorm:"type:cidr"` // 输入的ip地址
	Country  string `gorm:"type:text"` // 国家
	Province string `gorm:"type:text"` // 省份
	City     string `gorm:"type:text"` // 城市
	Area     string `gorm:"type:text"` // 区域
	Isp      string `gorm:"type:text"` // 运营商
	Idc      string `gorm:"type:text"` // 运营商
}

func (m *postgresqlIpv4) TableName() string {
	return "ip_v4"
}

func (app *App) postgresqlIpv4Log(res v4.Result, resInfo ip2region.IpInfo) {
	if app.Pgsql.Db == nil {
		return
	}
	var query postgresqlIpv4
	app.Pgsql.Db.Where("ip = ?", resInfo.IP).Select("id").Take(&query)
	if query.Id == 0 {
		app.Pgsql.Db.Create(&postgresqlIpv4{
			Ip:       resInfo.IP,
			Country:  resInfo.Country,
			Province: resInfo.Province,
			City:     resInfo.City,
			Isp:      resInfo.ISP,
			Idc:      res.Area,
		})
	} else {
		app.Pgsql.Db.Model(&postgresqlIpv4{}).
			Where("ip = ?", res.IP).
			Select("country", "province", "city", "isp", "idc").
			Updates(postgresqlIpv4{
				Country:  resInfo.Country,
				Province: resInfo.Province,
				City:     resInfo.City,
				Isp:      resInfo.ISP,
				Idc:      res.Area,
			})
	}
}

// Ipv6 数据库
type postgresqlIpv6 struct {
	Id       int64
	Ip       string `gorm:"type:cidr"` // 输入的ip地址
	Country  string `gorm:"type:text"` // 国家
	Province string `gorm:"type:text"` // 省份
	City     string `gorm:"type:text"` // 城市
	Area     string `gorm:"type:text"` // 区域
	Isp      string `gorm:"type:text"` // 运营商
}

func (m *postgresqlIpv6) TableName() string {
	return "ip_v6"
}

func (app *App) postgresqlIpv6Log(res v6.Result) {
	if app.Pgsql.Db == nil {
		return
	}
	var query postgresqlIpv6
	app.Pgsql.Db.Where("ip = ?", res.IP).Select("id").Take(&query)
	if query.Id == 0 {
		app.Pgsql.Db.Create(&postgresqlIpv6{
			Ip:       res.IP,
			Country:  res.Country,
			Province: res.Province,
			City:     res.City,
			Area:     res.Area,
			Isp:      res.Isp,
		})
	} else {
		app.Pgsql.Db.Model(&postgresqlIpv6{}).
			Where("ip = ?", res.IP).
			Select("country", "area").
			Select("country", "province", "city", "area", "isp").
			Updates(postgresqlIpv6{
				Country:  res.Country,
				Province: res.Province,
				City:     res.City,
				Area:     res.Area,
				Isp:      res.Isp,
			})
	}
}

package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB struct {
	*gorm.DB
	O map[string]*gorm.DB
}

func ConnectDB() {
	DB.O = make(map[string]*gorm.DB)
	mc := CONF.GetStringMap("mysql")
	mcs := make(map[string]map[string]any)
Loop:
	for k, v := range mc {
		switch v.(type) {
		case string, int64, float64:
			mcs["default"] = mc
			break Loop
		case map[string]any:
			mcs[k] = v.(map[string]any)
		default:
			panic("mysql配置解析失败")
		}
	}
	for k, v := range mcs {
		dsn := fmt.Sprintf(
			"%v:%v@tcp(%s:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			v["user"], v["password"], v["host"], v["port"], v["db"],
		)
		fmt.Println(dsn)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("数据库连接失败: " + err.Error())
		}
		DB.O[k] = db
		DB.DB = DB.O["default"]
	}
}

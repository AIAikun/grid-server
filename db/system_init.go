package db

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server/model"
)

var DB *gorm.DB

func InitMysql() {
	var Config model.Config
	vip := viper.New()
	vip.SetConfigFile("./config/config.yaml")
	err := vip.ReadInConfig()
	if err != nil {
		fmt.Printf("viper Read Config failed, err:%v\n", err)
		panic(err)
		return
	}
	vip.WatchConfig()
	vip.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	if err := vip.Unmarshal(&Config); err != nil {
		fmt.Printf("viper Unmarshal failed, err:%v\n", err)
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=10s&parseTime=true&loc=Local", Config.Mysql.Username, Config.Mysql.Password, Config.Mysql.Host, Config.Mysql.Port, Config.Mysql.Dbname)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Printf("gorm.Open failed, err:%v\n", err)
		panic(err)
	}
	DB = db
}

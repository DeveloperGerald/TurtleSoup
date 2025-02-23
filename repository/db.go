package repository

import (
	"fmt"
	"log"

	"github.com/DeveloperGerald/TurtleSoup/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 导入 MySQL 驱动
)

var db *gorm.DB

func Init() {
	var err error
	dbConfig := config.GetConfig().Database
	// 连接数据库，用户名:密码@tcp(主机:端口)/数据库名?charset=utf8&parseTime=True&loc=Local
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database))
	if err != nil {
		panic(fmt.Errorf("init database error: %v", err))
	}

	log.Println("database connected")
}

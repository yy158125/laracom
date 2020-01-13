package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

func OpenDB() *gorm.DB {
	// 从系统环境变量获取数据库信息
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	config := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, DBName,
	)
	log.Println(config)
	db, err := gorm.Open("mysql", config)

	if err != nil {
		log.Println(err)
		log.Printf("Database connection failed. Database name: %s", DBName)
	}

	return db
}

package mysql

import (
	"fmt"
	"tugas/drivers/mysql/orders"
	"tugas/drivers/mysql/products"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
}

func (config *ConfigDB) InitDB() *gorm.DB {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("error when connecting to the database")
	}

	fmt.Println("connected to the database")

	return db
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&orders.Order{}, &products.Product{})
}

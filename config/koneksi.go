package config

import (
	"user-api/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect ...
func Connect() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/user-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// db.Migrator().DropTable(domain.User{})
	db.AutoMigrate(domain.User{})

	return db
}

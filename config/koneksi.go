package config

import (
	"os"
	"user-api/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect ...
func Connect() *gorm.DB {
	godotenv.Load()
	dbName := os.Getenv("MYSQL_DB")
	dsn := "root:@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.Migrator().DropTable(domain.User{})
	db.AutoMigrate(domain.User{})

	user := domain.User{
		NamaLengkap: "admin",
		Username:    "admin",
		Password:    "admin",
		Foto:        "makanan.png",
	}

	db.Create(&user)

	return db
}

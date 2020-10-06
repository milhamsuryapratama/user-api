package config

import (
	"os"
	"user-api/domain"

	"crypto/md5"
	"encoding/hex"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect ...
func Connect() *gorm.DB {
	godotenv.Load()
	dbName := os.Getenv("MYSQL_DB")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dsn := dbUser + ":" + dbPassword + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.Migrator().DropTable(domain.User{})
	db.AutoMigrate(domain.User{})

	hashPassword := md5.New()
	hashPassword.Write([]byte("admin"))
	user := domain.User{
		NamaLengkap: "admin",
		Username:    "admin",
		Password:    hex.EncodeToString(hashPassword.Sum(nil)),
		Foto:        "makanan.png",
	}

	db.Create(&user)

	return db
}

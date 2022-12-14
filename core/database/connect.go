// @Title
// @Description
// @Author
// @Update
package database

import (
	"fmt"

	"github.com/chihabMe/jwt-auth/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("fail to connect with the database")
		panic(err)
	}
	fmt.Println(" database connected successfully")
	Instance = db
	Migrate()
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
}

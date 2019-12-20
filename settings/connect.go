package settings

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Database() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gobackbone?parseTime=true")
	if err != nil {
		fmt.Println("failed to connect database")
		panic(err)
	}
	return db
}

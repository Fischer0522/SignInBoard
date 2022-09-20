package models

import (
	"SignInBoard/setting"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var db *gorm.DB
func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}



}

type SignIn struct {
	Id int
	Name string
	Message string
	SignInTime time.Time

}
func (SignIn) TableName() string {
	return "sign_in_board"
}

type TrashData struct {
	Id int
	Name string
}
func(TrashData) TableName()string{
	return "trash_data"
}


func CloseDB() {
	defer db.Close()
}









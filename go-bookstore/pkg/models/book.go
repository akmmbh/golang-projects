package models
//contain all the logic how controller contact databases

import(
	"github.com/jinzhu/gorm"
	"github.com/akmmbh/go-bookstore/pkg/config"
)
var db *gorm.DB

 type Book struct{
	gorm.Model `gorm:""`
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`


 }
 func init(){
	config.Connect()
	db= config.GetDB()
	db.AutoMigrate(&Book{})
 }
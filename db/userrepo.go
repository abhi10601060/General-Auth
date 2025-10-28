package db

import (
	"log"
	"synapse/auth/config"
	"synapse/auth/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db_url  = config.GetEnv("AUTH_DB")
	auth_db *gorm.DB
)

func init() {
	auth_db = connectToAuthDB()

	if auth_db != nil {
		autoMigrateModels()
	}
}

func connectToAuthDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(db_url), &gorm.Config{})

	if err != nil {
		log.Println("error in db connection : ", err)
		return nil
	}
	log.Println("Auth Db connected Successfully")

	return db
}

func autoMigrateModels() {
	auth_db.AutoMigrate(&model.User{})
}

func AddUser(user *model.User) int {
	res := auth_db.Save(user)

	if res.Error != nil {
		log.Println("error in storing user : ", res.Error.Error())
		return -1
	}

	log.Println("user saved succesfully")
	return 1
}

func UserExist(user *model.User) bool {
	var res int64

	err := auth_db.Model(&model.User{}).Where("id = ?", user.Id).Count(&res).Error
	if err != nil {
		log.Println("err during user Exists : ", err.Error())
		return false
	}
	return res > 0
}

func IsValidPassword(incominUser *model.User) bool {
	var userFromDB model.User = model.User{Id: incominUser.Id}
	err := auth_db.First(&userFromDB).Error

	if err != nil {
		log.Println("err from isValidPassword : ", err.Error())
		return false
	}

	return incominUser.Password == userFromDB.Password
}

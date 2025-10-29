package model


type User struct{
	Id string `json:"id" binding:"required" gorm:"primarykey"`
	Password string `json:"password" binding:"required" gorm:"not null"`
	Salt string `gorm:"not null"`
}
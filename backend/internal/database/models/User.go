package models

type User struct {
	Base
	Email string `gorm:"unique"`
}

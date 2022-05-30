package do

import (
	"personal-website/app/database/mysql"
)

type User struct {
	BasicModel
	NickName string `gorm:"column:nickname;type:varchar(256);not null"`
	Role     int    `gorm:"column:role;type:int;not null"`
}

func (user User) TableName() string {
	return "user"
}

// InsertUser insert a user
func InsertUser(user *User) (err error) {
	err = mysql.Db.Create(&user).Error
	return err
}

// SelectUserById select the user by id
func SelectUserById(id int) (user User, err error) {
	err = mysql.Db.First(&user, id).Error
	return user, err
}

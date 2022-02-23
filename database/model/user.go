package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(500);not null"`
	Role     int    `gorm:"type:int;DEFAULT:2"`
}

// GetUserIdByName
// get the user id by a name
func GetUserIdByName(name string) (id uint, err error) {
	db := GetDb()

	err = db.Select("id").Where("username = ?", name).First(&id).Error
	return id, err
}

// GetUserIdAndNameByName
// get the user id and name by a name
func GetUserIdAndNameByName(name string) (id uint, userName string, err error) {
	var user User
	db := GetDb()
	err = db.Select("id, username").Where("username = ?", name).First(&user).Error
	return user.ID, user.Username, err
}

// CreateUser
// insert a user
func CreateUser(data *User) (err error) {
	db := GetDb()
	err = db.Create(&data).Error
	return err
}

// GetUserById
// get the user by id
func GetUserById(id int) (user User, err error) {
	db := GetDb()
	err = db.Limit(1).Where("ID = ?", id).Find(&user).Error
	return user, err
}

// ListUsersInPage
// list users by the username in page
func ListUsersInPage(username string, pageSize int, pageNum int) (users []User, total int64, err error) {
	db := GetDb()

	if username != "" {
		// list the user
		err = db.Select("id,username,role,created_at").Where(
			"username LIKE ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
		if err != nil {
			return users, total, err
		}

		// get the total number
		err = db.Model(&users).Where(
			"username LIKE ?", username+"%",
		).Count(&total).Error
		return users, total, err
	}

	// list the user
	err = db.Select("id,username,role,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil {
		return users, total, err
	}

	// get the total number
	err = db.Model(&users).Count(&total).Error
	return users, total, err
}

// UpdateUser
// update a user
func UpdateUser(id int, user *User) (err error) {
	db := GetDb()

	var userToUpdate User
	userInfo := map[string]interface{}{
		"username": user.Username,
		"role":     user.Role,
	}

	err = db.Model(&userToUpdate).Where("id = ? ", id).Updates(userInfo).Error
	return err
}

// UpdateUserPassword
// update a user's password
func UpdateUserPassword(id int, user *User) (err error) {
	db := GetDb()

	err = db.Select("password").Where("id = ?", id).Updates(&user).Error
	return err
}

// DeleteUser
// delete a user
func DeleteUser(id int) (err error) {
	db := GetDb()

	var user User
	err = db.Where("id = ? ", id).Delete(&user).Error
	return err
}

//// BeforeCreate 密码加密&权限控制
//func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
//	u.Password = ScryptPw(u.Password)
//	u.Role = 2
//	return nil
//}
//
//func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
//	u.Password = ScryptPw(u.Password)
//	return nil
//}
//
//// ScryptPw 生成密码
//func ScryptPw(password string) string {
//	const cost = 10
//
//	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return string(HashPw)
//}
//
//// CheckLogin 后台登录验证
//func CheckLogin(username string, password string) (User, int) {
//	var user User
//	var PasswordErr error
//
//	db.Where("username = ?", username).First(&user)
//
//	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
//
//	if user.ID == 0 {
//		return user, error-msg.ErrorUserNotExist
//	}
//	if PasswordErr != nil {
//		return user, error-msg.ErrorPasswordWrong
//	}
//	if user.Role != 1 {
//		return user, error-msg.ErrorUserNoRight
//	}
//	return user, error-msg.Success
//}
//
//// CheckLoginFront 前台登录
//func CheckLoginFront(username string, password string) (User, int) {
//	var user User
//	var PasswordErr error
//
//	db.Where("username = ?", username).First(&user)
//
//	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
//	if user.ID == 0 {
//		return user, error-msg.ErrorUserNotExist
//	}
//	if PasswordErr != nil {
//		return user, error-msg.ErrorPasswordWrong
//	}
//	return user, error-msg.Success
//}

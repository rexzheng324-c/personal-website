package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"personal-website/utils/errmsg"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(500);not null"`
	Role     int    `gorm:"type:int;DEFAULT:2"`
}

// CheckUser 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.Success
}

// CheckUpUser 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.Success
	}
	if user.ID > 0 {
		return errmsg.ErrorUsernameUsed //1001
	}
	return errmsg.Success
}

// CreateUser 新增用户
func CreateUser(data *User) *errmsg.ApiError {
	//data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return &errmsg.ApiError{
			StatusCode: http.StatusInternalServerError,
			Code:       errmsg.SystemError,
			Message:    err.Error(),
		} // 500
	}
	return nil
}

// GetUser 查询用户
func GetUser(id int) (User, int) {
	var user User
	err := db.Limit(1).Where("ID = ?", id).Find(&user).Error
	if err != nil {
		return user, errmsg.Error
	}
	return user, errmsg.Success
}

// GetUsers 查询用户列表
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64

	if username != "" {
		db.Select("id,username,role,created_at").Where(
			"username LIKE ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Where(
			"username LIKE ?", username+"%",
		).Count(&total)
		return users, total
	}
	db.Select("id,username,role,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	db.Model(&users).Count(&total)

	return users, total
}

// EditUser 编辑用户信息
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&user).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

// ChangePassword 修改密码
func ChangePassword(id int, data *User) int {
	//var user User
	//var maps = make(map[string]interface{})
	//maps["password"] = data.Password

	err := db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

// BeforeCreate 密码加密&权限控制
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	u.Role = 2
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return nil
}

// ScryptPw 生成密码
func ScryptPw(password string) string {
	const cost = 10

	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}

	return string(HashPw)
}

// CheckLogin 后台登录验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username = ?", username).First(&user)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if user.ID == 0 {
		return user, errmsg.ErrorUserNotExist
	}
	if PasswordErr != nil {
		return user, errmsg.ErrorPasswordWrong
	}
	if user.Role != 1 {
		return user, errmsg.ErrorUserNoRight
	}
	return user, errmsg.Success
}

// CheckLoginFront 前台登录
func CheckLoginFront(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username = ?", username).First(&user)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 {
		return user, errmsg.ErrorUserNotExist
	}
	if PasswordErr != nil {
		return user, errmsg.ErrorPasswordWrong
	}
	return user, errmsg.Success
}

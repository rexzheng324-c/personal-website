package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"personal-website/database/model"
	"personal-website/database/ro"
	"personal-website/utils/errmsg"
	"personal-website/utils/validator"
	"strconv"
)

// CreateUser 添加用户
func CreateUser(c *gin.Context) {
	var data ro.CreateUserBody
	_ = c.ShouldBindJSON(&data)

	err := validator.Validate(&data)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	err = model.CreateUser(&model.User{
		Model:    gorm.Model{ID: 1},
		Username: data.Username,
		Password: data.Password,
		Role:     2,
	})
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
	return
}

// GetUser 查询单个用户
func GetUser(c *gin.Context) {
	c.JSON(
		http.StatusOK, gin.H{
			"status":  500,
			"message": "not yet",
		},
	)
	return
}

// ListUsers 查询用户列表
func ListUsers(c *gin.Context) {
	c.JSON(
		http.StatusOK, gin.H{
			"status":  500,
			"message": "not yet",
		},
	)
	return
}

// UpdateUser 编辑用户
func UpdateUser(c *gin.Context) {
	c.JSON(
		http.StatusOK, gin.H{
			"status":  500,
			"message": "not yet",
		},
	)
	return
}

// ChangeUserPassword 修改密码
func ChangeUserPassword(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.ChangePassword(id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteUser(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

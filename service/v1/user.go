package v1

import (
	"github.com/gin-gonic/gin"
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

	statusCode, msg := validator.Validate(&data)
	if statusCode != http.StatusOK {
		c.JSON(
			statusCode, gin.H{
				"message": msg,
			},
		)
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"message": "not yet",
		},
	)
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

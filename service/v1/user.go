package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal-website/database/model"
	"personal-website/utils/errmsg"
	"strconv"
)


// CreateUser 添加用户
func CreateUser(c *gin.Context) {
	return
}

// GetUser 查询单个用户
func GetUser(c *gin.Context) {
	return
}

// ListUsers 查询用户列表
func ListUsers(c *gin.Context) {
	return
}

// UpdateUser 编辑用户
func UpdateUser(c *gin.Context) {
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

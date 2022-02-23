package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"personal-website/database/model"
	"personal-website/service/result"
	"personal-website/service/ro"
	"personal-website/utils/error-msg"
	"personal-website/utils/validator"
	"strconv"
)

// CreateUser create a user
func CreateUser(c *gin.Context) {
	var data ro.CreateUserBody
	_ = c.ShouldBindJSON(&data)

	apiErr := validator.Validate(&data)
	if apiErr != nil {
		c.JSON(http.StatusOK, apiErr)
		return
	}

	err := model.CreateUser(&model.User{
		Model:    gorm.Model{},
		Username: data.Username,
		Password: data.Password,
		Role:     2,
	})
	if err != nil {
		c.JSON(http.StatusOK, error_msg.NewApiError(result.Fail, err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.NewSuccessBox(nil))
	return
}

// GetUser 查询单个用户
func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, error_msg.NewApiError(result.Fail, "ID should be int!"))
		return
	}

	user, err := model.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusOK, error_msg.NewApiError(result.Fail, err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.NewSuccessBox(user))
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

//// ChangeUserPassword 修改密码
//func ChangeUserPassword(c *gin.Context) {
//	var data model.User
//	id, _ := strconv.Atoi(c.Param("id"))
//	_ = c.ShouldBindJSON(&data)
//
//	code := model.ChangePassword(id, &data)
//
//	c.JSON(
//		http.StatusOK, gin.H{
//			"status":  code,
//			"message": error_msg.GetErrMsg(code),
//		},
//	)
//}
//
//// DeleteUser 删除用户
//func DeleteUser(c *gin.Context) {
//	id, _ := strconv.Atoi(c.Param("id"))
//
//	code := model.DeleteUser(id)
//
//	c.JSON(
//		http.StatusOK, gin.H{
//			"status":  code,
//			"message": error_msg.GetErrMsg(code),
//		},
//	)
//}

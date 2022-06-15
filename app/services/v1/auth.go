package v1

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"personal-website/app/middlewares"
	"personal-website/app/models/do"
	"personal-website/app/models/mapper"
	"personal-website/app/models/ro"
	"personal-website/app/utils/result"
	"personal-website/app/utils/validator"
)

func RegisterUser(c *gin.Context) {
	var data ro.RegisterUserBody
	_ = c.ShouldBindJSON(&data)

	err := validator.Validate(&data)
	if err != nil {
		c.JSON(http.StatusOK, result.NewFailBox(result.ParamsError, err))
		return
	}
	userID := uuid.New().String()
	user := do.User{
		BasicModel: do.BasicModel{ID: userID},
		NickName:   userID,
		Role:       do.UserRoleNormal,
	}
	basicAuth := do.BasicAuth{
		BasicModel: do.BasicModel{
			ID: uuid.New().String(),
		},
		UserId:   user.ID,
		Username: data.Username,
		Password: data.Password,
	}
	_, err = do.SelectBasicAuthByUsername(basicAuth.Username)
	if err == nil {
		c.JSON(http.StatusOK, result.NewFailBox(result.UserNameExist, errors.New("username exists")))
		return
	}
	err = mapper.RegisterUser(&user, &basicAuth)
	if err != nil {
		c.JSON(http.StatusOK, result.NewFailBox(result.Fail, err))
		return
	}

	session := sessions.Default(c)
	// Save the username in the session
	session.Set(middlewares.UserKey, userID)
	if err = session.Save(); err != nil {
		c.JSON(http.StatusOK, result.NewFailBox(result.Fail, errors.New("failed to save session")))
		return
	}
	c.JSON(http.StatusOK, result.NewSuccessBox(nil))
}

func LoginUser(c *gin.Context) {
	var data ro.LoginUserBody
	_ = c.ShouldBindJSON(&data)

	err := validator.Validate(&data)
	if err != nil {
		c.JSON(http.StatusOK, result.NewFailBox(result.ParamsError, err))
		return
	}
	_, err = do.SelectBasicAuthByUsername(data.Username)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, result.NewFailBox(result.UserNameNotExist, errors.New("username does not exist")))
		return
	}
	basicAuth, err := do.SelectBasicAuthByUsernameAndPassword(data.Username, data.Password)
	if err != nil {
		c.JSON(http.StatusOK, result.NewFailBox(result.WrongPassword, errors.New("wrong password")))
		return
	}
	session := sessions.Default(c)
	session.Set(middlewares.UserKey, basicAuth.UserId)
	if err = session.Save(); err != nil {
		c.JSON(http.StatusOK, result.NewFailBox(result.Fail, errors.New("failed to save session")))
		return
	}
	c.JSON(http.StatusOK, result.NewSuccessBox(nil))
}

func LogoutUser(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(middlewares.UserKey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusOK, result.NewFailBox(result.Fail, errors.New("failed to save session")))
		return
	}
	c.JSON(http.StatusOK, result.NewSuccessBox(nil))
}

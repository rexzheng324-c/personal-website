package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"personal-website/app/middlewares"
	"personal-website/app/models/do"
	"personal-website/app/models/dto"
	"personal-website/app/models/ro"
	"personal-website/app/utils/result"
	"personal-website/app/utils/validator"
)

func CreateBlog(c *gin.Context) {
	var data ro.CreateBlogBody
	_ = c.ShouldBindJSON(&data)

	err := validator.Validate(&data)
	if err != nil {
		c.JSON(http.StatusOK, result.NewFailBox(result.ParamsError, err))
		return
	}
	userId := c.GetString(middlewares.UserIdKey)
	user, err := do.SelectUserById(userId)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, result.NewFailBox(result.RecordNotFound, errors.New("user does not exist")))
		return
	}
	if user.Role != do.UserRoleAdmin {
		c.JSON(http.StatusOK, result.NewFailBox(result.NotAdmin, errors.New("user is not admin")))
		return
	}
	blog := &do.Blog{
		BasicModel: do.BasicModel{ID: uuid.New().String()},
		UserId:     userId,
		Title:      data.Title,
		Content:    data.Content,
	}

	err = do.InsertBlog(blog)
	if err != nil {
		c.JSON(http.StatusOK, result.NewFailBox(result.Fail, err))
		return
	}

	c.JSON(http.StatusOK, result.NewSuccessBox(nil))
}

func ListBlogs(c *gin.Context) {
	var data ro.ListBlogsBody
	_ = c.ShouldBindJSON(&data)

	err := validator.Validate(&data)
	if err != nil {
		c.JSON(http.StatusOK, result.NewFailBox(result.ParamsError, err))
		return
	}

	blogCondition := &do.Blog{
		Title: data.Title,
	}
	pageCondition := do.PageCondition{
		Limit:  data.Limit,
		Offset: data.Offset,
	}

	blogs, err := do.SelectBlogsByCondition(*blogCondition, pageCondition)
	if err != nil {
		c.JSON(http.StatusOK, result.NewFailBox(result.Fail, err))
		return
	}

	blogDTOs := make([]dto.Blog, 0)
	for _, blog := range blogs {
		blogDTOs = append(blogDTOs, dto.Blog{
			Id:      blog.ID,
			Title:   blog.Title,
			Content: blog.Content,
		})
	}

	c.JSON(http.StatusOK, result.NewSuccessBox(dto.ListBlogsBody{
		Blogs:  blogDTOs,
		Offset: pageCondition.Offset,
		Count:  len(blogDTOs),
	}))
}

func DeleteBlog(c *gin.Context) {
	var data ro.DeleteBlogBody
	_ = c.ShouldBindJSON(&data)

	err := validator.Validate(&data)
	if err != nil {
		c.JSON(http.StatusOK, result.NewFailBox(result.ParamsError, err))
		return
	}
	userId := c.GetString(middlewares.UserIdKey)
	user, err := do.SelectUserById(userId)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, result.NewFailBox(result.RecordNotFound, errors.New("user does not exist")))
		return
	}
	if user.Role != do.UserRoleAdmin {
		c.JSON(http.StatusOK, result.NewFailBox(result.NotAdmin, errors.New("user is not admin")))
		return
	}

	err = do.DeleteBlog(data.Id)
	if err != nil {
		c.JSON(http.StatusOK, result.NewFailBox(result.Fail, err))
		return
	}

	c.JSON(http.StatusOK, result.NewSuccessBox(nil))
}

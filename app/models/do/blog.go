package do

import (
	"personal-website/app/databases/mysql"
)

type Blog struct {
	BasicModel
	UserId  string `gorm:"column:user_id;type:varchar(128);not null"`
	Title   string `gorm:"column:title;type:varchar(128);not null"`
	Content string `gorm:"column:content;type:text;not null"`
}

func (b Blog) TableName() string {
	return "blog"
}

func InsertBlog(blog *Blog) error {
	err := mysql.Db.Create(&blog).Error
	return err
}

func SelectBlogsByCondition(blogCondition Blog, pageCondition PageCondition) (blogs []Blog, err error) {
	db := mysql.Db
	if blogCondition.Title != "" {
		db = db.Where("title = ?", blogCondition.Title)
	}
	if blogCondition.UserId != "" {
		db = db.Where("user_id = ?", blogCondition.UserId)
	}
	err = db.Offset(pageCondition.Offset).Limit(pageCondition.Limit).Find(&blogs).Offset(-1).Limit(-1).Error
	return blogs, err
}

func DeleteBlog(id string) error {
	var blog Blog
	err := mysql.Db.Where("id =  ?", id).Delete(&blog).Error
	return err
}

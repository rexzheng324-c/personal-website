package mapper

import (
	"personal-website/app/databases/mysql"
	"personal-website/app/models/do"
)

func RegisterUser(user *do.User, basicAuth *do.BasicAuth) error {
	tx := mysql.Db.Begin()

	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Create(&basicAuth).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

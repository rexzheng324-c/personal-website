package do

//go:generate mockgen -source=auth.go -destination=./mocks/mock_auth.go -package=mocks

import "personal-website/app/databases/mysql"

type BasicAuth struct {
	BasicModel
	UserId   string `gorm:"column:user_id;type:varchar(128);not null"`
	Username string `gorm:"column:username;type:varchar(128);not null"`
	Password string `gorm:"column:password;type:varchar(500);not null"`
}

func (b BasicAuth) TableName() string {
	return "basic_auth"
}

func SelectBasicAuthByUsernameAndPassword(username string, password string) (BasicAuth, error) {
	var basicAuth BasicAuth
	err := mysql.Db.Where("username = ? and password = ?", username, password).First(&basicAuth).Error
	return basicAuth, err
}

func SelectBasicAuthByUsername(username string) (BasicAuth, error) {
	var basicAuth BasicAuth
	err := mysql.Db.Where("username = ?", username).First(&basicAuth).Error
	return basicAuth, err
}

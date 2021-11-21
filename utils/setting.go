package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("config/local_config.ini")
	if err != nil {
		fmt.Println("config loading fails, please check the path:", err)
	}
	LoadDB(file)
}

func LoadDB(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("rex_website")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("website")
}

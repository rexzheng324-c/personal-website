package mysql

import (
	"errors"
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var Db *gorm.DB

func GetDb(configPath string) (*gorm.DB, error) {
	file, err := ini.Load(configPath)
	if err != nil {
		return nil, errors.New("config loading fails")
	}
	DbHost := file.Section("mysql").Key("DbHost")
	DbPort := file.Section("mysql").Key("DbPort")
	DbUser := file.Section("mysql").Key("DbUser")
	DbPassWord := file.Section("mysql").Key("DbPassWord")
	DbName := file.Section("mysql").Key("DbName")
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUser,
		DbPassWord,
		DbHost,
		DbPort,
		DbName,
	)
	Db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	sqlDB, _ := Db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	return Db, nil
}

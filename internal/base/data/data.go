package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Data struct {
	DB *gorm.DB
}

func NewData(db *gorm.DB) *Data {
	return &Data{
		DB: db,
	}
}

func NewDB() (*gorm.DB, error) {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/ClassMoments?charset=utf8&parseTime=True&loc=Local"
	dsn := "root:123456@tcp(mysql:3306)/ClassMoments?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	schema.RegisterSerializer("json", schema.JSONSerializer{})
	//err = db.AutoMigrate(
	//	&entity.User{},
	//	&entity.Like{},
	//	&entity.Comment{},
	//	&entity.Moment{},
	//	&entity.Notification{},
	//	&entity.Class{},
	//	&entity.School{},
	//	&entity.Teacher{},
	//	&entity.ParentRelChild{},
	//	&entity.Media{},
	//)
	if err != nil {
		return nil, err
	}
	return db, nil
}

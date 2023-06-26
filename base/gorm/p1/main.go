package main

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func main() {
	dsn := "taishan:E7XIzjXtQpv7mL32@tcp(rm-2zee82n4243c8o406.mysql.rds.aliyuncs.com:3306)/permission?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("Failed to connect to database!")
	}
	//get(db, Title("国宝大熊猫"), Status(0))
	//get(db, Title("国宝"), Status(0))

	ge(db, 1, 10)

	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to close database connection!")
	}
	sqlDB.Close()
}

func get(db *gorm.DB, options ...func(option *gorm.DB)) {
	var articles []*ArticlePrompt
	//db = db.Table(ArticlePrompt{}.TableName())
	fmt.Println(db)
	fmt.Println("------------")
	for _, option := range options {
		option(db)
	}

	res := db.WithContext(context.Background()).Find(&articles)
	fmt.Println(res)
}

func ge(db *gorm.DB, pageId, pageSize int) {
	var total int64
	var res []User
	db = db.WithContext(context.Background()).Model(&User{})
	db.Order("id DESC").Limit(pageSize).Offset(pageSize * (pageId - 1)).Find(&res)

	if db.Error != nil {
		fmt.Println(db.Error)
	}

	db.Count(&total)

	fmt.Println(res)

}

type Option func(db *gorm.DB)

func Title(title string) Option {
	return func(db *gorm.DB) {
		db.Where("title = ?", title)
	}
}

func Status(status int) Option {
	return func(db *gorm.DB) {
		db.Where("status = ?", status)
	}
}

func (User) TableName() string {
	return "user"
}

type User struct {
	ID         int64     `gorm:"id"`
	WorkCode   string    `gorm:"work_code"`
	Status     int       `gorm:"status"`
	CreateTime time.Time `gorm:"create_time"`
	UpdateTime time.Time `gorm:"update_time"`
}

type ArticlePrompt struct {
	ID        int64
	Grade     int
	Title     string
	Outline   string
	Require   string
	Example   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    int
}

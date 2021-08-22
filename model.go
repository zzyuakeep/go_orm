package main

import (
	"database/sql"
	"fmt"
	"time"
)

type UserBack struct {
	ID           uint   `gorm:"primaryKey:true autoIncrement:true"`
	Name         string `gorm:"index:name"`
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type BaseModel struct {
	UUID uint `gorm:"primaryKey"`
	Time time.Time
}

type User struct {
	//gorm.Model          // 这个里面包含了  ID, CreatedAt, UpdatedAt ,DeletedAt
	BaseModel BaseModel `gorm:"embedded ;embeddedPrefix:mm_"` // embedded	嵌套字段, 如果没有指定这个,gorm 认为这个是联表关系的, 需要创建外键关联
	Name      string    `gorm:"index; comment:姓名 ; default:张三 "`
	//ID           uint   `gorm:"primaryKey:true autoIncrement:true"`
	Email        *string `gorm:"not null; comment:邮箱; column:email"`
	Age          uint8   `gorm:"comment:年龄"`
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
}


func CreateUser() {
	err := GLOBAL_DB.AutoMigrate(&User{})
	if err != nil {

		fmt.Println("创建user表失败,err:", err)
		return
	}
}

func DROpUser() {
	err := GLOBAL_DB.Migrator().DropTable(&User{})
	if err != nil {
		fmt.Println(err)
		return
	}
}

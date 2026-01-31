package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	Name    string `gorm:"type:varchar(50);not null;comment:'姓名'"`
	Phone   string `gorm:"type:varchar(50);not null;comment:'电话号'"`
	Sex     string `gorm:"type:enum('男','女');not null;comment:'性别'"`
	Address string `gorm:"type:varchar(50);not null;comment:'家庭住址'"`
	History string `gorm:"type:varchar(50);not null;comment:'病史'"`
	SeralNo string `gorm:"type:varchar(50);not null;comment:'病例号'"`
}

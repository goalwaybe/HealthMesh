package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);not null;comment:'科室名称'"`
}

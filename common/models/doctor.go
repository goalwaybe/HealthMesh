package models

import (
	"gorm.io/gorm"
	"medi-biz/common/pkg"
	pb "medi-biz/proto/medi"
)

type Doctor struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(50);not null;comment:'医生姓名'"`
	Title        string  `gorm:"type:varchar(50);not null;comment:'职位'"`
	Skill        string  `gorm:"type:varchar(50);not null;comment:'擅长领域'"`
	DepartmentID int     `gorm:"type:int;not null;comment:'科室id'"`
	ServicePrice float64 `gorm:"type:float;not null;comment:'诊费'"`
	Status       int     `gorm:"type:int;not null;comment:'状态 1-正常 2-禁用'"`
	Author       string  `gorm:"type:varchar(200);comment:'照片'"`
}

func (d *Doctor) FindDoctorAdd(db *gorm.DB) error {
	return db.Debug().Create(d).Error
}

func (d *Doctor) FindDoctorByID(db *gorm.DB, id int64) error {
	return db.Debug().Where("id=?", id).First(d).Error
}

func (d *Doctor) FindDoctorList(db *gorm.DB, in *pb.DoctorListReq) ([]Doctor, int64) {
	var list []Doctor
	var total int64
	db.Debug().Model(&Doctor{}).Count(&total).
		Scopes(pkg.Paginate(int(in.Page), int(in.Size))).
		Select("doctors.`name` AS `Name`",
			"doctors.title AS Title",
			"doctors.skill AS Skill",
			"doctors.service_price AS ServicePrice",
		).
		Joins("LEFT JOIN departments ON doctors.department_id = departments.id").
		Find(&list)

	return list, total

	/*
		SELECT
		  doctors.`name` AS `Name`,
		  doctors.title AS Title,
		  doctors.skill AS Skill,
		  doctors.service_price AS ServicePrice
		FROM
		  doctors
		  LEFT JOIN departments ON doctors.department_id = departments.id;
	*/
}

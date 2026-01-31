package models

import (
	"gorm.io/gorm"
	pb "medi-biz/proto/medi"
	"time"
)

type Schedule struct {
	gorm.Model
	PatientID    int       `gorm:"type:int;not null;comment:'患者id'"`
	DoctorID     int       `gorm:"type:int;not null;comment:'医生id'"`
	DepartmentID int       `gorm:"type:int;not null;comment:'科室id'"`
	ScheduleDate time.Time `gorm:"type:datetime;not null;comment:'预约日期'"`
	ScheduleTime string    `gorm:"type:varchar(50);not null;comment:'预约时间段 1-上午 2-下午 3-晚上'"`
	TotalNum     int       `gorm:"type:int;not null;comment:'总号源'"`
	UsedNum      int       `gorm:"type:int;comment:'已使用号源(+1)'"`
	Status       int       `gorm:"type:int;not null;comment:'预约状态 1-可约 2-不可约'"`
}

func (s *Schedule) FindScheduleAdd(db *gorm.DB) error {
	return db.Debug().Create(s).Error
}

func (s *Schedule) FindScheDuleByID(db *gorm.DB, in *pb.ScheduleAddReq) error {
	return db.Debug().Where("patient_id = ? and doctor_id = ? and schedule_date = ? and schedule_time = ?",
		in.PatientID, in.DoctorID, in.ScheduleDate, in.ScheduleTime).Find(s).Error
}

func (s *Schedule) FindScheduleByID(db *gorm.DB, id int64) error {
	return db.Debug().Where("id=?", id).First(s).Error
}

package models

import (
	"gorm.io/gorm"
	pb "medi-biz/proto/medi"
	"time"
)

type Appointment struct {
	gorm.Model
	PatientID     int        `gorm:"type:int;not null;comment:'患者id'"`
	DoctorID      int        `gorm:"type:int;not null;comment:'医生id'"`
	DepartmentID  int        `gorm:"type:int;not null;comment:'科室id'"`
	ScheduleID    int        `gorm:"type:int;not null;comment:'预约id'"`
	AppointmentNo string     `gorm:"type:varchar(50);not null;comment:'排班号'"`
	ScheduleDate  time.Time  `gorm:"type:datetime;not null;comment:'预约日期'"`
	ScheduleTime  string     `gorm:"type:varchar(50);not null;comment:'预约时间段 1-上午 2-下午 3-晚上'"`
	PayStatus     int        `gorm:"type:int;not null;comment:'支付状态 1-已支付 2-未支付 3-已取消'"`
	PayServe      int        `gorm:"type:int;not null;comment:'支付方式 1-微信 2-支付宝 3-银行卡'"`
	PayTime       *time.Time `gorm:"type:datetime;comment:'支付时间'"`
	SeralNo       string     `gorm:"type:varchar(50);not null;comment:'就诊号'"`
}

func (a *Appointment) FindData(db *gorm.DB, in *pb.AppointmentAddReq) error {
	return db.Debug().Create(a).Error
}

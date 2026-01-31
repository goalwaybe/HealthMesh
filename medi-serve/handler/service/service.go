package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"math/rand"
	"medi-biz/common/config"
	"medi-biz/common/models"
	"medi-biz/common/pkg"
	pb "medi-biz/proto/medi" //注意这个路径
	"strconv"
	"time"
)

// 定义空接口
type Service struct {
	pb.UnimplementedMediServer
}

// 实现方法
func (s *Service) DoctorAdd(ctx context.Context, in *pb.DoctorAddReq) (*pb.DoctorAddResp, error) {
	data := models.Doctor{
		Name:         in.Name,
		Title:        in.Title,
		Skill:        in.Skill,
		DepartmentID: int(in.DepartmentID),
		ServicePrice: float64(in.ServicePrice),
		Status:       1,
	}

	err := data.FindDoctorAdd(config.DB)
	if err != nil {
		return nil, errors.New("添加失败")
	}

	return &pb.DoctorAddResp{
		Id: int64(data.ID),
	}, nil
}

func (s *Service) ScheduleAdd(ctx context.Context, in *pb.ScheduleAddReq) (*pb.ScheduleAddResp, error) {
	paresScheduleDate, err := time.Parse("2006-01-02", in.ScheduleDate)
	if err != nil {
		return nil, errors.New("日期转换失败")
	}
	var doctor models.Doctor
	err = doctor.FindDoctorByID(config.DB, in.DoctorID)
	if err != nil {
		return nil, errors.New("医生不存在")
	}
	var schedule models.Schedule
	err = schedule.FindScheDuleByID(config.DB, in)
	if err != nil {
		return nil, errors.New("一个医生在同一个时间不能重复预约")
	}
	data := models.Schedule{
		PatientID:    int(in.PatientID),
		DoctorID:     int(in.DoctorID),
		DepartmentID: int(in.DepartmentID),
		ScheduleDate: paresScheduleDate,
		ScheduleTime: in.ScheduleTime,
		TotalNum:     1,
		UsedNum:      0,
		Status:       1,
	}
	err = data.FindScheduleAdd(config.DB)
	if err != nil {
		return nil, errors.New("预约失败")
	}
	return &pb.ScheduleAddResp{
		Id: int64(data.ID),
	}, nil
}

func (s *Service) DoctorList(ctx context.Context, in *pb.DoctorListReq) (*pb.DoctorListResp, error) {
	var DoctorModel models.Doctor
	var DoctorList []models.Doctor
	DoctorList, total := DoctorModel.FindDoctorList(config.DB, in)

	var list []*pb.DoctorListData
	for _, v := range DoctorList {
		lis := &pb.DoctorListData{
			Name:         v.Name,
			Title:        v.Title,
			Skill:        v.Skill,
			DepartmentID: int64(v.DepartmentID),
			ServicePrice: float32(v.ServicePrice),
		}
		list = append(list, lis)
	}
	return &pb.DoctorListResp{
		List:  list,
		Total: total,
	}, nil
}

func (s *Service) AppointmentAdd(ctx context.Context, in *pb.AppointmentAddReq) (*pb.AppointmentAddResp, error) {
	var appointment models.Schedule
	err := appointment.FindScheduleByID(config.DB, in.ScheduleID)
	if err != nil {
		return nil, errors.New("预约不存在")
	}

	var Doctor models.Doctor
	err = Doctor.FindDoctorByID(config.DB, in.DoctorID)
	if err != nil {
		return nil, errors.New("医生不存在")
	}

	if appointment.UsedNum >= appointment.TotalNum {
		return nil, errors.New("不可预约")
	}
	appointmentNo := rand.Intn(9000) + 1000

	total := pkg.GenWareTotal()

	err = config.DB.Transaction(func(tx *gorm.DB) error {
		paresScheduleDate, _ := time.Parse("2006-01-02", in.ScheduleDate)
		data := models.Appointment{
			PatientID:    int(in.PatientID),
			DoctorID:     int(in.DoctorID),
			DepartmentID: int(in.DepartmentID),
			ScheduleID:   int(in.ScheduleID),
			ScheduleDate: paresScheduleDate,
			ScheduleTime: in.ScheduleTime,
			PayStatus:    1,
			PayServe:     2,
			PayTime:      nil,
		}
		err = data.FindData(config.DB, in)
		if err != nil {
			return err
		}
		return nil
	})

	payUrl := pkg.AliPagePay(total, float64(appointmentNo))

	return &pb.AppointmentAddResp{
		AppointmentNo: strconv.Itoa(appointmentNo),
		SeralNo:       total,
		PayUrl:        payUrl,
	}, nil
}

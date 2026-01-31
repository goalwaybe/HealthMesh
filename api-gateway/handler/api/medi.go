package api

import (
	"github.com/gin-gonic/gin"
	"medi-biz/api-gateway/handler/request"
	"medi-biz/common/config"
	pb "medi-biz/proto/medi"
)

func DoctorAdd(c *gin.Context) {
	var form request.DoctorAddRequest
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "请求参数缺失",
			"data": nil,
		})
		return
	}
	r, err := config.MediClient.DoctorAdd(c, &pb.DoctorAddReq{
		Name:         form.Name,
		Title:        form.Title,
		Skill:        form.Skill,
		DepartmentID: int64(form.DepartmentID),
		ServicePrice: float32(form.ServicePrice),
	})
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": r,
		})
		return
	}
}

func ScheduleAdd(c *gin.Context) {
	var form request.ScheduleAddRequest
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "请求参数缺失",
			"data": nil,
		})
		return
	}
	r, err := config.MediClient.ScheduleAdd(c, &pb.ScheduleAddReq{
		PatientID:    int64(form.PatientID),
		DoctorID:     int64(form.DoctorID),
		DepartmentID: int64(form.DepartmentID),
		ScheduleDate: form.ScheduleDate,
		ScheduleTime: form.ScheduleTime,
	})
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": r,
		})
		return
	}
}

func DoctorList(c *gin.Context) {
	var form request.DoctorListRequest
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "请求参数缺失",
			"data": nil,
		})
		return
	}
	r, err := config.MediClient.DoctorList(c, &pb.DoctorListReq{
		Page: int64(form.Page),
		Size: int64(form.Size),
		Name: form.Name,
	})
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": r,
		})
		return
	}
}

func AppointmentAdd(c *gin.Context) {
	var form request.AppointmentAddRequest
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "请求参数缺失",
			"data": nil,
		})
		return
	}
	r, err := config.MediClient.AppointmentAdd(c, &pb.AppointmentAddReq{
		PatientID:    int64(form.PatientID),
		DoctorID:     int64(form.DoctorID),
		DepartmentID: int64(form.DepartmentID),
		ScheduleID:   int64(form.ScheduleID),
		ScheduleDate: form.ScheduleDate,
		ScheduleTime: form.ScheduleTime,
	})
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": r,
		})
		return
	}
}

func Notify(c *gin.Context) {

}

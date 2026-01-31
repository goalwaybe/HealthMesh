package request

type DoctorAddRequest struct {
	Name         string  `form:"name" json:"name" xml:"name"  binding:"required"`
	Title        string  `form:"title" json:"title" xml:"title"  binding:"required"`
	Skill        string  `form:"skill" json:"skill" xml:"skill"  binding:"required"`
	DepartmentID int     `form:"department_id" json:"department_id" xml:"department_id"  binding:"required"`
	ServicePrice float64 `form:"service_price" json:"service_price" xml:"service_price"  binding:"required"`
}

type ScheduleAddRequest struct {
	PatientID    int    `form:"patient_id" json:"patient_id" xml:"patient_id"  binding:"required"`
	DoctorID     int    `form:"doctor_id" json:"doctor_id" xml:"doctor_id"  binding:"required"`
	DepartmentID int    `form:"department_id" json:"department_id" xml:"department_id"  binding:"required"`
	ScheduleDate string `form:"schedule_date" json:"schedule_date" xml:"schedule_date"  binding:"required"`
	ScheduleTime string `form:"schedule_time" json:"schedule_time" xml:"schedule_time"  binding:"required"`
}

type DoctorListRequest struct {
	Page int    `form:"page" json:"page" xml:"page"  binding:"required"`
	Size int    `form:"size" json:"size" xml:"size"  binding:"required"`
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
}

type AppointmentAddRequest struct {
	PatientID    int    `form:"patient_id" json:"patient_id" xml:"patient_id"  binding:"required"`
	DoctorID     int    `form:"doctor_id" json:"doctor_id" xml:"doctor_id"  binding:"required"`
	DepartmentID int    `form:"department_id" json:"department_id" xml:"department_id"  binding:"required"`
	ScheduleID   int    `form:"schedule_id" json:"schedule_id" xml:"schedule_id"  binding:"required"`
	ScheduleDate string `form:"schedule_date" json:"schedule_date" xml:"schedule_date"  binding:"required"`
	ScheduleTime string `form:"schedule_time" json:"schedule_time" xml:"schedule_time"  binding:"required"`
}

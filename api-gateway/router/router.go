package router

import (
	"github.com/gin-gonic/gin"
	"medi-biz/api-gateway/handler/api"
	"net/http"
)

func Router(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	r.POST("/doctorAdd", api.DoctorAdd)
	r.POST("/scheduleAdd", api.ScheduleAdd)
	r.POST("/doctorList", api.DoctorList)
	r.POST("/appointmentAdd", api.AppointmentAdd)
	r.POST("/notify", api.Notify)
}

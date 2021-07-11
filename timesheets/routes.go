package timesheets

import (
	"github.com/gin-gonic/gin"
)

func SetupTimesheetsRoutes(router *gin.Engine) {
	v1 := router.Group("v1") 

	go v1.GET("", GetTimesheets)
}
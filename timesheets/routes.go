package timesheets

import (
	"github.com/gin-gonic/gin"
)

func SetupTimesheetsRoutes(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		timesheetRouteGroup := v1.Group("/timesheets")
		{
			go timesheetRouteGroup.GET("", GetTimesheets)
			go timesheetRouteGroup.POST("", AddTimesheets)

			go timesheetRouteGroup.GET("/:id", GetTimesheet)
		}
	}
}
package timesheets

import (
	"github.com/gin-gonic/gin"
)

func SetupTimesheetsRoutes(router *gin.RouterGroup) {
	timesheetRouteGroup := router.Group("/timesheets")
	{
		go timesheetRouteGroup.GET("", GetTimesheets)
		go timesheetRouteGroup.POST("", AddTimesheets)
		go timesheetRouteGroup.DELETE("", RemoveTimesheets)

		go timesheetRouteGroup.GET("/:id", GetTimesheet)
		go timesheetRouteGroup.PUT("/:id", UpdateTimesheets)
	}
}

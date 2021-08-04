package timesheets

import (
	"github.com/gin-gonic/gin"
)

func SetupTimesheetsRoutes(router *gin.RouterGroup) {
	timesheetRouteGroup := router.Group("/timesheets")
	{
		timesheetRouteGroup.GET("", getTimesheets)
		timesheetRouteGroup.POST("", addTimesheets)
		timesheetRouteGroup.DELETE("", removeTimesheets)

		timesheetRouteGroup.GET("/:id", getTimesheet)
		timesheetRouteGroup.PUT("/:id", updateTimesheets)
	}
}

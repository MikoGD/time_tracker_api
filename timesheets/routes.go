package timesheets

import (
	"github.com/gin-gonic/gin"
)

func SetupTimesheetsRoutes(router *gin.RouterGroup) {
	timesheetRouteGroup := router.Group("/timesheets")
	{
		go timesheetRouteGroup.GET("", getTimesheets)
		go timesheetRouteGroup.POST("", addTimesheets)
		go timesheetRouteGroup.DELETE("", removeTimesheets)

		go timesheetRouteGroup.GET("/:id", getTimesheet)
		go timesheetRouteGroup.PUT("/:id", updateTimesheets)
	}
}

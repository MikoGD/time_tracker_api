package rows

import (
	"github.com/gin-gonic/gin"
)

func SetupTimesheetRowsRoutes(router *gin.RouterGroup) {
	timesheetRowsRouteGroup := router.Group("/rows")
	{
		go timesheetRowsRouteGroup.GET("", GetRows)

		go timesheetRowsRouteGroup.GET("/:id", GetRow)

		go timesheetRowsRouteGroup.GET("/timesheet/:id", GetRowByTimesheet)
	}
}

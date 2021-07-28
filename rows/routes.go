package rows

import (
	"github.com/gin-gonic/gin"
)

func SetupTimesheetRowsRoutes(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		timesheetRowsRouteGroup := v1.Group("/rows")
		{
			go timesheetRowsRouteGroup.GET("", GetRows)

			go timesheetRowsRouteGroup.GET("/:id", GetRow)

			go timesheetRowsRouteGroup.GET("/timesheet/:id", GetRowByTimesheet)
		}
	}
}

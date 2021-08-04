package rows

import (
	"github.com/gin-gonic/gin"
)

func SetupTimesheetRowsRoutes(router *gin.RouterGroup) {
	rowsRouteGroup := router.Group("/rows")
	{
		rowsRouteGroup.GET("", getRows)
		rowsRouteGroup.DELETE("", removeRowByIds)

		rowsRouteGroup.GET("/:id", getRow)

		timesheetRouteGroup := rowsRouteGroup.Group("/timesheet")
		{
			timesheetRouteGroup.GET("/:id", getRowByTimesheet)
			timesheetRouteGroup.POST("/:id", addRowToTimesheet)
			timesheetRouteGroup.PUT("/:id", updateRowInTimesheet)
		}
	}
}

package rows

import (
	"github.com/gin-gonic/gin"
)

func SetupTimesheetRowsRoutes(router *gin.RouterGroup) {
	rowsRouteGroup := router.Group("/rows")
	{
		go rowsRouteGroup.GET("", getRows)
		go rowsRouteGroup.DELETE("", removeRowByIds)

		go rowsRouteGroup.GET("/:id", getRow)

		timesheetRouteGroup := rowsRouteGroup.Group("/timesheet")
		{
			go timesheetRouteGroup.GET("/:id", getRowByTimesheet)
			go timesheetRouteGroup.POST("/:id", addRowToTimesheet)
			go timesheetRouteGroup.PUT("/:id", updateRowInTimesheet)
		}
	}
}

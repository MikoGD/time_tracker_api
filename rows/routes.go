package rows

import (
	"github.com/gin-gonic/gin"
)

func SetupTimesheetRowsRoutes(router *gin.RouterGroup) {
	rowsRouteGroup := router.Group("/rows")
	{
		go rowsRouteGroup.GET("", GetRows)
		go rowsRouteGroup.DELETE("", RemoveRowByIds)

		go rowsRouteGroup.GET("/:id", GetRow)

		timesheetRouteGroup := rowsRouteGroup.Group("/timesheet")
		{
			go timesheetRouteGroup.GET("/:id", GetRowByTimesheet)
			go timesheetRouteGroup.POST("/:id", AddRowToTimesheet)
		}
	}
}

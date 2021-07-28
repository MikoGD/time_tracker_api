package rows

import (
	"fmt"

	"example.com/time_tracker_api/utils"
	"github.com/gin-gonic/gin"
)

func GetRows(context *gin.Context) {
	rows, err := utils.SelectFromTable(tableName, "*", "")

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	timesheetRows, err := parseRows(rows)

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	sendQuerySuccessResponse(context, timesheetRows)
}

func GetRow(context *gin.Context) {
	id := context.Param("id")

	condition := fmt.Sprintf("row_id=%s", id)

	rows, err := utils.SelectFromTable(tableName, "*", condition)

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	timesheetRows, err := parseRows(rows)

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	sendQuerySuccessResponse(context, timesheetRows)
}

func GetRowByTimesheet(context *gin.Context) {
	id := context.Param("id")

	condition := fmt.Sprintf("timesheet_id=%s", id)

	rows, err := utils.SelectFromTable(tableName, "*", condition)

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	timesheetRows, err := parseRows(rows)

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	sendQuerySuccessResponse(context, timesheetRows)
}

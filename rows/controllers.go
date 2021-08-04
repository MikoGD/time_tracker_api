package rows

import (
	"fmt"

	"example.com/time_tracker_api/utils"
	"github.com/gin-gonic/gin"
)

func addRowToTimesheet(context *gin.Context) {
	values, err := parseRequestBodyForInsertValues(context)

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	rowsAffected, err := utils.InsertToTable(tableName, columns, values)

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	sendExecSuccessResponse(context, rowsAffected)
}

func getRowByTimesheet(context *gin.Context) {
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

func getRow(context *gin.Context) {
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

func getRows(context *gin.Context) {
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

func removeRowByIds(context *gin.Context) {
	condition, err := createDeleteConditions(context)

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	rowsAffected, err := utils.DeleteFromTable(tableName, condition)

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	sendExecSuccessResponse(context, rowsAffected)
}

package timesheets

import (
	"fmt"

	"example.com/time_tracker_api/utils"
	"github.com/gin-gonic/gin"
)

func addTimesheets(context *gin.Context) {
	values := parseRequestBodyForInsertValues(context)

	if values == "" {
		return
	}

	rowsAffected, err := utils.InsertToTable(tableName, columns, values)

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	sendExecSuccessResponse(context, rowsAffected)
}

func getTimesheet(context *gin.Context) {
	id := context.Param("id")

	condition := fmt.Sprintf("timesheet_id=%s", id)

	rows, err := utils.SelectFromTable(tableName, "*", condition)

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	timesheets, err := parseRows(rows)

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	sendQuerySuccessResponse(context, timesheets)
}

func getTimesheets(context *gin.Context) {
	rows, err := utils.SelectFromTable(tableName, "*", "")

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	timesheets, err := parseRows(rows)

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	sendQuerySuccessResponse(context, timesheets)
}

func removeTimesheets(context *gin.Context) {
	condition, err := createDeleteConditions(context)

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

func updateTimesheets(context *gin.Context) {
	id := context.Param("id")

	timesheets, err := parseRequestBodyForTimesheetsUpdates(context)

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	transaction, err := utils.DB.Begin()

	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	var totalRowsAffected int64 = 0
	for _, timesheet := range timesheets {
		rowsAffected, err := utils.UpdateRowInTable(transaction, tableName, fmt.Sprintf("timesheet_name='%s'", timesheet.Name), fmt.Sprintf("timesheet_id=%s", id))

		totalRowsAffected += rowsAffected

		if err != nil {
			transaction.Rollback()
			sendErrorResponse(context, err)
			return
		}
	}

	if err := transaction.Commit(); err != nil {
		sendErrorResponse(context, err)
		return
	}

	sendExecSuccessResponse(context, totalRowsAffected)
}

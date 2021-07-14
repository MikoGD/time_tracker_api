package timesheets

import (
	"database/sql"
	"fmt"
	"net/http"

	"example.com/time_tracker_api/utils"
	"github.com/gin-gonic/gin"
)

const tableName = "timesheets"
const columns = "(timesheet_name)"

func sendErrorResponse(context *gin.Context, err error) {
	response := TimesheetsErrorResponse{fmt.Sprintf("%s", err)}
	context.JSON(http.StatusNotFound, response)
}

func sendQuerySuccessResponse(context *gin.Context, timesheets []Timesheet) {
	response := CreateQuerySuccessResponse(timesheets)
	context.JSON(http.StatusOK, response)
}

func sendExecSuccessResponse(context *gin.Context, rowsAffected int64) {
	response := CreateExecSuccessResponse(rowsAffected)
	context.JSON(http.StatusOK, response)
}

func parseRows(rows *sql.Rows) ([]Timesheet, error) {
	var timesheets []Timesheet

	for rows.Next() {
		var timesheet Timesheet

		if err := rows.Scan(&timesheet.Id, &timesheet.Name); err != nil {
			return nil, err
		}

		timesheets = append(timesheets, timesheet)
	}

	return timesheets, nil
}

func parseRequestBody(context *gin.Context) string {
	var timesheet TimesheetRequestBody

	if err := context.Bind(&timesheet); err != nil {
		sendErrorResponse(context, err)
		return ""
	}

	return fmt.Sprintf("('%s')", utils.EscapeString(timesheet.Name))
}

func GetTimesheets(context *gin.Context) {
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

func AddTimesheets(context *gin.Context) {
	values := parseRequestBody(context)

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

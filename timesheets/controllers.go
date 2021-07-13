package timesheets

import (
	"database/sql"
	"fmt"
	"net/http"

	"example.com/time_tracker_api/utils"
	"github.com/gin-gonic/gin"
)

const tableName = "timesheets"
const columns = "(name)"

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

func AddTimesheets(context *gin.Context, values string) {
	rowsAffected, err := utils.InsertToTable(tableName, columns, values)
	
	if err != nil {
		sendErrorResponse(context, err)
		return
	}

	sendExecSuccessResponse(context, rowsAffected)
}

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

func parseRequestBodyForValues(context *gin.Context) string {
	var timesheets TimesheetRequestBody

	if err := context.ShouldBindJSON(&timesheets); err != nil {
		sendErrorResponse(context, err)
		return ""
	}

	values := ""

	for i, timesheet := range timesheets.Timesheets {
		if i < len(timesheets.Timesheets) - 1 {
			values += fmt.Sprintf("('%s'), ", timesheet.Name)
		} else {
			values += fmt.Sprintf("('%s')", timesheet.Name)
		}
	}

	return values
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

func GetTimesheet(context *gin.Context) {
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

func AddTimesheets(context *gin.Context) {
	values := parseRequestBodyForValues(context)

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


func createDeleteConditions(context *gin.Context) (string, error) {
	var requestBody TimesheetRequestBody	

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		return "", err
	}

	idsString := "("

	for i, id := range requestBody.Ids {
		if i < len(requestBody.Ids) - 1 {
			idsString += fmt.Sprintf("%d, ", id)
		} else {
			idsString += fmt.Sprintf("%d)", id)
		}
	}

	return fmt.Sprintf("timesheet_id IN %s", idsString), nil
}

func RemoveTimesheets(context *gin.Context) {
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

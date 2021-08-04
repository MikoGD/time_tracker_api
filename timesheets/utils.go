package timesheets

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createDeleteConditions(context *gin.Context) (string, error) {
	var requestBody TimesheetRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		return "", err
	}

	idsString := "("

	for i, id := range requestBody.Ids {
		if i < len(requestBody.Ids)-1 {
			idsString += fmt.Sprintf("%d, ", id)
		} else {
			idsString += fmt.Sprintf("%d)", id)
		}
	}

	return fmt.Sprintf("timesheet_id IN %s", idsString), nil
}

func createExecSuccessResponse(count int64) TimesheetsSuccessReponse {
	return TimesheetsSuccessReponse{count, make([]Timesheet, 0)}
}

func createQuerySuccessResponse(timesheets []Timesheet) TimesheetsSuccessReponse {
	return TimesheetsSuccessReponse{int64(len(timesheets)), timesheets}
}

func parseRequestBodyForTimesheetsUpdates(context *gin.Context) ([]Timesheet, error) {
	var requestBody TimesheetRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		return nil, err
	}

	return requestBody.Timesheets, nil
}

func parseRequestBodyForInsertValues(context *gin.Context) string {
	var timesheets TimesheetRequestBody

	if err := context.ShouldBindJSON(&timesheets); err != nil {
		sendErrorResponse(context, err)
		return ""
	}

	values := ""

	for i, timesheet := range timesheets.Timesheets {
		if i < len(timesheets.Timesheets)-1 {
			values += fmt.Sprintf("('%s'), ", timesheet.Name)
		} else {
			values += fmt.Sprintf("('%s')", timesheet.Name)
		}
	}

	return values
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

func sendErrorResponse(context *gin.Context, err error) {
	response := TimesheetsErrorResponse{fmt.Sprintf("%s", err)}
	context.JSON(http.StatusNotFound, response)
}

func sendExecSuccessResponse(context *gin.Context, rowsAffected int64) {
	response := createExecSuccessResponse(rowsAffected)
	context.JSON(http.StatusOK, response)
}

func sendQuerySuccessResponse(context *gin.Context, timesheets []Timesheet) {
	response := createQuerySuccessResponse(timesheets)
	context.JSON(http.StatusOK, response)
}

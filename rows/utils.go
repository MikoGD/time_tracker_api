package rows

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createDeleteConditions(context *gin.Context) (string, error) {
	var requestBody TimesheetRowsRequestBody

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

	return fmt.Sprintf("row_id IN %s", idsString), nil
}

func createExecSuccessResponse(count int64) TimesheetRowsSuccessReponse {
	return TimesheetRowsSuccessReponse{count, make([]TimesheetRow, 0)}
}

func createUpdateColumns(timesheetRow TimesheetRow) string {
	return fmt.Sprintf(
		"row_description='%s', row_start_time=%d, row_end_time=%d, row_elapsed_time=%d",
		timesheetRow.Description,
		timesheetRow.StartTime,
		timesheetRow.EndTime,
		timesheetRow.ElapsedTime,
	)
}

func createQuerySuccessResponse(timesheetRows []TimesheetRow) TimesheetRowsSuccessReponse {
	return TimesheetRowsSuccessReponse{int64(len(timesheetRows)), timesheetRows}
}

func parseRequestBodyForInsertValues(context *gin.Context) (string, error) {
	var rows TimesheetRowsRequestBody

	if err := context.ShouldBindJSON(&rows); err != nil {
		return "", err
	}

	values := ""

	for i, row := range rows.TimesheetRows {
		if i < len(rows.TimesheetRows)-1 {
			values += fmt.Sprintf("('%s', %d, %d, %d, %d), ", row.Description, row.StartTime, row.EndTime, row.ElapsedTime, row.TimesheetId)
		} else {
			values += fmt.Sprintf("('%s', %d, %d, %d, %d);", row.Description, row.StartTime, row.EndTime, row.ElapsedTime, row.TimesheetId)
		}
	}

	return values, nil
}

func parseRequestBodyForUpdateValues(context *gin.Context) ([]TimesheetRow, error) {
	var requestBody TimesheetRowsRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		return nil, err
	}

	return requestBody.TimesheetRows, nil
}

func parseRows(rows *sql.Rows) ([]TimesheetRow, error) {
	var timesheetRows []TimesheetRow

	for rows.Next() {
		var timesheetRow TimesheetRow

		if err := rows.Scan(
			&timesheetRow.Id,
			&timesheetRow.Description,
			&timesheetRow.StartTime,
			&timesheetRow.EndTime,
			&timesheetRow.ElapsedTime,
			&timesheetRow.TimesheetId); err != nil {

			return nil, err
		}

		timesheetRows = append(timesheetRows, timesheetRow)
	}

	return timesheetRows, nil
}

func sendErrorResponse(context *gin.Context, err error) {
	response := TimesheetRowsErrorResponse{fmt.Sprintf("%s", err)}
	context.JSON(http.StatusNotFound, response)
}

func sendExecSuccessResponse(context *gin.Context, rowsAffected int64) {
	response := createExecSuccessResponse(rowsAffected)
	context.JSON(http.StatusOK, response)
}

func sendQuerySuccessResponse(context *gin.Context, timesheetRows []TimesheetRow) {
	response := createQuerySuccessResponse(timesheetRows)
	context.JSON(http.StatusOK, response)
}

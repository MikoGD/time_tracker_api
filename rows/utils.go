package rows

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateExecSuccessResponse(count int64) TimesheetRowsSuccessReponse {
	return TimesheetRowsSuccessReponse{count, make([]TimesheetRows, 0)}
}

func CreateQuerySuccessResponse(timesheetRows []TimesheetRows) TimesheetRowsSuccessReponse {
	return TimesheetRowsSuccessReponse{int64(len(timesheetRows)), timesheetRows}
}

func parseRows(rows *sql.Rows) ([]TimesheetRows, error) {
	var timesheetRows []TimesheetRows

	for rows.Next() {
		var timesheetRow TimesheetRows

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
	response := CreateExecSuccessResponse(rowsAffected)
	context.JSON(http.StatusOK, response)
}

func sendQuerySuccessResponse(context *gin.Context, timesheetRows []TimesheetRows) {
	fmt.Println(timesheetRows)
	response := CreateQuerySuccessResponse(timesheetRows)
	context.JSON(http.StatusOK, response)
}

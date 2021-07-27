package rows

import (
	"example.com/time_tracker_api/utils"
	"github.com/gin-gonic/gin"
)

func GetRows(context *gin.Context) {
	rows, err := utils.SelectFromTable(tableName, "*", "")

	if err != nil {
		sendErrorResponse(context, err)
	}

	timesheetRows, err := parseRows(rows)

	if err != nil {
		sendErrorResponse(context, err)
	}

	sendQuerySuccessResponse(context, timesheetRows)
}

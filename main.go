package main

import (
	"example.com/time_tracker_api/rows"
	"example.com/time_tracker_api/timesheets"
	"example.com/time_tracker_api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	utils.ConnectToDatabaase()

	timesheets.SetupTimesheetsRoutes(router)
	rows.SetupTimesheetRowsRoutes(router)

	router.Run(":8090")
}

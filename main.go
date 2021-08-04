package main

import (
	"example.com/time_tracker_api/rows"
	"example.com/time_tracker_api/timesheets"
	"example.com/time_tracker_api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnvFile()

	var router *gin.Engine

	if utils.GetEnvVariable("MODE") == "RELEASE" {
		gin.SetMode(gin.ReleaseMode)
		router = gin.Default()
	} else {
		router = gin.Default()
	}

	utils.ConnectToDatabaase()

	v1 := router.Group("v1")

	timesheets.SetupTimesheetsRoutes(v1)
	rows.SetupTimesheetRowsRoutes(v1)

	router.Run(":8090")
}

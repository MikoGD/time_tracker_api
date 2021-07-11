package main

import (
	"example.com/time_tracker_api/timesheets"
	"example.com/time_tracker_api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	utils.ConnectToDatabaase()

	timesheets.SetupTimesheetsRoutes(router)

	router.Run(":8090")
	// db := utils.ConnectToDatabaase()

	// defer db.Close()

	// rows, err := db.Query("SELECT * FROM timesheets")

	// if err != nil {
	// 	fmt.Println(fmt.Errorf("timesheets %v", err))
	// }

	// defer rows.Close()


	// for rows.Next() {
	// 	var currTimesheet timesheet.Timesheet
	// 	rows.Scan(&currTimesheet.Id, &currTimesheet.Name)
	// 	fmt.Printf("id: %s, name: %s\n", currTimesheet.Id, currTimesheet.Name)
	// }
}

package rows

type TimesheetRow struct {
	Id          uint   `json:"id"`
	Description string `json:"description" binding:"alpha"`
	StartTime   uint   `json:"startTime"`
	EndTime     uint   `json:"endTime"`
	ElapsedTime uint   `json:"elapsedTime"`
	TimesheetId uint   `json:"timesheetId"`
}

type TimesheetRowsRequestBody struct {
	TimesheetRows []TimesheetRow `json:"timesheetRows"`
	Ids           []uint         `json:"ids"`
}

type TimesheetRowsSuccessReponse struct {
	Count int64
	Data  []TimesheetRow
}

type TimesheetRowsErrorResponse struct {
	ErrorMessage string
}

const tableName = "timesheet_rows"
const columns = "(row_description, row_start_time, row_end_time, row_elapsed_time, timesheet_id)"

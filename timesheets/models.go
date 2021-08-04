package timesheets

const tableName = "timesheets"
const columns = "(timesheet_name)"

type Timesheet struct {
	Id   string `json:"id" binding:"numeric"`
	Name string `json:"name" binding:"alpha"`
}

type TimesheetRequestBody struct {
	Timesheets []Timesheet `json:"timesheets"`
	Ids        []uint      `json:"ids"`
}

type TimesheetsSuccessReponse struct {
	Count int64
	Data  []Timesheet
}

type TimesheetsErrorResponse struct {
	ErrorMessage string
}

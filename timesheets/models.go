package timesheets

type Timesheet struct {
	Id string
	Name string
}

type TimesheetRequestBody struct {
	Name string `json:"name" binding:"required alpha`
}

type TimesheetsSuccessReponse struct {
	Count int64
	Data []Timesheet
}

type TimesheetsErrorResponse struct {
	ErrorMessage string
}

func CreateQuerySuccessResponse(timesheets []Timesheet) TimesheetsSuccessReponse {
	return TimesheetsSuccessReponse{int64(len(timesheets)), timesheets}
}

func CreateExecSuccessResponse(count int64) TimesheetsSuccessReponse {
	return TimesheetsSuccessReponse{count, make([]Timesheet, 0)}
}
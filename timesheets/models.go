package timesheets

type Timesheet struct {
	Id string
	Name string
}

type TimesheetsSuccessReponse struct {
	Count int
	Data []Timesheet
}

type TimesheetsErrorResponse struct {
	ErrorMessage string
}


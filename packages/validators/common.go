package validators

type ValidationErrorResponse struct {
	FailedField string
	Tag         string
	Param       string
	Value       interface{}
}

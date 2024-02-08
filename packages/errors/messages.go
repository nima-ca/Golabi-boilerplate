package errors

const (
	InternalServerErrorErrorMsg     string = "Something went wrong, please try again later!"
	FailedToParseBodyErrorMsg       string = "Failed to Parse Request Body"
	InvalidPaginationParamsErrorMsg string = "Invalid Page or Count in params"
	InvalidIDParamErrorMsg          string = "Invalid ID in params"
	FileIsNotPDFErrorMsg            string = "Uploaded file is not pdf"
	FailedToParseFileErrorMsg       string = "Failed to parse uploaded file"
)

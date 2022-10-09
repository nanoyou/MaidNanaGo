package response

type Response struct {
	Ok bool `json:"ok"`
}

type SuccessResponse struct {
	Response
	SuccessMessage string
}

type FailureResponse struct {
	Response
	Error        string
	ErrorMessage string
}

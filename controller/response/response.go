package response

type Response struct {
	Ok bool `json:"ok"`
}

type SuccessResponse struct {
	Response
	SuccessMessage string `json:"success_message"`
}

type FailureResponse struct {
	Response
	Error        string `json:"error"`
	ErrorMessage string `json:"error_message"`
}

package response

type Response struct {
	Ok bool `json:"ok"`
}

type SuccessResponse struct {
	Response
}

type ErrorResponse struct {
	Response
	ErrorMessage string `json:"error_message"`
}

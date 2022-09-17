package request

type RegisterRequest struct {
	Username         string `json:"username" validate:"required,min=6,max=20,urlsafename"`
	Password         string `json:"password" validate:"required,min=6,max=32"`
	VerificationCode int    `json:"verification_code" validate:"required"`
}

type LoginRequest struct {
	Password string `json:"password" validate:"required"`
}

package request

type RegisterRequest struct {
	Username         string `json:"username" validate:"required"`
	Password         string `json:"password" validate:"required"`
	VerificationCode int    `json:"verification_code" validate:"required"`
}

type LoginRequest struct {
	Password string `json:"password" validate:"required"`
}

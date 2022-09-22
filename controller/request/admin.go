package request

type AdminModifyUserRequest struct {
	Username string `json:"username" validate:"max=20,urlsafename"`
	Password string `json:"password" validate:"max=32"`
}

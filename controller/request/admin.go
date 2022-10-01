package request

type AdminModifyUserRequest struct {
	Username string `json:"username" validate:"min=4,max=20,urlsafename"`
	Password string `json:"password" validate:"min=6,max=32"`
}

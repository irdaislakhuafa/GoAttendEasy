package request

type AuthRegister struct {
	Name     string `json:"name,omitempty"     validate:"required"`
	Email    string `json:"email,omitempty"    validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
	RoleID   string `json:"role_id,omitempty"  validate:"required"`
}

type AuthLogin struct {
	Email    string `json:"email,omitempty"    validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

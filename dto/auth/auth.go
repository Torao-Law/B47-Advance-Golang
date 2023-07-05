package authdto

type AuthRequest struct {
	Name     string `json:"name" validate:"required" form:"name"`
	Email    string `json:"email" validate:"required" form:"email"`
	Password string `json:"password" validate:"required" form:"password"`
}

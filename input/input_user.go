package input

type RegisterUserInput struct {
	Email      string `json:"email" form:"email" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	FullName       string `json:"full_name" form:"full_name" binding:"required"`
	Role string `json:"role" form:"role" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UpdateUserInput struct {
	FullName       string `json:"full_name" form:"full_name"`
	Email      string `json:"email" form:"email"`
	Error      error
}

type FindByEmailInput struct {
	Email    string `json:"email" form:"email" binding:"required"`
}
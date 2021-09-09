package users

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdater struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

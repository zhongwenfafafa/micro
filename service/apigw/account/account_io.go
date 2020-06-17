package account

type Register struct {
	Username string `json:"username" binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required,gte=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type SignIn struct {
	Username string `json:"username" binding:"required_or=Mobile"`
	Mobile string `json:"mobile"`
	Password string `json:"password"`
	VerifyCode string `json:"verify_code"`
}
package web

type AuthLoginRequest struct {
	Username string `json:"username" binding:"required,max=12"`
	Password string `json:"password"`
}

type UserFromJWT struct {
	Username string `json:"username"`
	Exp      string `json:"exp"`
}

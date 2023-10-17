package iggcon

type LogInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogInAccessTokenRequest struct {
	Token string `json:"token"`
}

type LogInResponse struct {
	UserId uint32 `json:"userId"`
}

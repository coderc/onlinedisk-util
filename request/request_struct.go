package request

type RequestUserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

package response

type Register struct {
	Token string `json:"token"`
}

type Login struct {
	Token string `json:"token"`
}

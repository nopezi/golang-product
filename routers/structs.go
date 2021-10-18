package routers

type Result struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ResultToken struct {
	Token string `json:"token"`
}
package services

type PingService interface {
	Ping() (*PingResponse, error)
}

type PingResponse struct {
	Message  string `json:"message"`
}

package entities

type ErrResponse struct {
	Status     string `json:"status"`
	StatusCode uint32 `json:"status_code"`
	Message    string `json:"message"`
}

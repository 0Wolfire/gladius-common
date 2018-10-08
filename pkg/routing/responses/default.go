package responses

type DefaultResponse struct {
	Message     string      `json:"message"`
	Success     bool        `json:"success"`
	Error       string      `json:"error"`
	Response    interface{} `json:"response"`
	Transaction *TxHash     `json:"txHash"`
	Endpoint    string      `json:"endpoint"`
}

package model

type Result map[string]interface{}

type HTTPResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"status_code"`
	Result     Result `json:"result"`
}

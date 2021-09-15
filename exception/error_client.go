package exception

import "encoding/json"

type ErrorClient struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (e *ErrorClient) Error() string {
	errBytes, _ := json.Marshal(e)
	return string(errBytes)
}

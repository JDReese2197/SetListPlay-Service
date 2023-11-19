package setlistFmModels

import "fmt"

type ResponseError struct {
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func (e ResponseError) ToString() string {
	return fmt.Sprintf(" (%s) - code: %d, status: %s, message: %s",
		e.Timestamp, e.Code, e.Status, e.Message)
}

package bb

import "fmt"

type ErrorBB struct {
	Errors []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}

func (e *ErrorBB) Error() string {
	if len(e.Errors) > 0 {
		f := e.Errors[0]
		return fmt.Sprintf("Code: %v, message: %s",
			f.Code, f.Message)
	}
	return ""
}

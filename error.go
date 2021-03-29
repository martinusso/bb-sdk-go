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

type ErrosV1 struct {
	Erros []struct {
		Code    string `json:"codigo"`
		Message string `json:"mensagem"`
	} `json:"erros"`
}

func (e *ErrosV1) Error() string {
	if len(e.Erros) > 0 {
		f := e.Erros[0]
		return f.Message
	}
	return ""
}

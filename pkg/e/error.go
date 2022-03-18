package e

import "fmt"

func MakeCommonErr(message string, err *error, code int) SixCommonError {
	if err == nil {
		return SixCommonError{Text: message, Code: code}
	} else {
		return SixCommonError{Text: fmt.Sprintf("%s:【%v】", message, *err), Code: code}
	}
}

// 通用错误
type SixCommonError struct {
	Code int
	Text string
}

func (e SixCommonError) Error() string {
	return fmt.Sprintf("%v", e.Text)
}

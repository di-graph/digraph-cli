package error_handling

import "fmt"

type RequestError struct {
	StatusCode int
	Err        error
}

const FAIL = "\033[91m"
const ENDC = "\033[0m"

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func FormatError(input string) string {
	return fmt.Sprintf("\n%s%s%s\n", FAIL, input, ENDC)
}

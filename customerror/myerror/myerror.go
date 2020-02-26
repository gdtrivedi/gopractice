package myerror

// MyError is a trivial implementation of error.
type MyError struct {
	s string
	c string
}

// New returns an error that formats as the given text.
func New(text string, code string) error {
	return &MyError{text, code}
}

func (e *MyError) Error() string {
	return e.c + ": " + e.s
}

func (e *MyError) ErrCode() string {
	return e.c
}

package myerror

// MyError is a trivial implementation of error.
type MyError struct {
	s string
}

// New returns an error that formats as the given text.
func New(text string) error {
	return &MyError{text}
}

func (e *MyError) Error() string {
	return e.s
}

package MyError

func New(text string) error {
	return &myerrorString{text}
}

// errorString is a trivial implementation of error.
type myerrorString struct {
	s string
}

func (e *myerrorString) Error() string {
	return e.s
}

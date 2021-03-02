package apperror

type Error struct {
	Text       string
	StatusCode int
}

func (e Error) Error() string {
	return e.Text
}

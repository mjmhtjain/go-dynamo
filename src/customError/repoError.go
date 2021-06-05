package customError

type RecordNotFoundError struct {
	Message string
	Err     error
}

func (r *RecordNotFoundError) Error() string {
	return r.Message
}

func (r *RecordNotFoundError) Unwrap() error {
	return r.Err
}

func NewRecordNotFoundError(m string, err error) error {
	return &RecordNotFoundError{Message: m, Err: err}
}

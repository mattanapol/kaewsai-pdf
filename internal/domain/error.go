package domain

type MissingIdError struct {
}

func (e *MissingIdError) Error() string {
	return "id is required"
}

type InvalidRequestTypeError struct {
}

func (e *InvalidRequestTypeError) Error() string {
	return "invalid request type"
}

package error_code

type ErrorCode string

const (
	Server     ErrorCode = "0000"
	Input      ErrorCode = "0001"
	Acceptable ErrorCode = "0002"
)

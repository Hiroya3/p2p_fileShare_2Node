package errorStatus

type ErrorCode int

const (
	ErrCode300 ErrorCode = 300
)

func (e ErrorCode) Error() string {
	switch e {
	case ErrCode300:
		return "error code 300"
	default:
		return "unknown error code"
	}
}

func ReturnErrorCode300() error {
	return ErrCode300
}

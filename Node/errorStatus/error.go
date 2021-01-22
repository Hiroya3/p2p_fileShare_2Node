package errorStatus

type ErrorCode int

const (
	ErrCode300 ErrorCode = 300 //チェックサムが一致せず改竄があったとき
)

func (e ErrorCode) Error() string {
	switch e {
	case ErrCode300:
		return "tampering happend because checksum is wrong"
	default:
		return "unknown error code"
	}
}

//改竄があったときのエラーコードを返す
func ReturnErrorCode300() error {
	return ErrCode300
}

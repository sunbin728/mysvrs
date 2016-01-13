package errors

type MyError struct {
	code int64
	msg  string
}

func (err *MyError) Error() string {
	return err.msg
}

func (err *MyError) GetMessage() string {
	return err.msg
}

func (err *MyError) GetCode() int64 {
	return err.code
}

const (
	RET_SUC       = 0
	RET_DATAERROR = 1
)

var (
	DataError = &MyError{RET_DATAERROR, "DataError"}
)

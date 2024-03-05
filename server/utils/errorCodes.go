package utils

type ErrorCode int

const (
	ErrUnauthorized ErrorCode = 0x01
)

func (ec ErrorCode) GetMessage() string {
	switch ec {
	case ErrUnauthorized:
		return "Unauthorized"
	}

	return ""
}

func (ec ErrorCode) Code() int {
	return int(ec)
}

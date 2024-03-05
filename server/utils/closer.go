package utils

func Closer(closer func() error) {
	_ = closer()
}

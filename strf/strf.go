package strf

import "fmt"

func Noop(msg string) string {
	return msg
}

func F(msg string, a ...any) string {
	return fmt.Sprintf(msg, a...)
}

func F2(msg string, a ...any) string {
	return F(msg, a...)
}

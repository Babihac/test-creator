package types

type IStep interface {
	LteFieldError(fieldName string) string
}

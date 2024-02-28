package types

type IForm interface {
	LteFieldError(fieldName string) string
}

package differ

import (
	"fmt"
	"strconv"
)

func StrconvInt(i int) string {
	return strconv.Itoa(i)
}

func StrconvString(str string) string {
	return str
}

func StrconvAny(e any) string {
	return fmt.Sprintf("%v", e)
}

func StrconvFmt[T any](e T) string {
	return StrconvAny(e)
}

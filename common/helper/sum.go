package helper

import (
	"strconv"
)

func Sum(a, b int) int {
	return a + b
}

func Diff(a, b int) string {
	return strconv.Itoa(a - b)
}

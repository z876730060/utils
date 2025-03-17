package utils

import (
	"strconv"

	"github.com/pkg/errors"
)

// StrToInt string 转换 int
func StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		Errorf("StrToInt value: %v, error: %+v\n", str, errors.New(err.Error()))
	}
	return i
}

// StrToInt64 string 转换 int64
func StrToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		Errorf("StrToInt64 value: %v, error: %+v\n", str, errors.New(err.Error()))
	}
	return i
}

// StrToFloat64 string 转换 float64
func StrToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		Errorf("StrToFloat64 value: %v, error: %+v\n", str, errors.New(err.Error()))
	}
	return f
}

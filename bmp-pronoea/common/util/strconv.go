package util

import (
	"strconv"
)

func TurnStringToFloat64(numberStr string, prec int) (float64, error) {
	float64F, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		return 0, err
	}
	return Float64WithPrec(float64F, prec)
}
// Float64WithPrec float64 保留小数点后几位
func Float64WithPrec(inFloat64 float64, prec int) (float64, error) {
	float64With, err := strconv.ParseFloat(strconv.FormatFloat(inFloat64, 'f', prec, 64), 64)
	if err != nil {
		return 0, err
	}

	return float64With, nil
}

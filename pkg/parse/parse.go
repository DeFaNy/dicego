package parse

import (
	"errors"
	"math"
	"strconv"
)

func FloatWithoutExtra(snum string) (num float64, err error) {
	n, err := strconv.ParseFloat(snum, 64)
	if err != nil || math.IsNaN(n) || math.IsInf(n, 0) {
		return 0, errors.New("invalid number")
	}

	return n, err
}

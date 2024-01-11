package validator

import (
	"math"
	"strconv"
	"strings"
)

func Validate(payload string, checkDigit int) bool {
	split := strings.Split(payload, "")

	sum := 0
	for i := len(split) - 1; i >= 0; i-- {
		el, err := strconv.Atoi(split[i])
		if err != nil {
			return false
		}
		if (i+1)%2 == 0 {
			multiplied := el * 2
			floor := math.Floor(float64(multiplied / 10))
			remainder := float64(multiplied) - floor*10
			sum += int(floor) + int(remainder)

		} else {
			sum += el
		}
	}
	eval := 10-(sum%10) == checkDigit
	return eval
}

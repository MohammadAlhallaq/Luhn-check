package luhnAlgo

import (
	"strconv"
)

func Valid(number int) bool {

	runes := strconv.Itoa(number)
	length := len(runes)
	sum := 0
	alternate := false

	for i := length - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(runes[i]))
		if err != nil {
			return false
		}

		if alternate {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		alternate = !alternate
	}
	return sum%10 == 0
}

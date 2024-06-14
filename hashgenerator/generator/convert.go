package generator

import "strconv"

func convertNumber(num int) (newNum string) {
	for num != 0 {
		symbInt := num % 62
		var symbStr string
		if symbInt >= 0 && symbInt <= 9 {
			symbStr = strconv.Itoa(symbInt)
		}
		if symbInt >= 10 && symbInt <= 35 {
			symbStr = string(rune(symbInt - 10 + int('A')))
		}
		if symbInt >= 36 && symbInt <= 61 {
			symbStr = string(rune(symbInt - 36 + int('a')))
		}

		newNum += symbStr

		num = num / 62
	}

	return
}

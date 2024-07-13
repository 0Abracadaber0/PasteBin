package converter

import (
	"fmt"
	"strconv"
)

func FormatingTime(time string) string {

	year, _ := strconv.Atoi(time[:4])
	month, _ := strconv.Atoi(time[5:7])
	day, _ := strconv.Atoi(time[8:10])
	hour, _ := strconv.Atoi(time[11:13])
	minute, _ := strconv.Atoi(time[14:])

	hour += 3
	if hour >= 24 {
		hour -= 24
		day++
	}

	days := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		days[2]++
	}

	if day > days[month] {
		day -= days[month]
		month++
	}

	if month > 12 {
		month -= 12
		year++
	}

	newTime := fmt.Sprintf(
		"%04d-%02d-%02d %02d:%02d:00:00",
		year, month, day, hour, minute,
	)

	return newTime
}

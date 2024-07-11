package converter

func FormatingTime(time string) (string, error) {
	newTime := time[:10] + " " + time[11:] + ":00:00 +03:00"

	return newTime, nil
}

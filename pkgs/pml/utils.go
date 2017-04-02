package pml

import (
	"strings"
	"strconv"
)

const SECONDS_TO_MIN = 60
const SECONDS_TO_HOUR = 3600
const SECONDS_TO_DAY = 86400
const SECONDS_TO_WEEK = 604800
const SECONDS_TO_MONTH = 2592000

func convertToSeconds(timeValue string) (timeInSeconds int) {
	result := strings.Split(timeValue, " ")
	num, _ := strconv.Atoi(result[0])
	unit := result[1]

	switch unit {
	case "min","mins":
		timeInSeconds = num * SECONDS_TO_MIN
	case "hr","hrs":
		timeInSeconds = num * SECONDS_TO_HOUR
	case "day","days":
		timeInSeconds = num * SECONDS_TO_DAY
	case "week","weeks":
		timeInSeconds = num * SECONDS_TO_WEEK
	case "month", "months":
		timeInSeconds = num * SECONDS_TO_MONTH
	default:
		timeInSeconds = num //assume it is seconds
	}
	return timeInSeconds
}


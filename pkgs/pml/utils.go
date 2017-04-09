package pml

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const SECONDS_TO_MIN = 60
const SECONDS_TO_HOUR = 3600
const SECONDS_TO_DAY = 86400
const SECONDS_TO_WEEK = 604800
const SECONDS_TO_MONTH = 2592000

var (
	WEEKDAYS        = []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	DAYTIMES        = []string{"morning", "afternoon", "evening"}
	START_TIMESTAMP = int64(378000) // 378000 is the first monday at 9am timestamp
)

func convertToSeconds(timeValue string) (timeInSeconds int) {
	result := strings.Split(timeValue, " ")
	num, _ := strconv.Atoi(result[0])
	unit := result[1]

	switch unit {
	case "min", "mins":
		timeInSeconds = num * SECONDS_TO_MIN
	case "hr", "hrs":
		timeInSeconds = num * SECONDS_TO_HOUR
	case "day", "days":
		timeInSeconds = num * SECONDS_TO_DAY
	case "week", "weeks":
		timeInSeconds = num * SECONDS_TO_WEEK
	case "month", "months":
		timeInSeconds = num * SECONDS_TO_MONTH
	default:
		timeInSeconds = num //assume it is seconds
	}
	return timeInSeconds
}

func secondsToHumanReadable(seconds int) string {

	if seconds/SECONDS_TO_MONTH > 0 {
		return fmt.Sprintf("%d months", seconds/SECONDS_TO_MONTH)
	} else if seconds/SECONDS_TO_WEEK > 0 {
		return fmt.Sprintf("%d weeks", seconds/SECONDS_TO_WEEK)
	} else if seconds/SECONDS_TO_DAY > 0 {
		return fmt.Sprintf("%d days", seconds/SECONDS_TO_DAY)
	} else if seconds/SECONDS_TO_HOUR > 0 {
		return fmt.Sprintf("%d hrs", seconds/SECONDS_TO_HOUR)
	} else if seconds/SECONDS_TO_MIN > 0 {
		return fmt.Sprintf("%d mins", seconds/SECONDS_TO_MIN)
	} else {
		return fmt.Sprintf("%d secs", seconds)
	}
}

/*
  Monday Morning => Monday 9am
  Tuesday Afternoon => Tuesday 2pm
  Friday Evening => Friday 7pm
  Wednesday => Wednesday Morning
  Evening => This Evening
  Monday 9pm
  5pm
  10am
*/
func calculateOffsetDelay(currentDelay Delay, offset Wait) Delay {
	arr := strings.Split(string(offset), " ")
	hrs := 9
	day := WEEKDAYS[time.Now().Weekday()] // default to current day

	for _, ele := range arr {
		if contains(DAYTIMES, ele) {
			switch ele {
			case "morning":
				hrs = 9
			case "afternoon":
				hrs = 14
			case "evening":
				hrs = 19
			}
		}
		if contains(WEEKDAYS, ele) {
			day = ele
		}
		if len(ele) < 5 && (strings.HasSuffix(ele, "am") || strings.HasSuffix(ele, "pm")) {
			hrs, _ = strconv.Atoi(ele[:len(ele)-2])
		}
	}

	tm := time.Unix(START_TIMESTAMP+int64(currentDelay), 0)
	numberOfDays := 0
	for tm.Weekday() != indexOf(WEEKDAYS, day) {
		tm = tm.AddDate(0, 0, 1)
		numberOfDays++
	}

	startOfDay := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, tm.Location())
	secondsSinceMidnight := int(tm.Sub(startOfDay).Seconds())

	timestampOffset := hrs * 60 * 60
	currentDayDifference := timestampOffset - secondsSinceMidnight

	delay := int(currentDelay) + (numberOfDays * 24 * 60 * 60) + currentDayDifference
	return Delay(delay)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.ToLower(a) == strings.ToLower(e) {
			return true
		}
	}
	return false
}

func indexOf(s []string, e string) time.Weekday {
	for i, a := range s {
		if strings.ToLower(a) == strings.ToLower(e) {
			return time.Weekday(i)
		}
	}
	return time.Weekday(1)
}

func writeProcessToFile(process *Element, fileName string) {
	bytes := []byte(process.Encode("  "))
	ioutil.WriteFile(resDir+"/"+fileName, bytes, 0644)
}

func drugPairListsEqual(listOne, listTwo []DrugPair) (equal bool) {
	for _, x := range listOne {
		if !drugPairListContains(listTwo, x) {
			return false
		}
	}
	return true
}

func drugPairListContains(drugPairs []DrugPair, elem DrugPair) (contains bool) {
	for _, x := range drugPairs {
		if x == elem {
			return true
		}
	}
	return false
}

func JoinPMLProcesses(processes []*Element) (joinedProcess *Element) {
	joinedProcess = &Element{
		Name:        "merged",
		Children:    []ElementInterface{},
		Loops:       0,
		elementType: ProcessType,
	}
	branch := &Element{
		Name:        "mergedpathways",
		Children:    []ElementInterface{},
		Loops:       0,
		elementType: BranchType,
	}
	joinedProcess.Children = append(joinedProcess.Children, branch)
	for i, process := range processes {
		process.ChangeNames("_" + strconv.Itoa(i+1))
		process.elementType = SequenceType
		branch.Children = append(branch.Children, process)
	}
	return
}

func (el *Element) ChangeNames(modifier string) {
	el.Name = el.Name + modifier
	for _, child := range el.Children {
		child.ChangeNames(modifier)
	}
}

func (act *Action) ChangeNames(modifier string) {
	act.Name = act.Name + modifier
}

func (delay Delay) ChangeNames(modifier string) {
	return
}

func (w Wait) ChangeNames(modifier string) {
	return
}

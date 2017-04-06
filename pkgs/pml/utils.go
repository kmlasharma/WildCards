package pml

import (
	"strconv"
	"strings"
	"fmt"
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

//func JoinPMLProcesses(processes []*Element) (joinedProcess *Element) {
//	joinedProcess, _ = processFromFile("base.pml")
//	for _, process := range processes {
//		CombineProcesses(joinedProcess, process)
//	}
//	return
//}

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

func CombineProcesses(baseProcess, processToAdd *Element) {
	for _, child := range processToAdd.Children {
		baseProcess.Children = append(baseProcess.Children, child)
	}
	for i, child := range baseProcess.Children {
		fmt.Println(fmt.Sprintf("%s: %s", i, child.GetName()))
	}
}


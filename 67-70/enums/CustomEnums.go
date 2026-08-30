package main

import "fmt"

type weekdays int

const (
	saturday weekdays = iota
	sunday
)

func getwordDay(day weekdays) string {
	switch day {
	case saturday:
		return "offday"
	case sunday:
		return "on day"
	default:
		return "invalid"

	}
}

func CustomEnums() {
	fmt.Println(getwordDay(sunday))
	fmt.Println(getwordDay(saturday))

}

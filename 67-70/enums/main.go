package main

import "fmt"

type weekday int

const (
	Monday weekday = iota
	TuesDay
	WednessDay
	thursDay
	SaturDay
	FriDay
)

func getworkdayStatus(day weekday) string {
	switch day {
	case Monday, TuesDay, WednessDay:
		return "Office is open"
	case thursDay:
		return "Half day open"
	case SaturDay, FriDay:
		return "off day"
	default:
		return "invalid day"
	}
}

type officeStatus string

const (
	statusopen    officeStatus = "open"
	statusclose   officeStatus = "close"
	statusHalfDay officeStatus = "half_day"
)

func main() {
	fmt.Println(getworkdayStatus(Monday))
	fmt.Println(getworkdayStatus(FriDay))
	fmt.Println(statusHalfDay)

}

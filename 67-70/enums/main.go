package main

// Enum (Enumeration) হলো আগে থেকে নির্দিষ্ট করে রাখা কিছু value-এর তালিকা।
// Go-তে আলাদা enum keyword নেই। সাধারণত const + custom type দিয়ে Enum-এর মতো তৈরি করা হয়।

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
	// fmt.Println(getworkdayStatus(Monday))
	// fmt.Println(getworkdayStatus(FriDay))
	// fmt.Println(statusHalfDay)
	CustomEnums()

}

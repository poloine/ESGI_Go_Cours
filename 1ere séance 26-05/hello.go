package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(math.Pow(64, 8))
	fmt.Println(time.Now().Date())
	hourMinute := strconv.Itoa(time.Now().Hour()) + ":" + strconv.Itoa(time.Now().Minute())
	fmt.Println(hourMinute)
	currentDate := time.Now()
	fmt.Println(currentDate)
	birthday := time.Date(2000, time.December, 28, 0, 0, 0, 0, time.Local)
	fmt.Println(birthday)
	fmt.Println(dateDiff(currentDate, birthday).Year())
}

func dateDiff(a, b time.Time) (c time.Time) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year := int(y2 - y1)
	month := int(M2 - M1)
	day := int(d2 - d1)
	hour := int(h2 - h1)
	minute := int(m2 - m1)
	second := int(s2 - s1)

	// Normalize negative values
	if second < 0 {
		second += 60
		minute--
	}
	if minute < 0 {
		minute += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
}

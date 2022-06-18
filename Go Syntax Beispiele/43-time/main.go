package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, loc)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())

	p(then.Add(diff))
	p(then.Add(-diff))
}

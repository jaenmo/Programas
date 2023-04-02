package main

import ("fmt"
		"github.com/jaenmo/Programs/calendar")

func main(){
	date:= calendar.Date{}
	date.Year = 2019
	date.Day = 50
	date.Month = 14

	fmt.Println(date)
	date = calendar.Date{Year: -1, Month: -3, Day: -4}
	fmt.Println(date)


}
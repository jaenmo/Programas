//Add the grade and tell if the student passed or not. Over 60 = passed

package main

import "fmt"

func main() {
	fmt.Println("Type your grade, please: ")
	var value float64
	fmt.Scan(&value)

	if value < 60 && value >= 0 {
		fmt.Println("Oops sorry but you are stupid")
	} else if value >= 60 && value <= 100 {
		fmt.Println("Yeah baby, you got it")
	} else if value < 0 || value > 100 {
		fmt.Println("WTF are you typing?!")
	}
}

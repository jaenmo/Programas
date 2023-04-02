//Average calculates the average of several numbers from a file

package main

import (
	"fmt"
	//"github.com/jaenmo/Programs/datafile"
	//"github.com/jaenmo/Programs/average/readfile"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	var sum float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(average(number...))
	// 	sum += number
	// }
	// sampleCount := float64(len(arguments))
	// fmt.Printf("Average = %.2f", sum/sampleCount)
}

func average(number ...float64)float64{
	var sum float64
	for _, v:= range number{
		sum += v
	}
	return sum/float64(len(number))
}
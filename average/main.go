//Average calculates the average of several numbers from a file

package main

import (
	"fmt"
	"github.com/jaenmo/Programs/datafile"
	//"github.com/jaenmo/Programs/average/readfile"
	"log"
)

func main() {
	numbers, err := datafile.GetFloat("data.txt")

	if err != nil {
		log.Fatal(err)
	}
	var sum float64 
	for _, number := range numbers {
		sum += number
		
	}
	sampleCount := float64(len(numbers))
		fmt.Printf("Average: %.2f\n", sum/sampleCount)
}



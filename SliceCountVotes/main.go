package main

import ("fmt"
		"github.com/jaenmo/Programs/datafile"
		"log"
)

func main(){
	slice, err := datafile.GetString("votos.txt")
	if err != nil {
		log.Fatal(err)
	}
	var count []int
	var names []string

	for _, line := range slice{
		matched := false
		for i, name := range names{
			if name == line{
				count[i]++
				matched = true
			}
		}
		if matched == false{
			names = append(names, line)
			count = append(count, 1)
		}
	
	}
	for i, name := range names {
		fmt.Printf("%s, %d\n", name, count[i])

	}
	}

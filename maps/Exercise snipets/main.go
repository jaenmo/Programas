package main

import (
	"fmt"
)

func main() {
	ranks := map[string]int{
		"gold":   1,
		"silver": 2,
		"bronze": 3,
	}

	for medal, rank := range ranks {

		fmt.Printf("The %s medal-s rank is %d\n", medal, rank)
	
	}
	type mesa struct{
		material string
		area float64
	}
	var cocina mesa
	cocina.area = 45.65
	cocina.material = "madera"

	fmt.Println(cocina)
elementos(c.cocina)
	
}
func elementos (c cocina){
	fmt.Println(c.area)
	fmt.Println(c.material)
} 

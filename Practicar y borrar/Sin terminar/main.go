
package main

import (
	"fmt"
	"strconv"
)

func Label(colors []string) string {

	colorValues := map[string]int{
		"Black":  0,
		"Brown":  1,
		"Red":    2,
		"Orange": 3,
		"Yellow": 4,
		"Green":  5,
		"Blue":   6,
		"Violet": 7,
		"Grey":   8,
		"White":  9}

	color1, color2, color3 := colors[0], colors[1], colors[2]

	// // calculate resistance value
	mainValue := (colorValues[color1]*10 + colorValues[color2]) * powerOf10(colorValues[color3])
	s := strconv.Itoa(mainValue)
	
	return s
	//fmt.Println(value, "ohms")
}

func powerOf10(n int) int {
	if n == 0 {
		return 1
	}
	return 10 * powerOf10(n-1)
}

func main() {
	color1 := "Orange"
	color2 := "Orange"
	color3 := "Orange"
	colors := []string{color1, color2, color3}

	fmt.Println(Label(colors))
}

// value 1 *10, followed by value 2, and value 3 is the number of zeros
// orange-orange-black = 33 ohms
// giga	G	109	1000000000	1960
// mega	M	106	1000000	1873
// kilo	k	103	1000	1795
// hecto	h	102	100	1795
// deca	da	101	10

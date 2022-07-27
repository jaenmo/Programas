//Calculate how many litres of paint is needed to paint the wall.
//1l can paint 10 m2

package main

import (
	"fmt"
	"log"
)

func main(){
	var amount, total float64

	amount, err := paintNeeded(4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f l needed\n", amount)
	
	total += amount

	amount, err = paintNeeded(5.2, 3.5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f l needed\n", amount)

	total += amount
	fmt.Printf("Total: %.2f l needed\n", total)

}
func paintNeeded(w, l float64)(float64, error){
	if w < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid\n", w)
	}
	if l < 0 {
		return 0, fmt.Errorf("a length of %.2f is invalid\n", l)
	}

	return w * l /10, nil
 
}

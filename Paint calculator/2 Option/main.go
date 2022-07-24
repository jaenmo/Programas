//Calculate how many litres of paint is needed to paint the wall.
//1l can paint 10 m2

package main

import "fmt"

func main(){
	paintNeeded(4.2, 3.0)
	paintNeeded(5.2, 3.5)
}
func paintNeeded(w, l float64)float64{
	area := w * l	
	fmt.Printf("%.2f liters needed\n", area/10)

}
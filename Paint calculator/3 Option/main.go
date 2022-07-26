//Calculate how many litres of paint is needed to paint the wall.
//1l can paint 10 m2

package main

import "fmt"
		//"errors")

func main(){
	//err := errors.New("Number can't be negative")
	//fmt.
	var amount, total float64

	// if err != nil{
	// 	log.Fatal(err)
	// }

	amount = paintNeeded(4.2, 3.0)
	fmt.Printf("%.2f l needed\n", amount)

	total += amount

	amount = paintNeeded(5.2, 3.5)
	fmt.Printf("%.2f l needed\n", amount)

	total += amount
	fmt.Printf("Total: %.2f l needed\n", total)

	// myInt := 3
	// myBool := false
	// myString := "bye"
	// myInt, myBool, myString := manyReturns()
	fmt.Println(manyReturns())

}
func paintNeeded(w, l float64)(float64){

	return w * l /10
 

}

func manyReturns() (int, bool, string){
	return 1, true, "hola"
}
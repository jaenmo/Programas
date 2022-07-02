package main

import "fmt"

func main() {

	// clothes1 := []string{"Folded & Hung", "299.75"}
	// clothes2 := []string{"Folded & Hung", "700"}
	// clothes3 := []string{"American Eagle", "500"}

	// m := map[string][]string{
	// 	"T-shirt":     clothes1,
	// 	"Jeans pants": clothes2,
	// 	"Shorts Jean": clothes3,
	// }

	// for prenda, value := range m {
	// 	fmt.Println("prenda", prenda)
	// 	for _, value2 := range value {
	// 		fmt.Println("\t", value2)
	// 	}
	// }

	dinero := map[string]float32{
		"T-shirt":     299.75,
		"Jeans pants": 500.25,
		"Shorts Jean": 700.70,
		"coat":        860.50,
		"Cap":         299,
		"thongs":      300,
		"Top":         1599,
		"Dress":       799,
	}
	// for piece, price := range m1 {
	// 	fmt.Println(piece)
	// 	for _, precio := range price {
	// 		fmt.Println("\t", precio)
	// 	}
	// }
	if dinero, ok := dinero["Top"]; ok {
		if dinero < 500 {
			fmt.Println("its a good price")
		} else if dinero >= 500 && dinero < 800 {
			fmt.Println("it's expensive, but we can buy it")
		} else {
			fmt.Println("no way!")
		}
	}
}

package main

import (
    "fmt"
   // "math"
)

// func InterestRate(balance float64) float32 {
// 	if balance < 0 {
// 		return float32(3.213)
// 	}else if balance >= 0 && balance < 1000{
// 		return float32(0.5)
// 	}else if balance >= 1000 && balance < 5000{
// 		return float32(1.621)
// 	}
// 	return float32(2.475)
// }

// func Interest(balance float64) float64 {
// 	if balance < 0 {
// 		return balance * 3.213/100
// 	}else if balance >= 0 && balance < 1000{
// 		return balance * 0.5/100
// 	}else if balance >= 1000 && balance < 5000{
// 		return balance * 1.621/100
// 	}
// 	return balance * 2.475/100
// }

// func AnnualBalanceUpdate(balance float64) float64 {
// 	if balance < 0 {
// 		return balance * 3.213/100 + balance
// 	}else if balance >= 0 && balance < 1000{
// 		return balance * 0.5/100 + balance
// 	}else if balance >= 1000 && balance < 5000{
// 		return balance * 1.621/100 + balance
// 	}
// 	return balance * 2.475/100 + balance
// }

// func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
// 	var years int
// 	for years = 0; balance < targetBalance; years++{
// 		balance = AnnualBalanceUpdate(balance)
// 	}
// 	return years
// }

// func main(){
// 	fmt.Println(InterestRate(200.75))
// 	fmt.Println(Interest(200.75))
// 	fmt.Println(AnnualBalanceUpdate(200.75))

// 	balance := 1000.000
// 	targetBalance := 1032.682765
// 	fmt.Println(YearsBeforeDesiredBalance(balance, targetBalance))
// }

// var years int
// for balance < targetBalance {
// 	balance = AnnualBalanceUpdate(balance)
// 	years += 1
// }
// return years

var myMap = map[string] int{"Black": 0, "Brown": 1, "Red": 2, "Orange": 3, "Yellow": 4, "Green": 5, "Blue": 6, "Violet": 7, "Grey": 8, "White": 9,}

func Label(colors []string) string {
	
}
	




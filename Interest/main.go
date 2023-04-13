package main

import (
    "fmt"
)

func InterestRate(balance float64) float32 {
	if balance < 0 {
		return float32(3.213)
	}else if balance < 1000{
		return float32(0.5)
	}else if balance < 5000{
		return float32(1.621)
	}
	return float32(2.475)
}

func Interest(balance float64) float64 {
	return float64(InterestRate(balance)) * balance / 100
}

func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int
    for years = 0 ; balance < targetBalance; years++ {
		balance = AnnualBalanceUpdate(balance)	
		
    }
	return years
}

func main(){
	fmt.Println(InterestRate(200.75))
	fmt.Println(Interest(200.75))
	fmt.Println(AnnualBalanceUpdate(200.75))

	balance := 1000.000
	targetBalance := 1032.682765
	fmt.Println(YearsBeforeDesiredBalance(balance, targetBalance))
}

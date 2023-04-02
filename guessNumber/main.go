//guess a number

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	
	target := rand.Intn(100) + 1
	fmt.Println(target)
	fmt.Println("guess the number")
	var number int
	

	for i:=9; i>=0; i--{
		fmt.Scanln(&number)
		
	if number == target{
		fmt.Println("que crack")
		break
	}else if number < target{
		fmt.Println("por debajo, sigue probando")
	
	}else if number > target{
		fmt.Println("por encima, sigue probando")
	}
	fmt.Printf("te quedan %v intentos\n", i)
	if i <= 0{
		fmt.Printf("vaya por dios")
	}
	}
}
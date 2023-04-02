// program that allows a student to type in their percentage grade and tells them whether they passed or not. 
//a grade of 60% or more is passing, and less than 60% is failing. 

package main

import ("fmt"
		"log")

func main(){
	fmt.Println("Type your percentage grade")
	var grade int
	fmt.Scanln(&grade)

	if grade < 60 {
		fmt.Println("You have failed :(")
	}else {
		fmt.Println("You have passed! :)")
	}

	//testing with error

	_, error := fmt.Scanln(&grade)
	if error != nil{
		log.Fatal(error)
	}
}
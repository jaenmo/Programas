package main

import "fmt"


func status (name string){
	var ok bool
	grades := map[string]float64{"Alma": 0, "Salma": 90}
	grade, ok := grades[name]
	if !ok {
		fmt.Println("Mostrooooo, nostas")
	}else if grade < 60{
		fmt.Println("invalid", ok)
	}else{
	fmt.Println("Valid", ok)
	}
}
func main (){
	status("Alma")	
	status("Salma")	
	status("Calma")	

}
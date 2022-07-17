package main

import (
	"fmt"
	"math"
)

const (dWood = 1.5
	dMarble = 2.52
	dPlastic = 0.9)

type cube struct{
	length float64
	width float64
	thickness float64 
	density float64
}
func (c cube) weight()float64{
	return c.length * c.thickness * c.width * c.density
}

type sphere struct{
	radius float64
	density float64
}
func (s sphere) weight() float64{
	return 3/4 * math.Pi * s.radius * s.radius * s.radius
}

type composed struct{
	shape []shape
}
func (co composed) weight() float64{
	shape := 
	
	return 
}

func main(){
	fmt.Println("Insert length in cm")
	var l float64
	fmt.Scanln(&l)

	fmt.Println("Insert width in cm")
	var w float64
	fmt.Scanln(&w)

	fmt.Println("Insert thikness in cm")
	var t float64
	fmt.Scanln(&t)

	fmt.Println("Choose the material")
	material := []string{"wood", "marble", "plastic"}
	for i, v := range material{
		fmt.Println(i+1, v)
	}
	fmt.Scanln(&m)
	var 
	//total := table{length: l, width: w, thickness: t, density: density {2.52, 1.5, 0.9} }

	table.volume(l, w, t)
	
	
	

}

func (t table) volume(float64, float64, float64){
	v := t.length * t.width * t.thickness
	fmt.Sprintln(v)
}

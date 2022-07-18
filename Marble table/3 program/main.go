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

type shape interface{
	weight()float64
}

type composed struct{
	shapes []shape
}

func (co composed) weight() float64{
	shapes := []shape {cube{}, sphere{}}
	for _, v := range shapes{
		co.weight() = shape.weight()
	}
	return co.weight()
}

func main(){
	
	
	fmt.Println("What is the shape of the table?")
	sliceShapes := []string {cube, sphere, composed}
	for i, v := range sliceShapes {
		fmt.Println(i+1, v)
	}
	var table int
	fmt.Scanln(&table)
	
	if table == 1 {
	parametersCube()

	chooseMaterial()
	var mat int
	fmt.Scanln(&mat)
	fmt.Println(cube.weight)

	}else if table == 2{
		
		parametersSphere()
		chooseMaterial()
		var mat int
		fmt.Scanln(&mat)
	}else if table == 3{
		
		parametersCube()
		parametersSphere()

		chooseMaterial()
		var mat int
		fmt.Scanln(&mat)
	}


}
func chooseMaterial(){
	fmt.Println("Choose the material")
	material := []string{"wood", "marble", "plastic"}
	for i, v := range material{
		fmt.Println(i+1, v)
	}		
}

func parametersCube() {
	fmt.Println("Insert length in cm")
	var l float64
	fmt.Scanln(&l)

	fmt.Println("Insert width in cm")
	var w float64
	fmt.Scanln(&w)

	fmt.Println("Insert thikness in cm")
	var t float64
	fmt.Scanln(&t)
}

func parametersSphere() {
	fmt.Println("Insert radius in cm")
	var r float64
	fmt.Scanln(&r)
}
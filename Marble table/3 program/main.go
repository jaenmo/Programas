package main

import "fmt"

type density struct{
	marble float64
	wood float64
	plastic float64
}

type table struct{
	length float64
	width float64
	thickness float64 
	density
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

	total := table{length: l, width: w, thickness: t, density: density {2.52, 1.5, 0.9} }

	//fmt.Println(total.volume(l, w, t))
	total.volume(l, w, t, m)
	//marble: 2.52, wood: 1.5, plastic: 0.9
	
	

}

func (t table) volume(float64, float64, float64, float64)float64{
	return t.length * t.width * t.thickness * density
}

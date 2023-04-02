package main

import (
	"fmt"
	"github.com/jaenmo/Programs"
	"log"
)

func main() {
	coordinates := geo.Coordinates{}
	err := coordinates.SetLatitude(37.5)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(200)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())
}

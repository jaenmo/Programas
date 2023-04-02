package main

import (
	"fmt"
	"github.com/jaenmo/Programs/datafile"
	"log"
	"sort"
)

func main(){
	listado, err := datafile.GetString("votos.txt")
	if err != nil {
		log.Fatal(err)
	}
	sort.Strings(listado)
	count := make(map[string]int)
	for _, nombres := range listado{
		count[nombres]++
		
	}
	for i, v:= range count{
		fmt.Printf("%v: %v\n",i, v)
		
	 
	
}
}
	
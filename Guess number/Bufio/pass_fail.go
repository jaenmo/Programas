package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){

fmt.Print("Enter a grade: ")
reader := bufio.NewReader(os.Stdin)
input, err := reader.ReadString('\n')
// fmt.Println(input)
if err != nil {
	log.Fatal(err)
}
input = strings.TrimSpace(input)
grade, err := strconv.ParseFloat(input, 64)
if err != nil {
	log.Fatal(err)
}

var status string
if grade >= 60{
	status = "passing"
}else {
	status = "failed" 
	}
	fmt.Println("a grade of", grade, "is", status)
}
	
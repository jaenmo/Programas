package main

import "fmt"

func main(){

	// count the items from the slice
data := []string {"a", "c", "e", "a", "e"}
counts := make (map [string] int)
for _ , v := range data {
	counts [v] ++
}

// choose the items to be counted from the slice
letters := [] string {"a", "b", "c", "d", "a"}
for _ , value := range letters{
	count, ok := counts[value]
	if !ok {
	fmt.Printf ("%s: was not found\n", value)
	}else {
	fmt.Printf ("%s: %d\n", value, count)
}
}
}
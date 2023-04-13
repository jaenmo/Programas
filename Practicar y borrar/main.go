package main

import (
	"fmt"
	"errors"
)

func CollatzConjecture(n int) (int, error) {

	for n := 0; n == 1; n++ {
		if n%2 == 0 {
			return n / 2, errors.New("no se puede dividir por 2")
		}
		return n*3 + 1, errors.New("no se puede dividir por 3")
	}
	return n, nil
}

func main() {
	fmt.Println(CollatzConjecture(12))
}

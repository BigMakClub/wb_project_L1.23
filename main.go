package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice)
	i := 3
	copy(slice[i:], slice[i+1:])
	slice[len(slice)-1] = *new(int)
	slice = slice[:len(slice)-1]
	fmt.Println(slice)

}

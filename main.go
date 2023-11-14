package main

import "fmt"

func main() {
	var week int = 2
	var slices = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	for _, val := range slices[(week-1)*7 : week*7] {
		fmt.Println(val)
	}
}

package main

import "fmt"

func main() {
	var num [5]int

	num[0] = 5

	fmt.Println(num)

	nums := [5]int{1, 2, 3, 4}

	fmt.Println(nums)

	for i := 0; i < len(num); i++ {
		fmt.Printf("%v ", num[i])
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func make_random_array(num_items, max int) []int {
	var rand_array []int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num_items; i++ {
		rand_array = append(rand_array, rand.Intn(max))
	}

	return rand_array
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func print_array(arr []int, num_items int) {
	var arraylen = len(arr)
	var items_to_print = Min(arraylen, num_items)

	for i := 0; i < items_to_print; i++ {
		fmt.Println(arr[i])
	}
}

func check_sorted_if_sorted(arr []int) bool {
	prev := 0
	for i, v := range arr {
		if i == 0 {
			prev = v
		} else {
			if prev > v {
				return false
			}
		}
	}
	return true
}

func check_sorted(arr []int) {
	if check_sorted_if_sorted(arr) {
		fmt.Println("The array is sorted")
	} else {
		fmt.Println("The array is NOT sorted!")
	}
}

func main() {
	arr := []int{10, 20, 5, 30}
	check_sorted(arr)
}

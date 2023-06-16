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
	num_elems := Min(len(arr), num_items)
	fmt.Println(arr[:num_elems])
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

func first_cut_bubblesort(arr []int) {
	arr_len := len(arr)
	for {
		swapped := false
		for i := 1; i < arr_len; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
}

func bubblesort(arr []int) {
	n := len(arr)
	for {
		swapped := false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
		n = n - 1
		if swapped == false {
			break
		}
	}
}

func main() {
	arr := make_random_array(10, 25)
	print_array(arr, len(arr)+1)
	bubblesort(arr)
	print_array(arr, len(arr)+1)
}

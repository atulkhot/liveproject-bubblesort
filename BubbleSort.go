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

func main() {
	fmt.Println(make_random_array(10, 40))
}

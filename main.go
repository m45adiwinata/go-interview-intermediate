package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := 30
	int_list := make([]int, n)
	for i := range n {
		x := rand.Intn(6)
		int_list[i] = x
	}
	fmt.Println(int_list)

	//kelompokkan tiap bilangan yang sama, contoh:
	//int_list: [1 3 2 2 5 1 4 4 4 5 3 5 5 4 0 4 4 4 2 4 2 4 4 2 1 5 2 0 3 4]
	//grouped: [
	//         [1 1 1],
	//         [3 3 3],
	//         [2 2 2 2 2 2],
	//         [5 5 5 5 5],
	//         [4 4 4 4 4 4 4 4 4 4 4],
	//         [0 0],
	// ]
	grouped := [][]int{}

	print_grouped(grouped)
}

func print_grouped(grouped [][]int) {
	fmt.Println("[")
	for _, g := range grouped {
		fmt.Printf("\t%v,\n", g)
	}
	fmt.Println("]")
}

package main

import (
	"fmt"
	"time"
)

func sumRow(row []int) int {
	time.Sleep(1 * time.Second)
	var sum int
	for _, value := range row {
		if value != 1 {
			sum += value
		}
	}
	return sum
}

func oddEven(n int) string {
	if n%2 == 0 {
		return "gn"
	}
	return "gj"
}

func generate(numRows int) [][]int {
	var res = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}
	return res
}

func main() {
	pascal := generate(15)
	for _, a := range pascal {
		for _, b := range a {
			if b != 1 {
				fmt.Printf("%d%s", b, oddEven(b))
			} else {
				fmt.Print(b)
			}
			fmt.Print(" ")
		}
		fmt.Printf(" == %d", sumRow(a))
		fmt.Println("")
	}
}

package main

import "fmt"

func main() {
	input := [][]int{
		{1, 2, 3},
		{4, 2, 1},
	}
	//fmt.Println(input)
	var sum int

	var addArray = []int{}

	for i := 0; i < len(input); i++ {
		sum = 0
		for j := 0; j < len(input[i]); j++ {

			//fmt.Println(input[i][j])
			sum = input[i][j] + sum
		}
		addArray = append(addArray, sum)
		//fmt.Println(addArray)
	}
	var largest = addArray[0]
	for k := 0; k < len(addArray); k++ {
		if largest < addArray[k] {
			largest = addArray[k]
		}

	}
	fmt.Println("largest value:", largest)
}

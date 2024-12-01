package main

import "fmt"

func funcDeepDive() {
	numbers := []int{1, 2, 3, 4}

	// passing double function utility as argument to another func
	doubleNumber := transformNumbers(&numbers, double)

	// ananomouse func
	doubleNumberAgain := transformNumbers(&numbers, func(i int) int {
		return i * 2
	})

	//variadic func
	sumUpNumber := sumup(1, 2, 3, 4, 5)     // arg with single value
	anotherSumupNumber := sumup(numbers...) // arg with array as value

	fmt.Println("Number doubled by 2 ", doubleNumber)
	fmt.Println("Number doubled by 2 ananomouse", doubleNumberAgain)
	fmt.Println("Number sum up with variadic func as single value", sumUpNumber)
	fmt.Println("Number sum up with variadic func as list args", anotherSumupNumber)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	doubleNumbers := []int{}

	for _, val := range *numbers {
		doubleNumbers = append(doubleNumbers, transform(val))
	}

	return doubleNumbers
}

func double(val int) int {
	return 2 * val
}

func sumup(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

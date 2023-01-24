package main

import (
	"fmt"
	"sort"
)

var age int = 23

func checkEven(number int) bool {
	return number%2 == 0
}

func checkList(numbers []int, checker func(int) bool) {
	values := []bool{}
	for _, value := range numbers {
		values = append(values, checker(value))
	}
	fmt.Println(numbers)
	fmt.Println(values)
}

func checkModeMean(numbers []int) (int, int, int) {
	sort.Ints(numbers)
	total := 0
	for _, v := range numbers {
		total += v
	}
	mean := total / len(numbers)
	mode := numbers[len(numbers)-1]
	return mode, mean, total

}

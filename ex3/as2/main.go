package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

func isNumberInArray(number int, arr []int) bool {
	for _, val := range arr {
		if val == number {
			return true
		}
	}
	return false
}

func numDifferentIntegers(word string) int {
	re := regexp.MustCompile(`\d+`)
	numbersStr := re.FindAllString(word, -1)

	var numbers []int
	for _, numStr := range numbersStr {
		num, err := strconv.Atoi(numStr)
		if err == nil && !isNumberInArray(num, numbers) {
			fmt.Println("sort.SearchInts(numbers, num) < len(numbers)", sort.SearchInts(numbers, num) < len(numbers))
			numbers = append(numbers, num)
		}
	}
	return len(numbers)
}

func main() {
	word := "asdasd677ca12s09912e12sd677sdfs1"

	count := numDifferentIntegers(word)
	fmt.Printf("%v", count)
}

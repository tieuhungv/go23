package util

import "strconv"

func ParseIntArray(inputs []string) []int {
	var intArray []int
	for _, val := range inputs {
		num, err := strconv.Atoi(val)
		if err == nil {
			intArray = append(intArray, num)
		}
	}
	return intArray
}

func ParseMixedArray(inputs []string) []interface{} {
	var mixedArray []interface{}
	for _, val := range inputs {
		if num, err := strconv.Atoi(val); err == nil {
			mixedArray = append(mixedArray, num)
		} else if num, err := strconv.ParseFloat(val, 64); err == nil {
			mixedArray = append(mixedArray, num)
		} else {
			mixedArray = append(mixedArray, val)
		}
	}
	return mixedArray
}

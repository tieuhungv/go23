package util

import (
	"sort"
	"strconv"
)

func SortIntArray(intArray []int) {
	sort.Ints(intArray)
}

func SortStringArray(stringArray []string) {
	sort.Strings(stringArray)
}

func SortMixedArray(arr []interface{}) {
	sort.Slice(arr, func(i, j int) bool {
		strI := toString(arr[i])
		strJ := toString(arr[j])

		if isNumeric(strI) && isNumeric(strJ) {
			numI, _ := strconv.Atoi(strI)
			numJ, _ := strconv.Atoi(strJ)
			return numI < numJ
		} else if isNumeric(strI) {
			return true
		} else if isNumeric(strJ) {
			return false
		} else {
			return strI < strJ
		}
	})
}

func toString(v interface{}) string {
	switch v.(type) {
	case int:
		return strconv.Itoa(v.(int))
	case string:
		return v.(string)
	default:
		return ""
	}
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

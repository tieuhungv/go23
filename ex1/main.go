package main

import (
	"os"
	"strings"
)

func isCountryCodeValid(code string) bool {
	validCountryCodes := []string{"VN", "US"} // List of valid country codes

	for _, countryCode := range validCountryCodes {
		if countryCode == code {
			return true
		}
	}

	return false
}

func reorderNames(data []string, size int) string {
	fistName := data[0]
	lastName := data[1]
	var countryCode = data[size-1]

	// check if the country code is not valid
	if !isCountryCodeValid(strings.ToUpper(data[size-1])) {
		return "Error: Country code not supported, only VN or US are supported."
	}

	// with middle name
	var middleName string
	if size > 3 {
		// Remove the firstName, lastName and countryCode from the slice.
		middleSlice := data[2 : size-1]
		middleName = strings.Join(middleSlice, " ")
	}

	fullName := []string{fistName, middleName, lastName}

	// Reverse the name with the country code VN
	if strings.ToUpper(countryCode) == "VN" {
		fullName = []string{lastName, middleName, fistName}
	}

	return "Output: " + strings.Join(fullName, " ")
}

func main() {
	args := os.Args[1:]
	size := len(args)
	println(reorderNames(args, size))
}

package main

import (
	"ex2/util"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	size := len(args)

	rootCmd := flag.NewFlagSet("root", flag.ExitOnError)
	isInt := rootCmd.Bool("int", false, "integer")
	isString := rootCmd.Bool("string", false, "string")
	isMix := rootCmd.Bool("mix", false, "integer and string")

	err := rootCmd.Parse(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if !*isInt && !*isString && !*isMix {
		os.Exit(1)
	}

	inputs := args[1:size]
	var output interface{}
	if *isInt {
		intArray := util.ParseIntArray(inputs)
		util.SortIntArray(intArray)
		output = intArray
	}

	if *isString {
		util.SortStringArray(inputs)
		output = inputs
	}

	if *isMix {
		mixArray := util.ParseMixedArray(inputs)
		fmt.Println("mixArray: ", mixArray)
		util.SortMixedArray(mixArray)
		output = mixArray
	}

	result := strings.Trim(fmt.Sprint(output), "[]")
	fmt.Println("Output: ", result)
}

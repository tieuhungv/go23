# Assignment for Day 2

## Goal

The goal of this assignment is to create a Go package and a command-line tool to sort input provided by the user. The program should be able to handle different types of input (integer array, float array, string array, or mixed) and output the sorted result based on the provided input type.

## Inputs

The program will accept the following inputs:

- Number (integer or float) array
- String array

## Outputs

The program will produce the following output:

- Sorted result based on the provided input type

## Package and Functions

Create a Go package with the following functions to handle sorting for different data types:

1. `SortIntegers`: Function to sort integer arrays.
2. `SortFloats`: Function to sort float arrays.
3. `SortStrings`: Function to sort string arrays.

Implement the sorting logic for each data type using appropriate algorithms.

## Command-line Tool (CLI)

Create a command-line tool to interact with the package. The CLI should:

1. Parse the input from the command line.
2. Determine the type of input (integer array, float array, string array, or mixed).
3. Utilize the corresponding sorting function from the package to sort the elements.
4. Output the sorted result.

## External Libraries

For building the CLI, you can use the following Go packages:

- [Cobra](https://github.com/spf13/cobra)
- [Gocmd](https://github.com/devfacet/gocmd)

Use the `flag` package to parse command line arguments.

## Considerations

1. Create separate functions in the package for sorting each data type to keep the code organized and modular.
2. Consider implementing generic sorting functions using interfaces to handle mixed input types efficiently.

Feel free to explore and experiment with different sorting algorithms to achieve the best performance for each data type.

**Note:** Please make sure to have the required external packages (`Cobra` and `Gocmd`) installed and imported before running the command-line tool.

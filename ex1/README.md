# Assignment for Day 1

## Reordering Names Based on Country Code

**Goal:** Create a program that reorders names based on the specified country's format.

**Inputs:** First name, last name, middle name (optional), and country code.

**Outputs:** Reordered name based on the country's format.

### Overview

The objective of this assignment is to create a program that takes in a person's name and country code as input, and then rearranges the name based on the specified country's format. The program will support the inclusion of a middle name if provided.

### Instructions

1. Parse command line arguments to extract the first name, last name, middle name (if provided), and the country code.

2. Determine the name format based on the country code.

3. Reorder the name according to the determined format.

4. Output the reordered name.

### Implementation Hints

- Use the `os` package to access command line arguments (`os.Args`).

- Handle cases where the middle name is not provided.

- Consider using a `switch` statement to determine the name format based on the country code.

### Example

Suppose we have the following input:
```shell
$ go run main.go John Smith VN

Output: Smith John 

$ go run main.go Emily Rose Watson US

Output: Emily Rose Watso

```

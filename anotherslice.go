package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var strnum string
	var intko bool
	sli := make([]int, 0)

	for strnum != "X" { // - While the user didn't enter X character
		intko = true
		for intko {
			fmt.Printf("Please enter an integer (or X to exit) : ")
			_, err := fmt.Scan(&strnum) // - Read the user input
			if err == nil {
				if strnum != "X" {
					intnum, err := strconv.Atoi(strnum)
					if err == nil {
						sli = append(sli, intnum) // - Add the new int
						sort.Ints(sli)            // - Sort the slice
						fmt.Printf("%v\n", sli)   // - Print the slice
						intko = false             // - The value field in was correct (no need to ask to the user to enter (again) an int.
					} else {
						// The user didn't enter an integer nor X charater. Then print the relevant error message
						fmt.Printf("%s \n", err)
					}
				} else {
					intko = false // Mandatory if we wanna exit the loop.
				}
			} else {
				// Should not happen because the readen value is a string but we never know :)
				fmt.Printf("%s \n", err)
			}
		}
	}
}

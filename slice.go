package main

// Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
//The program should be written as a loop.
//Before entering the loop, the program should create an empty integer slice of size (length) 3.
//During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
//The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
//The slice must grow in size to accommodate any number of integers which the user decides to enter.
//The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	intSlice := make([]int, 3)            //inits int slice length = 3
	scanner := bufio.NewScanner(os.Stdin) //makes scanner object
	fmt.Printf("enter number: ")

	// input, _ := strconv.ParseInt(scanner.Text(),10, 64) // stores the input
	// fmt.Printf("You put:%q", input) // checks to see if input working

	var index int

	for scanner.Scan() { // scans the line of input
		input := scanner.Text()                 // stores the input
		inputString, err := strconv.Atoi(input) // string to int

		if strings.EqualFold(input, "X") {
			fmt.Println("Good bye")
			break

		} else if err != nil {
			fmt.Println("re-enter input: ", err.Error())
			break

		} else {
			intSlice[index] = inputString // index = slice
			index++
			fmt.Println("index + length of slice ", index, len(intSlice))

			if len(intSlice) == index {
				intSlice = append(intSlice, index) // add slice length and capacity
				fmt.Println("appended ", intSlice, len(intSlice), cap(intSlice))
			}
			sort.Ints(intSlice) // sorts slice
			fmt.Println("slice = ", intSlice)

		}
		fmt.Println("enter # or press x to leave")

	}

}

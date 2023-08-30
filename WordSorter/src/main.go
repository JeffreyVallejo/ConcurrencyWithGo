package main

import (
	"WordSorter/src/pkg"
	"fmt"
)

func main() {
	fmt.Println("Please enter the data.")

	fmt.Print("Runtime: ")
	var runtime int
	_, err := fmt.Scan(&runtime)
	if err != nil {
		fmt.Print("There was an error for Runtime")
	}

	fmt.Print("Word Length: ")
	var wordLength int
	_, err = fmt.Scan(&wordLength)
	if err != nil {
		fmt.Println("There was an error with Word Length")
	}

	fmt.Print("Substring to sort: ")
	var substring string
	_, err = fmt.Scan(&substring)
	if err != nil {
		fmt.Println("There was an error with Substring")
	}

	fmt.Println(runtime, wordLength, substring)

	for i := 0; i < 10; i++ {
		fmt.Println(pkg.GenerateRandomString(wordLength))
	}
	return
}

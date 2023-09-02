package main

import (
	"WordSorter/src/pkg"
	"fmt"
	"strings"
	"sync"
	"time"
)

/*
Build an executable that can take command line inputs to do the following.
Have a random word generator (a..z) up to a given length (n)
Fan-Out the produced words to a number of Go-Routines (x)
The Go-Routines should filter out words that do not contain the given substring (s)
Fan-In the remaining words and output them to the console
The program should quit after a specified timeout (t)
*/

func main() {
	fmt.Println("Please enter the data.")

	fmt.Print("How many routines: ")
	var rout int
	_, err := fmt.Scan(&rout)
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

	fmt.Println(rout, wordLength, substring)

	wg := &sync.WaitGroup{}
	ch := make(chan string)
	wg.Add(rout)
	go waitToClose(wg, ch)

	start := time.Now()

	count := rout * 1000
	wordSlice := make([]string, count)
	for i := 0; i < count; i++ {
		wordSlice[i] = pkg.GenerateRandomString(wordLength)
	}

	for i := 0; i < rout; i++ {
		go eliminateInvalidStrings(ch, wg, substring, wordSlice, i, i*1000+1000)
	}

	readChannels(ch)

	wg.Wait()
	fmt.Printf("The execution time taken is %s", time.Since(start))

	fmt.Println("Now perform the same operation without go routines")
	start = time.Now()
	wordCount := 0
	for i := 0; i < count; i++ {
		if strings.Contains(wordSlice[i], substring) {
			//fmt.Println(word)
			wordCount++
			time.Sleep(10 * time.Millisecond)
		}
	}
	fmt.Println(wordCount)
	fmt.Printf("The execution time taken is %s", time.Since(start))
}

func waitToClose(wg *sync.WaitGroup, ch chan<- string) {
	wg.Wait()
	close(ch)
}

// FanOut
func eliminateInvalidStrings(out chan string, wg *sync.WaitGroup, s string, words []string, start int, end int) {
	for i := start; i < end; i++ {
		if strings.Contains(words[i], s) {
			out <- words[i]
		}
	}
	wg.Done()
}

// FanIn
func readChannels(in <-chan string) {
	wordCount := 0
	for _ = range in {
		//fmt.Println(word)
		wordCount++
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println(wordCount)
}

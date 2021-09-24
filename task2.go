// 2- Search and count the word available in file.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// initiate file-handle to read from
	fileHandle, err := os.Open("../tasks.txt")

	// check if file-handle was initiated correctly
	if err != nil {
		panic(err)
	}

	// make sure to close file-handle upon return
	defer fileHandle.Close()

	// initiate scanner from file handle
	fileScanner := bufio.NewScanner(fileHandle)

	// tell the scanner to split by words
	fileScanner.Split(bufio.ScanWords)

	// initiate counter
	count := 0

	// for looping through results
	for fileScanner.Scan() {
		//fmt.Printf("word: '%s' - position: '%d'\n", fileScanner.Text(), count)
		count++
	}

	// check if there was an error while reading words from file
	if err := fileScanner.Err(); err != nil {
		panic(err)
	}

	// print total word count
	fmt.Printf("total word count: '%d'", count)
}

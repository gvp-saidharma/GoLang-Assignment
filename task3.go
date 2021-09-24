// 3- Replace the word into file and count.
package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func visit(path string, fi os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil //
	}

	read, err := ioutil.ReadFile("../data.txt")
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(read))
	//fmt.Println(path)

	newContents := strings.Replace(string(read), "pyton", "python", -1)

	//fmt.Println(newContents)

	err = ioutil.WriteFile("../out.txt", []byte(newContents), 0)
	if err != nil {
		panic(err)
	}

	return nil
}

func main() {
	err := filepath.Walk(".", visit)
	if err != nil {
		panic(err)
	}
}

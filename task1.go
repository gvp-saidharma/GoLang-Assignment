// 1-Create employee data excel file with duplicate record  that contain the fallowing field
// Emp id, Emp Name, Emp Salary
// Write a python code to display duplicate employee record and count no of duplicate record available into a employee data file.

package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("../employee test2.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

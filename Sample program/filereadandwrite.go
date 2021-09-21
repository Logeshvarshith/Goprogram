package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	f.SetCellValue("Sheet1", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	f.SetCellValue("Sheet2", "A1", 200)
	f.SetActiveSheet(index)
	//f.SetActiveSheet(sheet1)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("sample.xlsx"); err != nil {
		fmt.Println(err)
	}
	fmt.Println("File creation completed")
	openFile()
}
func openFile() {
	f, err := excelize.OpenFile("sample.xlsx")

	if err != nil {
		log.Fatal(err)
	}

	c1, err := f.GetCellValue("Sheet1", "A2")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sheet1-A2:", c1)

	c2, err := f.GetCellValue("Sheet1", "B2")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sheet1-B2:", c2)

	c3, err := f.GetCellValue("Sheet2", "A1")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sheet2-A1:", c3)
}

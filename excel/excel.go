package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func NewXLSXFile() *excelize.File {
	return excelize.NewFile()
}

func WriteDataIntoXLSX(f *excelize.File, sheetTitle string, header *[]interface{}, datas []*[]interface{}) {
	_, err := f.SearchSheet(sheetTitle, "")
	if err != nil {
		f.NewSheet(sheetTitle)
		f.SetSheetRow(sheetTitle, fmt.Sprintf("A1"), header)
		for index, data := range datas {
			f.SetSheetRow(sheetTitle, fmt.Sprintf("A%d", index+2), data)
		}
	} else {
		rows, err := f.GetRows(sheetTitle)
		if err != nil {
			fmt.Println(err)
			return
		}
		numberOfRows := len(rows)
		for index, data := range datas {
			f.SetSheetRow(sheetTitle, fmt.Sprintf("A%d", index+1+numberOfRows), data)
		}
	}
	// style, _ := f.NewStyle(`{"number_format":14}`)
	// f.SetColStyle("Sheet1", "F", style)
}

func SaveXLSX(f *excelize.File, filePath, fileName string) string {
	path := fmt.Sprintf("%s/%s", filePath, fileName)
	if err := f.SaveAs(path); err != nil {
		fmt.Println(err)
	}
	return path
}

package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strconv"
	"strings"
)

func closeExcelFiles() {
	if !eu.WriteExcelFiles {
		return
	}
	if nezapisuj() {
		return
	}
	MyMkdDirIfNotExist()
	for i := 0; i < len(xlsxs); i++ {
		 xlsxs[i].SaveAs( fmt.Sprintf(fmt.Sprint("./files/%v%v%.",digitToSee,"f%v"),name[i],
			"_with_step_", eu.Epsilon,".xlsx"))
	}

}

func MyMkdDirIfNotExist() {
	if _, err := os.Stat("./files"); os.IsNotExist(err) {
		 os.Mkdir("./files", 777)
	}
}

func nezapisuj() bool {
	if orbit==nil  {
		return true
	}
	if len(orbit)==0 {
		return true
	}
	if orbit[0]==nil {
		return true
	}
	if len(orbit[0])<5  {
		return true
	}
	return false
}

func writeToExcel1(time, positionX, positionY, velocityX,velocityY, accelerationX, accelerationY float64,row, numberFile int) {
	if !eu.WriteExcelFiles {
		return
	}
	 xlsxs[numberFile].SetCellValue("Sheet1", "A"+strconv.Itoa(row), time)
	 xlsxs[numberFile].SetCellValue("Sheet1", "B"+strconv.Itoa(row), positionX)
	 xlsxs[numberFile].SetCellValue("Sheet1", "C"+strconv.Itoa(row), positionY)
	 xlsxs[numberFile].SetCellValue("Sheet1", "D"+strconv.Itoa(row), velocityX)
	 xlsxs[numberFile].SetCellValue("Sheet1", "E"+strconv.Itoa(row), velocityY)
	 xlsxs[numberFile].SetCellValue("Sheet1", "F"+strconv.Itoa(row), accelerationX)
	 xlsxs[numberFile].SetCellValue("Sheet1", "G"+strconv.Itoa(row), accelerationY)
}



func setExcelTitle(titles string, file int) {
	if !eu.WriteExcelFiles {
		return
	}
	abc := strings.Split("A B C D E F G H I J K L", " ")
	titlesArray := strings.Split(titles, " ")
	num := strconv.Itoa(1)
	for i := 0; i < len(titlesArray); i++ {
		cell := abc[i]+num
		 xlsxs[file].SetCellValue("Sheet1", cell, titlesArray[i])
	}
}

func makeExcelFile() []*excelize.File {
	result := make([]*excelize.File,0,numberObject)
	for i := 0; i < numberObject; i++ {
		result = append(result, excelize.NewFile())
	}
	return result
}


func writeToExcel2_1(row int, time, positionX, accelerationX, positionY, accelerationY, velocityX, velocityY, r float64, numberFile int) {
	if !eu.WriteExcelFiles {
		return
	}
	 xlsxs[numberFile].SetCellValue("Sheet1", "A"+strconv.Itoa(row), time)
	 xlsxs[numberFile].SetCellValue("Sheet1", "B"+strconv.Itoa(row), positionX)
	 xlsxs[numberFile].SetCellValue("Sheet1", "C"+strconv.Itoa(row), positionY)
	 xlsxs[numberFile].SetCellValue("Sheet1", "D"+strconv.Itoa(row), velocityX)
	 xlsxs[numberFile].SetCellValue("Sheet1", "E"+strconv.Itoa(row), velocityY)
	 xlsxs[numberFile].SetCellValue("Sheet1", "F"+strconv.Itoa(row), accelerationX)
	 xlsxs[numberFile].SetCellValue("Sheet1", "G"+strconv.Itoa(row), accelerationY)
}

func writeToExcel2_2(row int, time , velocityX, velocityY float64, numberFile int) {
	if !eu.WriteExcelFiles {
		return
	}
	 xlsxs[numberFile].SetCellValue("Sheet1", "A"+strconv.Itoa(row), time)
	 xlsxs[numberFile].SetCellValue("Sheet1", "D"+strconv.Itoa(row), velocityX)
	 xlsxs[numberFile].SetCellValue("Sheet1", "E"+strconv.Itoa(row), velocityY)
}

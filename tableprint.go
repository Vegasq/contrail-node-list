package main

import "fmt"

func getMaxColWidth(table [][]string) []int {
	if len(table) == 0 {
		return nil
	}
	columnCount := len(table[0])
	var result []int
	var colW int

	for i:=0; i<columnCount; i+=1 {
		result = append(result, 0)
	}

	for _, row := range table {
		for colI, col := range row {
			colW = len(col)
			if colW > result[colI] {
				result[colI] = colW
			}
		}
	}

	return result
}

func fillTableValues(table [][]string, sizes []int){
	var diffW int
	var colW int
	var tailEnd string

	for rowI, row := range table {
		for colI, col := range row {
			colW = len(col)
			diffW = sizes[colI] - colW

			tailEnd = ""
			for i:=0; i!=diffW; i++ {
				tailEnd += "_"
			}
			table[rowI][colI] = table[rowI][colI] + tailEnd
		}
	}


	var rowStr string
	for _, row := range table {
		rowStr = "| "
		for _, col := range row {
			rowStr = rowStr + col + " | "
		}
		fmt.Println(rowStr)
	}
}

func PrintTable(table [][]string) {
	maxColWidth := getMaxColWidth(table)
	fillTableValues(table, maxColWidth)
}
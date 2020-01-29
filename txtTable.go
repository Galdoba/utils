package utils

import (
	"strconv"
	"strings"
)

//TxtTable - текстовая таблица
type TxtTable struct {
	cells []string
	rows  int
	cols  int
}

//NewTxtTable - создает объект из файла. Содержит кол-во строк и колонок и текст всех ячеек
func NewTxtTable(path string) *TxtTable {
	if !tableValid(path) {
		return nil
	}
	tt := &TxtTable{}
	r, c := tableDimentions(path)
	tt.rows = r
	tt.cols = c
	for i := 1; i <= r; i++ {
		line := LinesFromTXT(path)[i]
		data := strings.Split(line, "	")
		for j := range data {
			tt.cells = append(tt.cells, data[j])
		}
	}
	return tt
}

//Rows - возвращает кол-во строк
func (table *TxtTable) Rows() int {
	return table.rows
}

//Cols - возвращает кол-во колонок
func (table *TxtTable) Cols() int {
	return table.cols
}

//AllCells - возвращает весь массив данных в виде []string
func (table *TxtTable) AllCells() []string {
	return table.cells
}

//Cell - возвращет текст ячейки с заданными координатами
func (table *TxtTable) Cell(col, row int) string {
	id := table.cols*(row) + col
	if col > table.cols-1 || col < 0 {
		return "ERROR: TABLE OUT OF BOUNDS Cell(" + strconv.Itoa(col) + ":" + strconv.Itoa(row) + ") not found!"
	}
	if row > table.rows-1 || row < 0 {
		return "ERROR: TABLE OUT OF BOUNDS Cell(" + strconv.Itoa(col) + ":" + strconv.Itoa(row) + ") not found!"
	}
	return table.cells[id]
}

func tableValid(table string) bool {
	lines := LinesFromTXT(table)
	head := lines[0]
	tail := lines[len(lines)-1]
	if !strings.Contains(head, "txtTableHead") {
		return false
	}
	if !strings.Contains(tail, "txtTableTail") {
		return false
	}
	return true
}

func tableDimentions(table string) (rows, cols int) {
	if !tableValid(table) {
		panic(0)
	}
	head := LinesFromTXT(table)[0]
	data := strings.Split(head, " = ")
	r := ParseValueInt(data[1], "rows")
	c := ParseValueInt(data[2], "cols")
	return r, c
}

//ParseValueInt - возвращает числовое значение X из строки "dataType X"
func ParseValueInt(dataStr, dataType string) int {
	str := strings.TrimLeft(dataStr, dataType)
	str = strings.Trim(str, " ")
	val, _ := strconv.Atoi(str)
	return val
}

func getCell(line string, n int) string {

	data := strings.Split(line, "	")
	return data[n-1]
}

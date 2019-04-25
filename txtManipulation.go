package utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//LinesFromTXT - открывает txt и возвращает построчно всё его содержимое
func LinesFromTXT(path string) []string {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

//FileNames - возвращает имена файлов содержащих marker
//(dir = "./" для текущей директории)
func FileNames(dir, marker string) []string {
	var names []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if strings.Contains(f.Name(), marker) {
			names = append(names, f.Name())
		}
	}
	return names

}

//EditLineInFile - заменяет строку номер n в файле поадресу file на newContent
func EditLineInFile(file string, n int, newContent string) {
	lines := LinesFromTXT(file)
	lines[n] = newContent
	output := strings.Join(lines, "\n")
	err := ioutil.WriteFile(file, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

//InFileContains - Возвращает номер строки в котором содержится content
func InFileContains(file, content string) int {
	lines := LinesFromTXT(file)
	for i := range lines {
		if strings.Contains(lines[i], content) {
			return i
		}
	}
	return -1
}

//InFileContainsN - Возвращает слайс с номерами строк в которых содержится content
func InFileContainsN(file, content string) []int {
	lines := LinesFromTXT(file)
	var numbers []int
	for i := range lines {
		if strings.Contains(lines[i], content) {
			numbers = append(numbers, i)
		}
	}
	return numbers
}

package utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	asciiBlack   = "\u001b[30;1m"
	asciiRed     = "\u001b[31;1m"
	asciiGreen   = "\u001b[32;1m"
	asciiYellow  = "\u001b[33;1m"
	asciiBlue    = "\u001b[34;1m"
	asciiMagenta = "\u001b[35;1m"
	asciiCyan    = "\u001b[36;1m"
	asciiWhite   = "\u001b[37;1m"
)

//ASCIIColor - возвращает string покрашенный в 1 из RGB/CMYK цветов для терминала. (регистро не чувствительно)
//если цвет не определен - возвращает только string
func ASCIIColor(col string, text string) string {
	col = strings.ToUpper(col)
	switch col {
	case "RED":
		return asciiRed + text + asciiWhite
	case "GREEN":
		return asciiGreen + text + asciiWhite
	case "YELLOW":
		return asciiYellow + text + asciiWhite
	case "BLUE":
		return asciiBlue + text + asciiWhite
	case "MAGENTA":
		return asciiMagenta + text + asciiWhite
	case "CYAN":
		return asciiCyan + text + asciiWhite
	case "BLACK":
		return asciiBlack + text + asciiWhite
	case "WHITE":
		return asciiWhite + text + asciiWhite

	}
	return text
}

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

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

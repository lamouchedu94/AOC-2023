package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data, err := read("input")
	if err != nil {
		fmt.Println(err)
		return
	}

	for x, line := range data {
		for y, car := range line {
			if isNumber(car) {
				fmt.Println(x, y)
				checkIsOk(data, x, y)
			}
		}
	}
}

func isNumber(val rune) bool {
	return 48 <= int(val) && int(val) <= 57
}

func checkIsOk(data []string, x int, y int) {
	var runeVal rune
	fmt.Printf(string(data[x-1][y-1]), "%+q", &runeVal)
	if data[x+1][y+1] != '.' && !isNumber(runeVal) {
		fmt.Println(data)
	}
}

func read(path string) ([]string, error) {
	var tab []string
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	test := bufio.NewScanner(file)
	test.Split(bufio.ScanLines)

	for test.Scan() {
		tab = append(tab, test.Text())
	}
	return tab, err

}

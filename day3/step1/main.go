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
	res := ""
	for x, line := range data {
		grp := false
		for y, car := range line {
			if isNumber(car) {
				// fmt.Println(string(data[x][y]), x, y)
				res += string(data[x][y])
				grp = true
				checkIsOk(data, x, y)
			} else {
				if res != "" {
					fmt.Println(res)
				}
				res = ""
				grp = false
			}
		}
		_ = grp
	}
}

func isNumber(val rune) bool {
	return 48 <= int(val) && int(val) <= 57
}

func checkIsOk(data []string, x int, y int) {

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

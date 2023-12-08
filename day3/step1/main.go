package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := read("input.old")
	if err != nil {
		fmt.Println(err)
		return
	}
	tabRes := []string{}
	res := ""
	for y, line := range data {
		grp := false
		start := 0
		end := 0
		for x, car := range line {
			if isNumber(rune(car)) {
				if res == "" {
					start = x
				}
				res += string(data[y][x])
				grp = true
			} else {
				if res != "" {
					end = x
					fmt.Println(res, checkIsOk(data, start, end, y))
					//fmt.Println(start, end)
					if checkIsOk(data, start, end, y) {
						tabRes = append(tabRes, res)
					}
					start = 0
					end = 0
				}
				res = ""
				grp = false
			}
		}
		if grp {
			if checkIsOk(data, start, end, y) {
				tabRes = append(tabRes, res)
			}
		}
	}
	sum := 0
	for _, val := range tabRes {
		temp, _ := strconv.Atoi(val)
		sum += temp
	}
	fmt.Println(sum)
}

func isNumber(val rune) bool {
	return 48 <= val && val <= 57
}

func checkIsOk(data []string, start int, end int, y int) bool {
	start = start - 1
	if end < start {
		end = len(data[0]) - 1
	} else {
		end = end + 1
	}

	if start < 0 {
		start += 1
	}
	if end > len(data[0]) {
		end -= 1
	}
	ligne := y - 1
	lastLine := y + 1
	if ligne < 0 {
		ligne = y
	}
	if lastLine >= len(data) {
		lastLine = y
	}
	valid := false
	for l := ligne; l <= lastLine; l++ {
		for k := start; k < end; k++ {
			//fmt.Println(string(data[l][k]))
			if string(data[l][k]) != "." && !isNumber(rune(data[l][k])) {
				//fmt.Println(string(data[l][k]))
				return true
			}
		}
	}
	return valid

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

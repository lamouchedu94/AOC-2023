package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := read("input")
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
					//fmt.Println(res)
					end = x
					//fmt.Println(start, end)
					if checkIsOk(data, start, end, y) {
						if !inTab(tabRes, res) {
							tabRes = append(tabRes, res)

						}
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
				if !inTab(tabRes, res) {
					tabRes = append(tabRes, res)
				}
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

func inTab(tab []string, val string) bool {
	in := false
	for _, value := range tab {
		if value == val {
			return true
		}
	}
	return in

}

func isNumber(val rune) bool {
	return 48 <= val && val <= 57
}

func checkIsOk(data []string, start int, end int, y int) bool {
	start = start - 1
	end = end + 1

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
			if string(data[l][k]) != "." && !isNumber(rune(data[l][k])) {
				//fmt.Println(string(data[l][k]))
				valid = true
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

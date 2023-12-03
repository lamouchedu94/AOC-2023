package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	tab, err := read("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var total int
	m := make(map[string]string)
	m["one"], m["two"], m["three"], m["four"], m["five"], m["six"], m["seven"], m["eight"], m["nine"] = "1", "2", "3", "4", "5", "6", "7", "8", "9"
	for _, line := range tab {
		var first, last string
		for _, car := range line {
			if 48 <= int(car) && int(car) <= 57 {
				if first == "" {
					first = string(car)
				} else {
					last = string(car)
				}
			}
		}
		if last == "" {
			last = first
		}
		temp, _ := strconv.Atoi(first + last)
		total += temp
	}
	fmt.Println(total)
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

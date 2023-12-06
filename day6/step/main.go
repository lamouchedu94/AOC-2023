package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

func main() {
	deb := time.Now()
	data, err := read("../input")
	if err != nil {
		fmt.Println(err)
		return
	}
	// test := []string{}

	r, _ := regexp.Compile(`\d+`)
	var a int
	fmt.Sscanf(r.FindAllString(data[0], -1)[0], "%d", &a)
	tabTime := r.FindAllString(data[0], -1)
	tabDistance := r.FindAllString(data[1], -1)
	res := 1
	for i := 0; i < len(tabTime); i++ {
		var time, distance int
		fmt.Sscanf(tabTime[i], "%d", &time)
		fmt.Sscanf(tabDistance[i], "%d", &distance)
		res *= calc(time, distance)
	}
	fmt.Println(res)
	fin := time.Now()
	fmt.Println(fin.Sub(deb))

}

func calc(time int, distance int) int {
	res := 0
	for i := 1; i <= time; i++ {
		// calculer temps restant et multiplier par i
		timeleft := time - i
		if timeleft*i > distance {
			res += 1
		}
	}
	return res
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

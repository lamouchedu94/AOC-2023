package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type result struct {
	id       int
	is_valid bool
}

func main() {
	tab, err := read("../input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	all_party := []result{}
	for _, line := range tab {
		var party result
		party.is_valid = true
		new_line := strings.Split(line, ":")
		fmt.Sscanf(new_line[0], "Game %d", &party.id)
		data := strings.Split(new_line[1], ";")
		for _, value := range data {
			game := strings.Split(string(value), ",")
			for _, car := range game {
				var number int
				var color string
				fmt.Sscanf(car, "%d %s", &number, &color)
				switch color {
				case "blue":
					if number > 14 {
						party.is_valid = false
					}
				case "red":
					if number > 12 {
						party.is_valid = false
					}
				case "green":
					if number > 13 {
						party.is_valid = false
					}
				}
			}
		}
		all_party = append(all_party, party)
	}
	var total int
	for _, val := range all_party {
		if val.is_valid {
			fmt.Printf("%+v\n", val)
			total += val.id
		}
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

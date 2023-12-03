package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func parse(text string, row int) {
	number := ""
	tempCords := [3]int{}
	isFound := false
	for ind, val := range text {
		if unicode.IsDigit(val) && !isFound {
			isFound = true
			number += string(val)
			tempCords[0] = ind
			tempCords[2] = row
		} else if unicode.IsDigit(val) {
			number += string(val)
		} else if !unicode.IsDigit(val) && isFound {
			tempCords[1] = ind - 1
			num, _ := strconv.Atoi(number)
			intMap[tempCords] = num
			number = ""
		}
	}
}

func first(text string) int {
	return 0
}

// get all number start/end index and row number in something like map
// [start, end, row]: number
// then get all sumbols cords [[row1, column1], [row2, column2]]
// and then try to get key by looking around symbol
var (
	intMap     = make(map[[3]int]int)
	symbolList = [][2]int{}
)

func main() {
	data, err := os.Open("3/test_data")
	// data, err := os.Open("3/data")
	if err != nil {
		fmt.Println("errir")
	}
	defer data.Close()

	sc := bufio.NewScanner(data)
	i := 0
	for sc.Scan() {
		t := sc.Text()
		parse(t, i)

		i++

	}
	fmt.Println(intMap)
}

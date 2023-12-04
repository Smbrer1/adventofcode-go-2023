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
			isFound = false
		}
		if val != '.' && !unicode.IsDigit(val) {
			symbolList = append(symbolList, [2]int{row, ind})
		}
	}
}

func first() int {
	ans := 0
	for _, i := range symbolList {
		for k := -1; k < 2; k++ {
			for j := -3; j < 3; j++ {
				for h := -3; h < 3; h++ {
					cords := [3]int{i[1] + h, i[1] + j, i[0] + k}
					elem, ok := intMap[cords]
					fmt.Println(cords)
					if ok {
						delete(intMap, cords)
						fmt.Println(elem)
						ans += elem
					}
				}
			}
		}
	}
	return ans
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
	ans := 0
	for sc.Scan() {
		t := sc.Text()
		parse(t, i)

		i++

		ans += first()
	}
	fmt.Println(intMap)
	fmt.Println(symbolList)
	fmt.Println(ans)
}

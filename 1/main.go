package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func firstAndLast(text string, fDigit chan<- string, lDigit chan<- string) {
	isFound := false
	fdigit := ""
	ldigit := ""
	fmt.Println(text)
	for _, digit := range text {
		if unicode.IsDigit(digit) && !isFound {
			fdigit = string(digit)
			isFound = true
		} else if unicode.IsDigit(digit) && isFound {
			ldigit = string(digit)
		}
	}
	if ldigit == "" {
		ldigit = fdigit
	}
	fDigit <- fdigit
	lDigit <- ldigit
}

func main() {
	data, err := os.Open("./data")
	if err != nil {
		fmt.Println("errir")
	}
	defer data.Close()

	sc := bufio.NewScanner(data)
	fdChan := make(chan string)
	ldChan := make(chan string)
	ans := 0
	for sc.Scan() {
		t := sc.Text()
		go firstAndLast(t, fdChan, ldChan)
		num1 := <-fdChan
		num2 := <-ldChan
		temp, _ := strconv.Atoi(num1 + num2)
		fmt.Println(temp)
		ans += temp
	}
	fmt.Println(ans)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/Smbrer1/adventofcode-go-2023/utils"
)

func first(text string, fDigit chan<- string, lDigit chan<- string) {
	isFound := false
	fdigit := ""
	ldigit := ""
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

func numInString(str string, tokens map[string]string) string {
	for i := 0; i < len(str); i++ {
		for j := 3; j < 6; j++ {
			num := int(str[i]) - 48
			if num < 10 && num > 0 {
				fmt.Printf("num: %d\n", num)
				return fmt.Sprint(num)
			}
			if i <= len(str)-j {
				if num, ok := tokens[str[i:i+j]]; ok {
					return num
				}
			}
		}
	}
	return ""
}

func second(text string, nums chan<- int) {
	tokens := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	rTokens := map[string]string{
		"eno":   "1",
		"owt":   "2",
		"eerht": "3",
		"ruof":  "4",
		"evif":  "5",
		"xis":   "6",
		"neves": "7",
		"thgie": "8",
		"enin":  "9",
	}
	rText := utils.Reverse(text)
	num1 := numInString(text, tokens)
	num2 := numInString(rText, rTokens)
	fmt.Printf("secod num1: %s\n", num1)
	fmt.Printf("secod num2: %s\n", num2)
	ans, _ := strconv.Atoi(num1 + num2)

	fmt.Printf("secod ans: %d\n", ans)
	nums <- ans
}

func main() {
	data, err := os.Open("1/data")
	if err != nil {
		fmt.Println("errir")
	}
	defer data.Close()

	sc := bufio.NewScanner(data)
	fdChan := make(chan string)
	ldChan := make(chan string)
	ansChan := make(chan int)
	ans := 0
	ans2 := 0
	for sc.Scan() {
		t := sc.Text()

		fmt.Println(t)

		go first(t, fdChan, ldChan)
		num1 := <-fdChan
		num2 := <-ldChan
		temp, _ := strconv.Atoi(num1 + num2)
		ans += temp

		go second(t, ansChan)
		ans2 += <-ansChan
	}
	fmt.Println("=======")
	fmt.Println(ans)
	fmt.Println("=======")
	fmt.Println(ans2)
}

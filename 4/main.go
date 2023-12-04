package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func fisrt(text string) int {
	card := strings.Split(text, ":")
	nums := strings.Split(card[1], "|")
	nums[0] = strings.Trim(nums[0], " ")
	winnigNums := strings.Split(nums[0], " ")
	fmt.Println(winnigNums)
	nums[1] = strings.Trim(nums[1], " ")
	myNums := strings.Split(nums[1], " ")
	fmt.Println(myNums)
	points := 0
	win := false
	for _, i := range myNums {
		if i == "" {
			continue
		}
		if slices.Contains(winnigNums, i) {
			if !win {
				win = true
				points = 1
				fmt.Printf("winnig number: %v current points: %d\n", i, points)
			} else {
				points *= 2
				fmt.Printf("winnig number: (%v) current points: %d\n", i, points)
			}
		}
	}
	return points
}

type Card struct {
	number        int
	winnigNumbers []string
	myNumbers     []string
	count         int
}

func second(cards []Card) int {
	ans := 0
	for i, card := range cards {
		winnigCount := 0
		for _, num := range card.myNumbers {
			if num == "" {
				continue
			}
			if slices.Contains(card.winnigNumbers, num) {
				winnigCount += 1
				cards[i+winnigCount].count += 1 * card.count
			}
		}
		ans += card.count
	}
	for i, card := range cards {
		fmt.Println(i+1, "card count:", card.count)
	}
	return ans
}

func parse(text string) Card {
	card := strings.Split(text, ":")
	nums := strings.Split(card[1], "|")
	nums[0] = strings.Trim(nums[0], " ")
	winnigNums := strings.Split(nums[0], " ")
	nums[1] = strings.Trim(nums[1], " ")
	myNums := strings.Split(nums[1], " ")
	cardNumber := strings.Split(card[0], " ")[1]

	val, _ := strconv.Atoi(cardNumber)
	return Card{
		number:        val,
		winnigNumbers: winnigNums,
		myNumbers:     myNums,
		count:         1,
	}
}

func main() {
	data, err := os.Open("4/data")
	if err != nil {
		fmt.Println("errir")
	}
	defer data.Close()

	sc := bufio.NewScanner(data)
	ans := 0
	var cards []Card
	for sc.Scan() {
		t := sc.Text()
		// ans += fisrt(t)
		cards = append(cards, parse(t))
	}
	fmt.Println(ans)
	fmt.Println(second(cards))
}

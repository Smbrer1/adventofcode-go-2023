package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var allowedGames = map[string]int{
	"red": 12, "green": 13, "blue": 14,
}

func first(text string) int {
	arr := strings.Split(text, ":")
	id, _ := strconv.Atoi(strings.Split(arr[0], " ")[1])
	games := strings.Split(arr[1], ";")

	for _, game := range games {
		dices := strings.Split(game, ",")

		for _, dice := range dices {
			dice = strings.Trim(dice, " ")
			diceArr := strings.Split(dice, " ")
			// fmt.Printf("dice: %s\n", dice)

			intAmount, _ := strconv.Atoi(diceArr[0])

			if allowedGames[diceArr[1]] < intAmount {
				// fmt.Printf("Game: %d is not allowed in this color: %s\n\n", id, diceArr[1])
				return 0
			}
		}

	}
	return id
}

func second(text string) int {
	arr := strings.Split(text, ":")
	games := strings.Split(arr[1], ";")
	ans := 1
	minGames := map[string]int{
		"red": 0, "green": 0, "blue": 0,
	}
	for _, game := range games {
		dices := strings.Split(game, ",")

		for _, dice := range dices {
			dice = strings.Trim(dice, " ")
			diceArr := strings.Split(dice, " ")
			// fmt.Printf("dice: %s\n", dice)

			intAmount, _ := strconv.Atoi(diceArr[0])

			if minGames[diceArr[1]] < intAmount {
				minGames[diceArr[1]] = intAmount
			}
		}

	}
	fmt.Printf("minGames: %v\n", minGames)
	for _, val := range minGames {
		ans *= val
	}

	return ans
}

func main() {
	data, err := os.Open("2/data")
	if err != nil {
		fmt.Println("errir")
	}
	defer data.Close()

	sc := bufio.NewScanner(data)
	ans := 0
	ans2 := 0
	for sc.Scan() {
		t := sc.Text()
		fmt.Println(t)
		ans += first(t)
		ans2 += second(t)

	}
	fmt.Println("======")
	fmt.Println(ans)
	fmt.Println("======")
	fmt.Println(ans2)
}

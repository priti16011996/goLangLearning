package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("switch case ")
	rand.Seed(time.Now().Unix())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Dice Vlaue is 1 and you can Open with it")
	case 2:
		fmt.Println("Dice Vlaue is 2 and you can move 2 step ahead")
	case 3:
		fmt.Println("Dice Vlaue is 3 and you can move 3 step ahead")
	case 4:
		fmt.Println("Dice Vlaue is 4 and you can move 4 step ahead")
	case 5:
		fmt.Println("Dice Vlaue is 5 and you can move 5 step ahead")
	case 6:
		fmt.Println("Dice Vlaue is 6 and you can move 6 step ahead and get one more chance to through dice")
	default:
		fmt.Println("invalid input. Please retry")
	}
}

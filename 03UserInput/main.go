package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Take User Input from User")
	welcome := "Welcome Coders.."
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate us based on teaching standard:")

	//comma ok // err
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us", input)
	fmt.Printf("type of rating %T", input)

}

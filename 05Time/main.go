package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time study")
	predefinedTm := time.Now()
	fmt.Println(predefinedTm)
	fmt.Println(predefinedTm.Format("01-02-2006 15:30:04 Monday"))
}

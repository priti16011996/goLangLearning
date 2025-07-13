package main

import "fmt"

func main() {
	fmt.Println("Loop, Continue, Break, Goto Syntax Understanding")

	days := []string{"Mon", "Tue", "Wed", "Thus", "Fri", "Sat", "Sun"}

	fmt.Println("days array contains:", days)

	fmt.Println("1 st way to create loop")
	for j := 0; j < len(days); j++ {
		fmt.Println(days[j])
	}

	fmt.Println("2 nd way to create loop")
	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("3 rd way to create loop")

	for index, value := range days {
		fmt.Printf("value for index %v, value is %v\n", index, value)
	}

	b := 10

	for b < 25 {
		if b == 13 {
			b++
			continue
		}
		if b == 20 {
			b++
			goto lco
		}
		fmt.Println(b)
		b++
	}

lco:
	fmt.Println("using go to jump directly to this")
}

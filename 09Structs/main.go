package main

import "fmt"

func main() {
	fmt.Println("Struct basics")
	// there is no concept of inheritance and super ,parents
	priti := User{"Priti maurya", "Software Engineer", 4, "Venera Technologies"}
	fmt.Println(priti)
	fmt.Printf("priti structdetails %v\n", priti)
	fmt.Printf("priti structdetails %+v\n", priti)
	fmt.Printf("Name  %v and Designation %v\n", priti.Name, priti.Designation)
	priti.getName()
	priti.updateCompanyName("facebook")
}

type User struct {
	Name        string
	Designation string
	Experiance  int
	CompanyName string
}

func (u User) getName() {
	fmt.Println("Username is:", u.Name)
}

func (u User) updateCompanyName(newCompanyName string) {
	u.CompanyName = newCompanyName
	fmt.Println(u.CompanyName)
}

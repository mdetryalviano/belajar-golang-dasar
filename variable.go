package main

import "fmt"

func main() {
	var name string

	name = "Eko"
	fmt.Println(name)

	var country = "Indonesia"
	fmt.Println(country)

	city := "Semarang"
	fmt.Println(city)

	var (
		fullName string = "M. Detryalviano Maharandi"
		age      int    = 17
		province        = "Central Java"
	)

	fmt.Println("Name = ", fullName)
	fmt.Println("Age = ", age)
	fmt.Println("Province = ", province)

}

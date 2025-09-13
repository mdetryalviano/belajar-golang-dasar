package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Eko"
	//person["address"] = "Subang"

	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}

	fmt.Println(len(person))
	fmt.Println(person["name"])

	person["name"] = "Joko"
	fmt.Println(person["name"])

}

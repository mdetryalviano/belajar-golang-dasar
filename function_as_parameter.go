package main

import (
	"fmt"
	"slices"
	"strings"
)

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	badWords := []string{"anjing", "bangsat", "kontol", "memek", "cuki", "tolol", "bego", "tai", "ngentod"}

	if slices.Contains(badWords, strings.ToLower(name)) {
		return "***"
	} else {
		return name
	}

}

func main() {
	sayHelloWithFilter("Eko", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
	sayHelloWithFilter("BANGSAT", spamFilter)
}

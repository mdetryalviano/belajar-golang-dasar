package main

import (
	"fmt"
	"slices"
	"strings"
)

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blacklisted", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		blacklistedUser := []string{"eko", "kurniawan", "khannedy", "joko", "budi", "nugraha"}
		if slices.Contains(blacklistedUser, strings.ToLower(name)) {
			return true
		} else {
			return false
		}
	}

	registerUser("kurniawan", blacklist)
	registerUser("sigma", blacklist)

	registerUser("khannedy", func(name string) bool {
		blacklistedUser := []string{"eko", "kurniawan", "khannedy", "joko", "budi", "nugraha"}
		if slices.Contains(blacklistedUser, strings.ToLower(name)) {
			return true
		} else {
			return false
		}
	})

}

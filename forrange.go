package main

import (
	"fmt"
)

func main() {
	names := []string{"eko", "kurniawan", "khannedy"}
	for index, name := range names {
		fmt.Println("Indeks", index, "Name =", name)
	}

	for _, name := range names {
		fmt.Println("Name =", name)
	}

	for i := 0; i <= len(names)-1; i++ {
		fmt.Println("Name =", names[i])
	}

}

package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Eko"
	names[1] = "Khannedy"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [2]int{
		12,
		22,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])

}

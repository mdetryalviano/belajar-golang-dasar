package main

import "fmt"

func main() {
	type NoKTP string

	var ktpEko NoKTP = "111111111"
	fmt.Println(ktpEko)
	fmt.Println(NoKTP("222222222"))
}

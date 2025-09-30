package main

import "fmt"

type User struct {
	Id, Age              int
	Name, Email, Address string
}

func main() {
	var user User
	user.Id = 1
	user.Age = 18
	user.Name = "Eko"
	user.Email = "eko@mail.com"
	user.Address = "Jl. Eko"

	fmt.Println(user)

	var joko User
	joko.Id = 1
	joko.Age = 18
	joko.Name = "Joko"
	joko.Email = "joko@mail.com"
	joko.Address = "Jl. Eko"
	fmt.Println(joko)

	budi := User{1, 18, "Budi", "budbudi@mail.com", "Jl. Eko"}
	fmt.Println(budi)
}

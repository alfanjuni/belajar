package main

import "fmt"

func main() {
	//cara deklarasi variable
	var name string
	var age = 30
	var nama = "ALFAN JUNI"
	name = "KODINGSQUAD"

	//tidak perlu var
	alamat := "Pecangaan"

	var (
		firstName = "ahmad"
		lastName  = "alfan"
	)

	fmt.Println(name)

	name = "ALFAN"
	fmt.Println(name)
	fmt.Println(nama)
	fmt.Println(age)
	fmt.Println(alamat)
	fmt.Println(firstName, lastName)

}

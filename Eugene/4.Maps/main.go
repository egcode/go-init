package main

import "fmt"

func main() {
	// var nums map[string]string

	// nums := make(map[string]string)

	nums := map[string]string{
		"one": "odin",
		"two": "dva",
	}
	nums["three"] = "tri"
	fmt.Println(nums)

	dict := make(map[int]string)
	dict[8] = "vosem"
	dict[10] = "desyat"
	fmt.Println(dict)

	// // Delete key with value
	// delete(nums, "one")
	// fmt.Println(nums)

	fmt.Println("---------------------------")

	printMap(nums)
}

func printMap(c map[string]string) {
	for english, russian := range c {
		fmt.Println("Engilsh:", english, " Russian:", russian)
	}
}

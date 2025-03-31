package main

import "fmt"

/*
Write program to return if count frrom 1-3 is correct
1. "1","2","3" entered return true
2. "one", "two", "three" entered return true
3. "first", "second", "third" entered return true
4. all other cases return false
*/
func main() {
	var str1, str2, str3 string

	fmt.Print("Enter three strings: ")
	fmt.Scan(&str1, &str2, &str3)
	switch {
	case str1 == "1" && str2 == "2" && str3 == "3":
		fmt.Println(true)
	case str1 == "one" && str2 == "two" && str3 == "three":
		fmt.Println(true)
	case str1 == "first" && str2 == "second" && str3 == "third":
		fmt.Println(true)
	default:
		fmt.Println(false)
	}
}

package main

import (
	"fmt"
	s "strings"
)

func main() {
	var b string
	fmt.Println("Enter a string")
	fmt.Scan(&b)
	a := palindrome(b)
	if a == "Y" {
		fmt.Println("is a palindrome")
	}
	if a == "N" {
		fmt.Println("is not a palindrome")
	}
}
func palindrome(a string) string {
	b := []byte{}
	for i := len(a) - 1; i >= 0; i-- {
		b = append(b, a[i])
	}
	c := string(b)
	fmt.Println(c)
	d := s.ToLower(a)
	e := s.ToLower(c)
	if d == e {
		return "Y"
	}

	return "N"

}

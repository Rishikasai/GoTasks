package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func username() string {
	var c string
	fmt.Println("Enter the username (for sample: username=ram, password: sita)")
	fmt.Scan(&c)
	return c
}
func password() string {
	var b string
	fmt.Println("Enter the password")
	fmt.Scan(&b)
	return b
}

func withdraw(a string, b int) string {

	i1, _ := strconv.Atoi(a)
	if b > i1 {
		return "N"
	}
	i1 = i1 - b
	i2 := strconv.Itoa(i1)
	return i2

}

func deposit(a string, b int) string {

	i1, _ := strconv.Atoi(a)
	i1 = i1 + b
	i2 := strconv.Itoa(i1)
	return i2
}
func main() {
	csvfile, err := os.Open("bank.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	c := username()
	b := password()

	for {
		record, err := r.Read()
		if err == io.EOF {
			fmt.Println("Invalid username or password pls try again")
			break
		}
		if c == record[0] && b == record[1] {
			b := 1
			fmt.Println("Welcome ", record[0])
			for b > 0 {
				fmt.Println("**********************************")
				fmt.Println(` 
		Enter 1 to check balance
		Enter 2 to withdraw Amount
		Enter 3 to deposit Amount
		Enter 4 to exit
							`)
				fmt.Println("**********************************")
				var a int
				fmt.Scan(&a)
				if a == 1 {
					fmt.Println("The balance amount is", record[2])
				}
				if a == 2 {
					fmt.Println("enter the amount to withdraw")
					var b int
					fmt.Scan(&b)
					d := withdraw(record[2], b)
					if d == "N" {
						fmt.Println("Insufficient balance")
					}
					if d != "N" {
						record[2] = d
						fmt.Println("Sucessfully withdrawn ", b, "Rs.  The remaining balance is", record[2])
					}
				}
				if a == 3 {
					fmt.Println("enter the amount to deposit")
					var b int
					fmt.Scan(&b)
					record[2] = deposit(record[2], b)
					fmt.Println("Sucessfully deposited", b, "Rs. The remaining balance is", record[2])
				}
				if a == 4 {
					b = 0
				}
			}
			break
		}

	}
}

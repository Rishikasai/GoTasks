package main

import "fmt"

type data struct {
	password string
	amount   int
}

func main() {

	users := make(map[string]*data)
	users["rishika"] = &data{password: "rishi", amount: 1000}
	users["sita"] = &data{password: "ram", amount: 2000}
	users["ram"] = &data{password: "sita", amount: 2500}
	users["bharath"] = &data{password: "india", amount: 3000}
	fmt.Println("===================================")
	fmt.Println(`Welcome to sample bank!
Enter the username and password:`)
	fmt.Println("===================================")
	fmt.Println("username:")
	var username string
	fmt.Scan(&username)
	if users[username] == nil {
		fmt.Println("invalid username pls try again")
	}
	if users[username] != nil {
		var pwd string
		fmt.Println("password:")
		fmt.Scan(&pwd)
		if users[username].password != pwd {
			fmt.Println("wrong password pls try again")
		}
		if users[username].password == pwd {
			b := 1
			for b > 0 {
				fmt.Println("Welcome ", username)
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
					fmt.Println("The balance amount is", users[username].amount)
				}
				if a == 2 {
					fmt.Println("enter the amount to withdraw")
					var b int
					fmt.Scan(&b)
					users[username].amount = users[username].amount - b
					fmt.Println("Sucessfully withdrawn ", b, "Rs.  The remaining balance is", users[username].amount)
				}
				if a == 3 {
					fmt.Println("enter the amount to deposit")
					var b int
					fmt.Scan(&b)
					users[username].amount = users[username].amount + b
					fmt.Println("Sucessfully deposited", b, "Rs. The remaining balance is", users[username].amount)
				}
				if a == 4 {
					b = 0
				}
			}
		}
	}
}

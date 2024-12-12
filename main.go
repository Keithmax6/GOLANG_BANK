package main

import "fmt"
import "main/fileops"

const accountBalanceFileName = "balanceTxtFile.txt"


func main() {
	var accountBalance, err =       fileops.ReadFromFile(accountBalanceFileName,1000)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}
	fmt.Println("Welcome to Go Bank!")
	for {
		com()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			} else {
				accountBalance += depositAmount
				fmt.Println("Your new balance is: ", accountBalance)
				fileops.WriteBalanceToFile(accountBalance,accountBalanceFileName)
			}
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds")
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Your new balance is: ", accountBalance)
				fileops.WriteBalanceToFile(accountBalance,accountBalanceFileName)
			}
		case 4:
			fmt.Println("Thank you for using Go Bank!")
		default:
			fmt.Println("Invalid choice")

		}

	}

}

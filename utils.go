package main

import "fmt"

func GetName() string {
	name := ""

	fmt.Println("Welcome to Pari's Casino....")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s , let's Play!\n", name)
	return name
}

func GetBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Enter your bet (balance = $%d): ", balance)
		fmt.Scan(&bet) // only int input for now
		if bet > balance {
			fmt.Println("Bet cannot be larger then balance.")
		} else {
			break
		}
	}
	return bet
}
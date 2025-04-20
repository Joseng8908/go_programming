package mywithdraw

import "fmt"

func Withdraw(balance int, amount int) (int, error) {
	if balance - amount < 0{
		return 0, fmt.Errorf("balance is %d, amount is %d, invalid input", balance, amount)
	}
	return (balance - amount), nil
} 
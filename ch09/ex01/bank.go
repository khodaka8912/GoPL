// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraw = make(chan transaction)

type transaction struct {
	amount int
	result chan<- bool
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	result := make(chan bool)
	withdraw <- transaction{amount, result}
	return <-result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case trans := <-withdraw:
			if balance < trans.amount {
				trans.result <- false
			} else {
				balance -= trans.amount
				trans.result <- true
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-

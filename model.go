package main

import "time"

type Account struct {
	Account_ID      uint64
	Document_Number uint64
}

type OperationType struct {
	OperationType_ID uint64
	Description      []string
}

type Transaction struct {
	Transaction_ID   uint64
	Account_ID       uint64
	OperationType_ID uint64
	Amount           int64
	EventDate        time.Time
}

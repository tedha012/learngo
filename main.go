package main

import (
	"fmt"

	"github.com/tedha012/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}

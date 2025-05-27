package main

import (
	"cs50-go/utils"
	"fmt"
)

func main() {
	name := utils.GetString("What's your name? ")
	fmt.Printf("hello, %v\n", name)
}

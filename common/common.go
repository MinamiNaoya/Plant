package common

import (
	"fmt"
	"log"
)

func inputString(prompt string) string {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}
	return input
}

func inputInt(prompt string) int {
	var input int
	fmt.Print(prompt)
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}
	return input
}

package common

import (
	"fmt"
	"log"
)

func InputString(prompt string) string {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}
	return input
}

func InputInt(prompt string) int {
	var input int
	fmt.Print(prompt)
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}
	return input
}

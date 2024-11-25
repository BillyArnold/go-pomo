package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the pomodoro cli application")

	inputVal, err := inputString("Enter your name")
	if err != nil {
		fmt.Println("there was an input error")
	}

	input, err := inputString("Enter another word")
	if err != nil {
		fmt.Println("there was an input error")
	}

	fmt.Println(inputVal, input)

}

func inputString(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%v \n", prompt)

	val, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	val = strings.TrimSpace(val)

	return val, nil
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to the pomodoro cli application")

	pomoSettings := collectSettings()

	cycle(pomoSettings.minutesPerPomo)

	// inputVal, err := inputString("Enter your name")
	// if err != nil {
	// 	fmt.Println("there was an input error")
	// }

	// input, err := inputString("Enter another word")
	// if err != nil {
	// 	fmt.Println("there was an input error")
	// }

	// fmt.Println(inputVal, input)

}

type SettingsObject struct {
	tasks          []string
	cycles         int
	minutesPerPomo int
}

func collectSettings() *SettingsObject {
	cyclesInput, err := inputString("How many cycles of focused work would you like?")
	if err != nil {
		fmt.Println("there was an input error", err)
	}

	cyclesInputInt, err := strconv.Atoi(cyclesInput)
	if err != nil {
		fmt.Println("A number was not entered")
	}

	minutesInput, err := inputString("How many minutes per cycle should there be?")
	if err != nil {
		fmt.Println("there was an input error", err)
	}

	minutesInputInt, err := strconv.Atoi(minutesInput)
	if err != nil {
		fmt.Println("A number was not entered")
	}

	return &SettingsObject{
		tasks:          []string{"task 1", "task 2"},
		cycles:         cyclesInputInt,
		minutesPerPomo: minutesInputInt,
	}
}

func cycle(minutes int) {
	numberOfSeconds := minutes * 60

	for i := 1; i <= numberOfSeconds; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(i)
	}
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

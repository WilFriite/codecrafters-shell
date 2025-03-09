package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input: ", err)
			os.Exit(1)
		}

		// List of supported commands so far
		supportedCommands := []string{"echo", "type", "exit"}

		// Display the message once received
		length := len(command) - 1 // Remove the newline character
		if command[:length] == "exit 0" {
			os.Exit(0)
		}
		commandToArr := strings.Split(command[:length], " ")
		mainCommand := commandToArr[0]
		args := commandToArr[1:]

		if strings.Contains(mainCommand, "echo") {
			text := strings.Join(args, " ")
			if text == "" {
				fmt.Println("Error: you must provide a string to echo")
			} else {
				fmt.Println(strings.TrimSpace(text))
			}
		} else if strings.Contains(mainCommand, "type") {
			text := args[0]
			if text == "" {
				fmt.Println("Error: you must provide a command")
			} else {
				if slices.Contains(supportedCommands, text) {
					fmt.Println(text + " is a supported command")
				} else {
					fmt.Println(text + ": not found")
				}
			}
		} else {
			fmt.Println(mainCommand + ": command not found")
		}
	}
}

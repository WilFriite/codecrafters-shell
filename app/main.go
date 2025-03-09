package main

import (
	"bufio"
	"fmt"
	"os"
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

		// Display the message once received
		length := len(command) - 1 // Remove the newline character
		if command[:length] == "exit 0" {
			os.Exit(0)
		}
		if strings.Contains(command[:length], "echo") {
			text := command[:length][4:]
			if text == "" {
				fmt.Println("Error: you must provide a string to echo")
			}
			fmt.Println(strings.TrimSpace(text))
		} else {
			fmt.Println(command[:length] + ": command not found")
		}
	}
}

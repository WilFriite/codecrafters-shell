package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

		// Display the message once received
		length := len(command) - 1 // Remove the newline character
		if command[:length] == "exit 0" {
			os.Exit(0)
		}
		commandToArr := strings.Split(command[:length], " ")
		mainCommand := commandToArr[0]
		args := commandToArr[1:]

		if strings.Contains(mainCommand, "echo") {
			echoText(args)
			continue
		}
		if strings.Contains(mainCommand, "type") {
			text := args[0]
			if text == "" {
				fmt.Println("Error: you must provide a command")
				continue
			}

			if isShellBuiltin(text) {
				fmt.Println(text + " is a shell builtin")
			} else if path, err := exec.LookPath(text); err == nil {
				fmt.Println(text + " is " + path)
			} else {
				fmt.Println(text + ": not found")
			}
		} else {
			cmd := exec.Command(mainCommand, args...)
			err := cmd.Start()
			if err != nil {
				fmt.Println(mainCommand + ": command not found")
			}
		}
		//else {
		//	fmt.Println(mainCommand + ": command not found")
		//}
	}
}

func isShellBuiltin(text string) bool {
	// List of supported commands so far
	supportedCommands := []string{"echo", "type", "exit"}
	return slices.Contains(supportedCommands, text)
}

func echoText(text []string) string {
	formatted := strings.Join(text, " ")
	if formatted == "" {
		fmt.Println("Error: you must provide a string to echo")
	} else {
		fmt.Println(strings.TrimSpace(formatted))
	}
	return formatted
}

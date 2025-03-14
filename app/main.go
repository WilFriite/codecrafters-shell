package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for {
		// Uncomment this block to pass the first stage
		_, _ = fmt.Fprint(os.Stdout, "$ ")

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

		// mainCommand, args := ParseCommand(command[:length])
		mainCommand, args, err := ParseShellWords(command[:length])

		if err != nil {
			fmt.Println("Error parsing command: ", err)
			continue
		}

		switch mainCommand {
		case "echo":
			EchoCommand(args)
			break
		case "pwd":
			PwdCommand()
			break
		case "cd":
			CdCommand(args)
		case "type":
			TypeCommand(args)
			break
		default:
			cmd := exec.Command(mainCommand, args...)
			stdout, err := cmd.Output()
			if err != nil {
				fmt.Println(mainCommand + ": command not found")
			}
			fmt.Print(string(stdout))
			break
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

		switch mainCommand {
		case "echo":
			EchoCommand(args)
			break
		case "type":
			TypeCommand(args)
			break
		default:
			cmd := exec.Command(mainCommand, args...)
			stdout, err := cmd.Output()
			if err != nil {
				fmt.Println(mainCommand + ": command not found")
			}
			fmt.Println(string(stdout))
			break
		}
	}
}

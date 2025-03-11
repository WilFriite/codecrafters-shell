package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func isShellBuiltin(text string) bool {
	// List of supported commands so far
	supportedCommands := []string{"echo", "type", "exit"}
	return slices.Contains(supportedCommands, text)
}

func EchoCommand(args []string) string {
	formatted := strings.Join(args, " ")
	if formatted == "" {
		fmt.Println("Error: you must provide a string to echo")
	} else {
		fmt.Println(strings.TrimSpace(formatted))
	}
	return formatted
}

func TypeCommand(args []string) {
	if len(args) != 1 {
		fmt.Println("Error: you must provide only one argument")
	}
	text := args[0]
	if text == "" {
		fmt.Println("Error: you must provide a command")
		return
	}

	if isShellBuiltin(text) {
		fmt.Println(text + " is a shell builtin")
	} else if path, err := exec.LookPath(text); err == nil {
		fmt.Println(text + " is " + path)
	} else {
		fmt.Println(text + ": not found")
	}
}

func PwdCommand() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
}

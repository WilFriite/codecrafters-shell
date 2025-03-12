package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func hasBalancedQuotes(s string) bool {
	inQuote := false
	escaped := false

	for _, r := range s {
		if escaped {
			escaped = false
			continue
		}

		if r == '\\' {
			escaped = true
			continue
		}

		if r == '\'' {
			inQuote = !inQuote
		}
	}

	return !inQuote
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func removeQuotes(s string) string {
	var result strings.Builder

	for _, r := range s {
		if r != '\'' {
			result.WriteRune(r)
		}
	}

	return result.String()
}

func IsSupported(text string) bool {
	return slices.Contains(SupportedCommands, text)
}

func EchoCommand(args []string) string {
	trimmed := deleteEmpty(args)
	formatted := strings.Join(trimmed, " ")

	if formatted == "" {
		fmt.Println("Error: you must provide a string to echo")
	} else {
		processed := removeQuotes(formatted)
		fmt.Println(strings.TrimSpace(processed))
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

	if IsSupported(text) {
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

func CdCommand(args []string) {
	dir := args[0]
	var targetDir string
	if dir != "~" {
		targetDir = dir
	} else {
		targetDir = os.Getenv("HOME")
	}
	err := os.Chdir(targetDir)
	if err != nil {
		fmt.Println(dir + ": No such file or directory")
	}
}

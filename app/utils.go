package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
)

func IsSupported(text string) bool {
	return slices.Contains(SupportedCommands, text)
}

func EchoCommand(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stdout)
		return
	}

	for i := 0; i < len(args)-1; i++ {
		fmt.Fprintf(os.Stdout, "%s ", args[i])
	}

	fmt.Fprintln(os.Stdout, args[len(args)-1])
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

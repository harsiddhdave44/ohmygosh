package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	// path := os.Getenv("PATH")
	// fmt.Println("PATH:", path)
	currentPath, err := os.Getwd()
	for {
		if err != nil {
			fmt.Println("There was an error:", err)
			return
		}

		fmt.Print("gosh ", currentPath, "> ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.TrimSpace(command)

		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			return
		}

		command_split := strings.Fields(command)

		if len(command_split) == 0 {
			continue
		}

		main_command := command_split[0]
		args := command_split[1:]

		switch main_command {
		case "exit":
			return
		case "ls":
			listCurrentDirectories()
		case "cd":
			currentPath = changeDirectory(args)
		case "pwd":
			printWorkingDirectory()
		case "":
			continue
		default:
			fmt.Println("You entered:", command)
			continue
		}
	}
}

func printWorkingDirectory() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("There was an error:", err)
		return
	}
	fmt.Println(currentDirectory)
}

func listCurrentDirectories() {
	currentDirectories, err := os.ReadDir(".")

	if err != nil {
		fmt.Println("There was an error:", err)
		return
	}

	for _, dir := range currentDirectories {
		fmt.Println(dir.Name())
	}
}

func changeDirectory(path []string) string {
	err := os.Chdir(path[0])
	if err != nil {
		fmt.Println("There was an error:", err)
		return ""
	}
	currentDirectory, _ := os.Getwd()

	return currentDirectory
}

func clearTerminal() bool {
	currentOS := runtime.GOOS

	if currentOS == "windows" {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
	}
	return true
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

//Checks whether a command prefixed 'type' is a builtin command or not(unrecognizable)
func isShellBuiltin(command string ) bool {
	builtins := []string{"echo", "exit", "type"}
	if slices.Contains(builtins, command) {
		return true
	} else {
		return false
	}

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		// TODO: Uncomment the code below to pass the first stage
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
		}

		command = strings.TrimSpace(command)

		//builtin exit command
		if command == "exit" {
			break

		//builtin echo command
		}else if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])

		//outputs the type of the command 
		}else if strings.HasPrefix(command, "type ") {
			if isShellBuiltin(command[5:]) {
				fmt.Println(command[5:] + " is a shell builtin")
			} else {
				fmt.Println(command[5:] + ": not found")
			}
		}else {
			fmt.Println(command + ": command not found")
		}
		// command = strings.TrimRight(command, "\r\n")
		// fmt.Println(command + ": command not found")

	}

}

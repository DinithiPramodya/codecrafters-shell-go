package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

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
	if command == "exit" {
		break
	}

	fmt.Println(command[:len(command) - 1] + ": command not found")
	// command = strings.TrimRight(command, "\r\n")
	// fmt.Println(command + ": command not found")

	}

}

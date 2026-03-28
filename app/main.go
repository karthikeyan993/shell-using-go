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

	for {
		fmt.Print("$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("read error: %v\n", err)
		}
		command = strings.TrimSpace(command)
		if command == "exit" {
			os.Exit(0)
		}
		fmt.Printf("%s: command not found\n", command)
	}
}

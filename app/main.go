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
	fmt.Print("$ ")
	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	command = strings.TrimSpace(command)
	fmt.Printf("%s: command not found\n", command)
}

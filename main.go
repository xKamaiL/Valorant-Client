package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

const REGION = "au"

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username : ")
	username, _ := reader.ReadString('\n')
	fmt.Print("Enter Password : ")
	password, _ := terminal.ReadPassword(0)

	fmt.Printf("%s %s\n", username, password)
}

package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println(("Run encrypt or decrypt command with help flag for more information"))
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("file encryption")
	fmt.Println("Simple file encrypter for your day-to-day needs.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tgo run . encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands: ")
	fmt.Println("")
	fmt.Println("\t encrypt\tEncrypts a file given a password")
	fmt.Println("\t decrypt\tTries to decrypt a file using a password")
	fmt.Println("\t help\t\tDisplays help text")
	fmt.Println("")
}

func encryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a file to encrypt")
		os.Exit(0)
	}

	file := os.Args[2]

	if !validateFile(file) {
		panic("File does not exist")
	}

	password := getPassword()
	fmt.Println("Encrypting file...")
}

func decryptHandle() {

}

func getPassword() []byte {
	fmt.Println("Please enter a password")
	password , _ : =term.ReadPassword(0)

	fmt.Println("Please confirm your password")
	confirmPassword , _ : =term.ReadPassword(0)

	if !validatePassword(password, confirmPassword) {
		fmt.Println("\nPasswords do not match")
		return getPassword()
	}

	return password
}

func validatePassword(password :[]byte,confirmPassword: [byte]) bool {
	if(!bytes.Equal(password,confirmPassword)){
		return false
	}
	return true
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

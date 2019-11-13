package userinput

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetEncryptionKey will ask the user for the encryption key
func GetEncryptionKey(isEncrypt bool) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the encryption password: ")
	first, _ := reader.ReadString('\n')
	if !isEncrypt {
		return strings.TrimSuffix(first, "\n")
	}

	fmt.Print("Confirm Password: ")
	second, _ := reader.ReadString('\n')
	if first != second {
		fmt.Println("Encryption keys don't match")
		fmt.Println("Terminating")
		os.Exit(3)
	}
	return strings.TrimSuffix(first, "\n")
}

// ceaser cipher tool
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// caesarCipher decrypts a string using a caesar cipher and the given
// rotation
func caesarCipher(input string, rotation byte) string {
	output := make([]byte, utf8.RuneCountInString(input))
	for idx, c := range input {
		output[idx] = byte(c) - rotation
	}
	return string(output)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ceasar Cipher decryption tool")
	for {
		fmt.Print("Enter phrase (:quit to exit): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == ":quit" {
			return
		}
		fmt.Print("Enter rotation: ")
		rot, _ := reader.ReadString('\n')
		rot = strings.TrimSpace(rot)
		rotation, err := strconv.ParseInt(rot, 10, 64)
		if err != nil {
			fmt.Printf("error - %s", err)
			continue
		}
		decryptedStr := caesarCipher(text, byte(rotation))
		fmt.Printf("Decrypted string:\n%s\n", decryptedStr)
		fmt.Println("\n=============================================\n")
	}
}

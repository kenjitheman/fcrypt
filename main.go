package main

import (
	"fmt"

	"main.go/core"
)

func main() {
	key := []byte("examplekeyexampl") // 16, 24, or 32 bytes for AES-128, AES-192, AES-256

	inputFile := "./man.txt"
	encryptedFile := "./encrypted.enc"
	decryptedFile := "./decrypted.txt"

	err := core.EncryptFile(key, inputFile, encryptedFile)
	if err != nil {
		fmt.Println("[ERROR] encryption error:", err)
		return
	}

	fmt.Println("[SUCCESS] file encrypted successfully")

	err = core.DecryptFile(key, encryptedFile, decryptedFile)
	if err != nil {
		fmt.Println("[ERROR] decryption error:", err)
		return
	}

	fmt.Println("[SUCCESS] file decrypted successfully")
}

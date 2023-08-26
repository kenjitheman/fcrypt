package fcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

func EncryptFile(key []byte, inputFile, outputFile string) error {
	plaintext, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = output.Write(iv)
	if err != nil {
		return err
	}

	_, err = output.Write(ciphertext)
	return err
}

func DecryptFile(key []byte, inputFile, outputFile string) error {
	ciphertext, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = output.Write(plaintext)
	return err
}

<h2 align="center">fcrypt - file encryption/decryption golang library</h2>

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###

## project structure:

```rust
├── fcrypt.go
├── go.mod
├── LICENSE
└── README.md
```

## installation

```shell
go get github.com/kenjitheman/fcrypt
```

## usage

```go
package main

import (
	"fmt"

	"github.com/kenjitheman/fcrypt"
)

func main() {
	key := []byte("examplekeyexampl") // 16, 24, or 32 bytes for AES-128, AES-192, AES-256

	inputFile := "./man.txt"
	encryptedFile := "./encrypted.enc"
	decryptedFile := "./decrypted.txt"

	err := fcrypt.EncryptFile(key, inputFile, encryptedFile)
	if err != nil {
		fmt.Println("[ERROR] encryption error:", err)
		return
	}

	fmt.Println("[SUCCESS] file encrypted successfully")

	err = fcrypt.DecryptFile(key, encryptedFile, decryptedFile)
	if err != nil {
		fmt.Println("[ERROR] decryption error:", err)
		return
	}

	fmt.Println("[SUCCESS] file decrypted successfully")
}
```

## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)

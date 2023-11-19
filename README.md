## fcrypt - file encryption/decryption tool

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###

## Project structure:

```rust
fcrypt
│
├── fcrypt.go
├── go.mod
├── LICENSE
└── README.md
```

## Installation

```sh
go get github.com/kenjitheman/fcrypt
```

## Usage

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

## Contributing

- Pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change.

## License

- [MIT](./LICENSE)

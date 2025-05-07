package main

import (
	"fmt"

	"github.com/thuylinh996ll/cryptit/decrypt"
	"github.com/thuylinh996ll/cryptit/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("Hehe!")
	fmt.Println("Encrypted String: ", encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}

package main

import (
	"crypto"
	"encoding/hex"
	"fmt"

	_ "crypto/md5"
	_ "embed"
)

//go:embed input.txt
var secret_key string

func main() {
	fmt.Println("Day 4")

	hash := crypto.MD5.New()

	var hes string = "....."
	var i int

	for i = 0; hes[:5] != "00000"; i++ {
		hash.Reset()
		hash.Write([]byte(fmt.Sprintf("%s%d", secret_key, i)))
		hes = hex.EncodeToString(hash.Sum(nil))
	}

	fmt.Println(i-1, hes) // 117946

	for i = 0; hes[:6] != "000000"; i++ {
		hash.Reset()
		hash.Write([]byte(fmt.Sprintf("%s%d", secret_key, i)))
		hes = hex.EncodeToString(hash.Sum(nil))
	}

	fmt.Println(i-1, hes)

}

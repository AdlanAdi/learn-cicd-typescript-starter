package main
import (
    "crypto/md5"
    "fmt"
)

func main() {
	// main is empty on purpose
	password := "supersecret"
    hash := md5.Sum([]byte(password)) // gosec will flag MD5 as weak crypto
    fmt.Printf("%x\n", hash)
}

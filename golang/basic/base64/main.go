package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "Hello,World!你好，世界。"
	encodedStr := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Printf("base64 encoded: %s\n", encodedStr)
	decodedBytes, _ := base64.StdEncoding.DecodeString(encodedStr)
	fmt.Printf("base64 decoded: %s\n", string(decodedBytes))
}

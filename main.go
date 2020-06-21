package main

import (
	"os"
	"fmt"
	"encoding/base64"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments passed. Example usage:", os.Args[0], "string_to_encode")
		os.Exit(1)
	}
	encodeString(os.Args[1])
}

func encodeString(s string) {
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write([]byte(s))
	encoder.Close()
	fmt.Print("\n")
}
package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("generate key")
	fmt.Println("encryption")

	_ = os.WriteFile("test", []byte("write"), 0644)

}

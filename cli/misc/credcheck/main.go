package main

import (
	"fmt"

	"github.com/profburke/bgurt/cli/utilities"
)

func main() {
	c := utilities.SetCredentials()
	fmt.Printf("creds: %v\n", c)

	fmt.Println("A-OK")
}

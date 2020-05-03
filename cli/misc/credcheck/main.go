package main

import (
	"fmt"

	"github.com/profburke/bgg/cli/utilities"
)

func main() {
	c := utilities.SetCredentials()
	fmt.Printf("creds: %v\n", c)

	fmt.Println("A-OK")
}

package main

import (
	"fmt"
	"os"
)

// Handle errors.
func checkErrors(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	err := runCommand()
	checkErrors(err)
}

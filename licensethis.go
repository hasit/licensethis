// Package main adds an OSS license of choice to the current directory
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 0 {
		switch args[0] {
		case "help":
			fmt.Println("help")
		case "generate":
			fmt.Println("generate")
		case "list":
			fmt.Println("list")
		case "config":
			fmt.Println("config")
		default:
			fmt.Println("help")
		}
	} else {
		fmt.Println("help")
	}
}

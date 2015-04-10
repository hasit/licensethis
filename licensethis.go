// Package main adds an OSS license of choice to the current directory
package main

import "fmt"

// Usage prints the proper usage of licensethis from godoc.org
func Usage() {
	fmt.Println("Usage of licensethis:")
	return
}

// ParseFlags parses command-line flags, checks for valid flags and stores the values in a global structure for further use
func ParseFlags() error {
	// Parse flags here
	var e error
	return e
}

func main() {
	if ParseFlags() != nil {
		Usage()
		return
	}
}

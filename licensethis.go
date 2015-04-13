package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

//UserInfo stores information of user
type UserInfo struct {
	First string
	Last  string
}

//Check checks for errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//printHelp opens project page (http://hasit.github.io/licensethis/) in the default browser.
func printHelp() {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", "http://hasit.github.io/licensethis/").Start()
		check(err)
	case "windows", "darwin":
		err = exec.Command("open", "http://hasit.github.io/licensethis/").Start()
		check(err)
	default:
		err = fmt.Errorf("unsupported platform")
		check(err)
	}
}

func userConfig() {
	first := os.Getenv("FIRSTNAME")
	last := os.Getenv("LASTNAME")

	//if first != nil; last != nil {
	fmt.Printf("First: %v\n", first)
	fmt.Printf("Last: %v\n", last)
	//}
}

//parseArgs parses command line arguments and calls appropriate functions.
func parseArgs(args []string) {
	if len(args) != 0 {
		switch args[0] {
		case "help":
			printHelp()
		case "config":
			userConfig()
		case "info":
			fmt.Println("info")
		case "list":
			fmt.Println("list")
		case "generate":
			fmt.Println("generate")
		default:
			fmt.Println("Incorrect usage!")
			fmt.Println("Type `licensethis help` for help on proper usage.")
		}
	} else {
		fmt.Println("Incorrect usage!")
		fmt.Println("Type `licensethis help` for help on proper usage.")
	}
}

func main() {
	args := os.Args[1:]

	parseArgs(args)
}

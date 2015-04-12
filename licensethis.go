package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

//printHelp opens project page (http://hasit.github.io/licensethis/) in the default browser.
func printHelp() {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", "http://hasit.github.io/licensethis/").Start()
		if err != nil {
			panic(err)
		}
	case "windows", "darwin":
		err = exec.Command("open", "http://hasit.github.io/licensethis/").Start()
		if err != nil {
			panic(err)
		}
	default:
		err = fmt.Errorf("unsupported platform")
		if err != nil {
			panic(err)
		}
	}
}

//parseArgs parses command line arguments and calls appropriate functions.
func parseArgs(args []string) {
	if len(args) != 0 {
		switch args[0] {
		case "help":
			printHelp()
		case "generate":
			fmt.Println("generate")
		case "list":
			fmt.Println("list")
		case "info":
			fmt.Println("info")
		case "config":
			fmt.Println("config")
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

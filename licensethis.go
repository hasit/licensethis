package main

import (
	"fmt"
	"os"
)

func printHelp() {

}

func parseArgs(args []string) {
	if len(args) != 0 {
		switch args[0] {
		case "help":
			fmt.Println("help")
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

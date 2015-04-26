package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//check() checks and panics if an error has occured
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//help() prints usage, commands and examples for licensethis.
func help() {
	helptext := `licensethis - Choose an OSS license to your project with ease.

Usage:
licensethis help
licensethis list
licensethis info <license-name>
licensethis generate <license-name>

Commands:
help		Show this help text.
list		Show a list of all available OSS licenses.
info		Show more information for a license.
generate	Generate LICENSE.txt file in current folder after asking for author name.

Examples:
licensethis help		Show this help text.
licensethis info mit		Show more information for MIT license.
licensethis generate mit	Geneate MIT license text file in current directory.\n`

	fmt.Printf("%v", helptext)
}

func info(licensename []string) {
	//fmt.Printf("%v\n", licensename[0])

	file, e := ioutil.ReadFile("licenses.json")
	if e != nil {
		log.Fatal(e)
	}
	var jsontype map[string][]map[string]interface{}
	err := json.Unmarshal(file, &jsontype)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Results: %v \n", jsontype["licenses"])
	for i := range jsontype["licenses"] {
		item := jsontype["licenses"][i]
		//fmt.Printf("%v\n",item["name"])
		if item["name"] == licensename[0] {
			fmt.Printf("%v\n%v\n%v", item["name"], item["version"], item["text"])
		}
	}
}

//parseArgs parses command line arguments and calls appropriate functions.
func parseArgs(args []string) {
	if len(args) != 0 {
		switch args[0] {
		case "help":
			help()
		case "info":
			info(args[1:])
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

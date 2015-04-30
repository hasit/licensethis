package main

import (
	"fmt"
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
	helptext := `licensethis - Choose an OSS license for your project with ease.

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
licensethis generate mit	Geneate MIT license text file in current directory.`

	fmt.Printf("%v\n", helptext)
}

func info(licensename string) {
	//fmt.Printf("%v\n", licensename[0])
	fmt.Println(licensename)
	// file, e := ioutil.ReadFile("licenses.json")
	// if e != nil {
	// 	log.Fatal(e)
	// }
	// var jsontype map[string][]map[string]interface{}
	// err := json.Unmarshal(file, &jsontype)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //fmt.Printf("Results: %v \n", jsontype["licenses"])
	// for i := range jsontype["licenses"] {
	// 	item := jsontype["licenses"][i]
	// 	//fmt.Printf("%v\n",item["name"])
	// 	if item["name"] == licensename[0] {
	// 		fmt.Printf("%v\n%v\n%v", item["name"], item["version"], item["text"])
	// 	}
	// }
}

func list() {
	fmt.Println("List of all available OSS licenses:")
}

func generate(licensename string) {
	fmt.Println(licensename)
}

//parseArgs parses command line arguments and calls appropriate functions.
func parseArgs(args []string) {
	// if len(args) != 0 {
	// 	switch args[0] {
	// 	case "help":
	// 		help()
	// 	case "info":
	// 		if len(args) != 2 {
	// 			fmt.Println("Incorrect usage!")
	// 			fmt.Println("Please provide name of the license for which more information is required.")
	// 			fmt.Println("Type `licensethis help` for help on proper usage.")
	// 		} else {
	// 			info(args[1])
	// 		}
	// 	case "list":
	// 		list()
	// 	case "generate":
	// 		if len(args) != 2 {
	// 			fmt.Println("Incorrect usage!")
	// 			fmt.Println("Please provide name of the license to generate a license file.")
	// 			fmt.Println("Type `licensethis help` for help on proper usage.")
	// 		} else {
	// 			generate(args[1])
	// 		}
	// 	default:
	// 		fmt.Println("Incorrect usage!")
	// 		fmt.Println("Type `licensethis help` for help on proper usage.")
	// 	}
	// }

	//better structed version of the above block. but very strict. does not give personalized error messages.
	//for example, generate should print more help regarding generate usage

	// !!---FUTURE---!!
	//might have to look into a way to accept more than one license name for `info`.
	if len(args) == 1 {
		switch args[0] {
		case "help":
			help()
		case "list":
			list()
		default:
			fmt.Println("Incorrect usage!")
			fmt.Println("Type `licensethis help` for help on proper usage.")
		}
	} else if len(args) == 2 {
		switch args[0] {
		case "info":
			info(args[1])
		case "generate":
			generate(args[1])
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

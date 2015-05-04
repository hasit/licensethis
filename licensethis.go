package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Licenselist holds short and long names of licenses
type Licenselist []struct {
	Short string `json:"short"`
	Long  string `json:"long"`
}

//Licenseinfo holds information of licenses
type Licenseinfo []struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Source      string `json:"source"`
	Description string `json:"description"`
	Tags        struct {
		Required  []string `json:"required"`
		Permitted []string `json:"permitted"`
		Forbidden []string `json:"forbidden"`
	} `json:"tags"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getfilepath(filename string) string {
	gopath := os.Getenv("GOPATH")
	licensethispath := gopath + "/src/github.com/hasit/licensethis/"
	filepath := licensethispath + filename

	return filepath
}

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

func list() {
	listfilepath := getfilepath("licenselist.json")
	file, err1 := ioutil.ReadFile(listfilepath)
	check(err1)

	var licenselist Licenselist
	err2 := json.Unmarshal(file, &licenselist)
	check(err2)

	fmt.Println("List of all available OSS licenses:")
	for i := range licenselist {
		fmt.Printf("%d. %v - %v\n", i, licenselist[i].Short, licenselist[i].Long)
	}
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

func generate(licensename string) {
	fmt.Println(licensename)
}

// !!---FUTURE---!!
//might have to look into a way to accept more than one license name for `info`.
func parseArgs(args []string) {
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

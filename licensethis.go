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

// !!---FUTURE---!!
// alow substrings of license names as arguments too
func info(licensename string) {
	infofilepath := getfilepath("licenseinfo.json")
	file, err1 := ioutil.ReadFile(infofilepath)
	check(err1)

	var licenseinfo Licenseinfo
	err2 := json.Unmarshal(file, &licenseinfo)
	check(err2)

	for i := range licenseinfo {
		if licenseinfo[i].Name == licensename {
			fmt.Printf("Title: %v\n", licenseinfo[i].Title)
			fmt.Printf("Category: %v\n", licenseinfo[i].Category)
			fmt.Printf("Source: %v\n", licenseinfo[i].Source)
			fmt.Printf("Description: %v\n", licenseinfo[i].Description)
			fmt.Println("Tags: ")
			fmt.Println("  Required: ")
			for r := range licenseinfo[i].Tags.Required {
				fmt.Printf("    -%v\n", licenseinfo[i].Tags.Required[r])
			}
			fmt.Println("  Permitted: ")
			for r := range licenseinfo[i].Tags.Permitted {
				fmt.Printf("    -%v\n", licenseinfo[i].Tags.Permitted[r])
			}
			fmt.Println("  Forbidden: ")
			for r := range licenseinfo[i].Tags.Forbidden {
				fmt.Printf("    -%v\n", licenseinfo[i].Tags.Forbidden[r])
			}
		}
	}
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

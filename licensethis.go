
package main
import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"log"
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
	helptext:=`
# licensethis
Add an OSS license to your project.

### SYNOPSIS

licensethis <option> <license-name>
licensethis help
licensethis config
licensethis list
licensethis info license-name
licensethis generate license-name

DESCRIPTION

licensethis lets you add an OSS license to your project with ease.

OPTIONS

These options let you do a variety of things.

  * help
	Open project page in the default browser. This provides more information on usage of licensethis with detailed information about each option.
  * config
	Configure user's full name for generating license files. User will be prompted to enter their full name for the first time. This information is stored in a JSON file at *~/.config/licensethis/user.json*.
  * list
	List all OSS licenses available on *http://choosealicense.com/*. This only prints a list of licenses with no more information. This could serve as a quick preview of all available licenses to choose from.
  * info license-name
	Get more information for <license-name>. This prints a synopsis of the license, tags (required, permitted, forbidden) associated to the license and a link to *http://choosealicense.com/licenses/<license-name>*.
  * generate license-name
	Generate <license-name> license for your project. This generates a LICENSE.txt file in your current directory with YEAR and FULLNAME already filled in. It picks up the name from config file automatically.

## CONTRIBUTORS 
* Hasit Mistry 
	Personal: [hasit.github.io](http://hasit.github.io/) 
	Github: [hasit](https://github.com/hasit)
* Anuj Deshpande 
	Personal: [anujdeshpande.com](http://www.anujdeshpande.com/) 
	Github: [anujdeshpande](https://github.com/anujdeshpande/)\n`
	fmt.Printf("%v",helptext)
}

func userConfig() {
	first := os.Getenv("FIRSTNAME")
	last := os.Getenv("LASTNAME")

	//if first != nil; last != nil {
	fmt.Printf("First: %v\n", first)
	fmt.Printf("Last: %v\n", last)
	//}
}

func info(licensename []string) {
	//fmt.Printf("%v\n", licensename[0])
	
	file, e := ioutil.ReadFile("licenses.json")
	if e != nil {
		log.Fatal(e)
	}
	var jsontype map[string][]map[string]interface{}
	err := json.Unmarshal (file,&jsontype)
	if err!=nil {
		log.Fatal(err)
	}
	//fmt.Printf("Results: %v \n", jsontype["licenses"])
	for i := range jsontype["licenses"] {
		item := jsontype["licenses"][i]
		//fmt.Printf("%v\n",item["name"])
		if item["name"]==licensename[0] {
			fmt.Printf("%v\n%v\n%v", item["name"], item["version"],item["text"])
			}
	}

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

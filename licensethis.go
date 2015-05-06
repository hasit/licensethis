package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"time"
)

//Licenselist holds short and long names of licenses
type Licenselist struct {
	Short string `json:"short"`
	Long  string `json:"long"`
}

//Licenseinfo holds information of licenses
type Licenseinfo struct {
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

//User holds full name of the user
type User struct {
	Fullname string
	Year     int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getfilepath(filename string) string {
	gopath := os.Getenv("GOPATH")
	licensethispath := gopath + "src/github.com/hasit/licensethis/"
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

	var licenselist []Licenselist
	err2 := json.Unmarshal(file, &licenselist)
	check(err2)

	fmt.Println("List of all available OSS licenses:")
	for i := range licenselist {
		fmt.Printf("%d. %v - %v\n", i, licenselist[i].Short, licenselist[i].Long)
	}
}

func info(licensename string) {
	infofilepath := getfilepath("licenseinfo.json")
	file, err1 := ioutil.ReadFile(infofilepath)
	check(err1)

	var licenseinfo []Licenseinfo
	err2 := json.Unmarshal(file, &licenseinfo)
	check(err2)

	var closelicensenames []string

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
			return
		} else if strings.Contains(licenseinfo[i].Name, licensename) {
			closelicensenames = append(closelicensenames, licenseinfo[i].Name)
		}
	}
	if len(closelicensenames) != 0 {
		fmt.Println("Did you mean any of these?")
		for i := range closelicensenames {
			fmt.Printf("%v\n", closelicensenames[i])
		}
	} else {
		fmt.Println("No such license was found")
		fmt.Println("Type `licensethis list` for a list of available licenses")
	}
}

func generate(licensename string) {
	templatepath := "templates/" + licensename + ".txt"
	licensetemplatepath := getfilepath(templatepath)
	file, err1 := ioutil.ReadFile(licensetemplatepath)
	check(err1)

	filecontent := string(file)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter full name: ")
	name, err2 := reader.ReadString('\n')
	check(err2)

	nowyear := int(time.Now().Year())

	var user = User{name, nowyear}

	t := template.Must(template.New("user").Parse(filecontent))
	err3 := t.Execute(os.Stdout, user)
	check(err3)
}

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

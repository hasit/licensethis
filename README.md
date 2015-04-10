# licensethis
Add an OSS license to your project.

## Usage:

```
$ licensethis command [arguments]
```

## Commands:
Following are the available commands:

command | description
---- | ----
config | store and read user's full name from a local JSON file
generate | generate LICENSE.txt in the current directory using stored name
help | display help related to licensethis
list | print list of OSS licenses with an option to list using tags

**config**
	
	$ licensethis config -n John Doe
Generate ~/.config/licensethis/config.json file and store the name 'John Doe' in it.

	$ licensethis config -s
Read ~/.config/licensethis/config.json file and print the name from it.

**generate**

	$ licensethis generate -l MIT
Generate LICENSE.txt in the current directory with MIT License.
	
**help**
	
	$ licensethis help
Print the help file.
	
**list**

	$ licensethis list [-t <tag1> <tag2> <tag3>...]
Print a list of all OSS licenses. Additionally, provide tags to search using specific tags.
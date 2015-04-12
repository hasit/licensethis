# licensethis
Add an oss license to your project.

## SYNOPSIS

`licensethis` `<option>` `<license-name>`<br />
`licensethis` `help`<br />
`licensethis` `config`<br />
`licensethis` `list`<br />
`licensethis` `info` `<license-name>`<br />
`licensethis` `generate` `<license-name>`

## DESCRIPTION

**licensethis** lets you add an OSS license to your project with ease.

## OPTIONS

These options let you do a variety of things.

  * `help`<br />
	Print help file. This provides more information on usage of **licensethis** with detailed information about each option.
  * `config`<br />
	Configure user's full name for generating license files. User will be prompted to enter their full name for the first time. This information is stored in a JSON file at *~/.config/licensethis/user.json*.
  * `list`<br />
	List all OSS licenses available on *http://choosealicense.com/*. This only prints a list of licenses with no more information. This could serve as a quick preview of all available licenses to choose from.
  * `info` `license-name`<br />
	Get more information for <license-name>. This prints a synopsis of the license, tags (required, permitted, forbidden) associated to the license and a link to *http://choosealicense.com/licenses/<license-name>*.
  * `generate` `license-name`<br />
	Generate <license-name> license for your project. This generates a LICENSE.txt file in your current directory with YEAR and FULLNAME already filled in. It picks up the name from config file automatically.

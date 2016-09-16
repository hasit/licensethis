# licensethis


Choose an OSS license for your project with ease.

## Install


```
go get github.com/hasit/licensethis
```

## Usage


```
licensethis <command> [<license-name>]
```

## Commands

### help

Show help text.

```
~/g/licensethis (master) $ licensethis help
licensethis - Choose an OSS license for your project with ease.

Usage:
licensethis help
licensethis list
licensethis info <license-name>
licensethis generate <license-name>

Commands:
help            Show this help text.
list            Show a list of all available OSS licenses.
info            Show more information for a license.
generate        Generate LICENSE.txt file in current folder after asking for author name.

Examples:
licensethis help                Show this help text.
licensethis info mit            Show more information for MIT license.
licensethis generate mit        Geneate MIT license text file in current directory.
```

![help](/gifs/help.gif)

### list

Show a list of all available OSS licenses.

```
~/g/licensethis (master) $ licensethis list
List of all available OSS licenses:
0. agpl-3.0 - GNU Affero General Public License v3.0
1. apache-2.0 - Apache License 2.0
2. artistic-2.0 - Artistic License 2.0
3. bsd-2-clause - BSD 2-clause Simplified License
4. bsd-3-clause - BSD 3-clause New or Revised License
5. cc0-1.0 - Creative Commons Zero v1.0 Universal
6. epl-1.0 - Eclipse Public License 1.0
7. gpl-2.0 - GNU General Public License v2.0
8. gpl-3.0 - GNU General Public License v3.0
9. isc - ISC License
10. lgpl-2.1 - GNU Lesser General Public License v2.1
11. lgpl-3.0 - GNU Lesser General Public License v3.0
12. mit - MIT License
13. mpl-2.0 - Mozilla Public License 2.0
14. no-license - No License
15. unlicense - The Unlicense
```

![list](/gifs/list.gif)

### info

Show more information for a license.

```
~/g/licensethis (master) $ licensethis info mit
Title: MIT License
Category: MIT
Source: http://opensource.org/licenses/MIT
Description: A permissive license that is short and to the point. It lets people do anything with your code with proper attribution and without warranty.
Tags: 
  Required: 
    -include-copyright
  Permitted: 
    -commercial-use
    -modifications
    -distribution
    -sublicense
    -private-use
  Forbidden: 
    -no-liability
```

![info](/gifs/info.gif)

### generate

Generate LICENSE.txt file in current folder after asking for author name.

![generate](/gifs/generate.gif)

## Contributors

-	Hasit Mistry <br /> Personal: [hasit.me](http://hasit.me/) <br /> Github: [hasit](https://github.com/hasit)
-	Anuj Deshpande <br /> Personal: [anujdeshpande.com](http://www.anujdeshpande.com/) <br /> Github: [anujdeshpande](https://github.com/anujdeshpande/)

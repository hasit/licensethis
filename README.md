# licensethis
Add an OSS license to your project


### Usage
	
	licensethis [options] <command>

### Options

	--help
		Prints the synopsis and a list of the most commonly used commands.
	
	--list
		Lists all the possible licensing options.
		
	
	--tags tag1 tag2 ...
		Will show a list of licenses with the mentioned tags. Tags can take the following values 
		
			attribution
			sharealike
			noncommercial
			permissive
			nonpermissive
			copyleft
			copyright
			liable
 					
### Command

	
	config
		licensethis uses a simple text format to store customizations that are per user. 
		Such a configuration file may look like this:
		
		; user identity
           	[user]
                   name = "John Doe"
                   email = "john@doe.com"
                   

	name-of-license
		Add above named license to current directory. Use globally set user name and email, if defined.
                   
                   
### Usage

	licensethis config user.name John Doe
		Sets name variable globally as John Doe, to be used in copyright headers.
		
	licensethis MIT
		Add MIT license to current directory. Use globally set user name and email, if defined.
	
### Notes

	1. Choose an OSS License  
		http://choosealicense.com/

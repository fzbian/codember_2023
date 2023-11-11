package main

import "fmt"

/*
"#" +1
"@" -1
"*" i*i
"&" Prints

input example:
&###@&*&###@@##@##&######@@#####@#@#@#@##@@@@@@@@@@@@@@@*&&@@@@@@@@@####@@@@@@@@@#########&#&##@@##@@##@@##@@##@@##@@##@@##@@##@@##@@##@@##@@##@@##@@##@@&
*/

func main() {
	var input string
	var output int
	fmt.Scanln(&input)

	for _, i := range input {
		switch i {
		case '#':
			output += 1
		case '@':
			output -= 1
		case '*':
			output *= output
		case '&':
			fmt.Print(output)
		default:
			// Ignore other characters
		}
	}

}

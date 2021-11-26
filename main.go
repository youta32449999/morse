package main

import (
	"flag"
	"fmt"
	"strings"
)

var morseMap = map[string]string{
	".-":     "a",
	"-...":   "b",
	".-.-":   "c",
	"-..":    "d",
	".":      "e",
	"..-.":   "f",
	"--.":    "g",
	"....":   "h",
	"..":     "i",
	".---":   "j",
	"-.-":    "k",
	".-..":   "l",
	"--":     "m",
	"-.":     "n",
	"---":    "o",
	".--.":   "p",
	"---.-":  "r",
	"...":    "s",
	"-":      "t",
	"..-":    "u",
	"...-":   "v",
	".--":    "w",
	"-..-":   "x",
	"-.--":   "y",
	"--..":   "z",
	".----":  "1",
	"..---":  "2",
	"...--":  "3",
	"....-":  "4",
	".....":  "5",
	"-....":  "6",
	"--...":  "7",
	"---..":  "8",
	"----.":  "9",
	"-----":  "0",
	".-.-.-": ".",
	"--..--": ",",
	"---...": ":",
	"..--..": "?",
	"-....-": "-",
	"-.--.":  "(",
	"-.--.-": ")",
	"-..-.":  "/",
	"...--.": "[",
	"...---": "]",
	".....-": ";",
}

var reverseMorseMap = map[string]string{}

func init() {
	for key, value := range morseMap {
		reverseMorseMap[value] = key
	}
}

func decode(morse string) string {
	return morseMap[morse]
}

func encode(ch string) string {
	return reverseMorseMap[ch]
}

func main() {
	var inputWords string
	flag.StringVar(&inputWords, "e", "", "input string to morse string")
	var inputMorse string
	flag.StringVar(&inputMorse, "d", "", "input morse to string")
	flag.Parse()

	if inputMorse != "" {
		for _, word := range strings.Split(inputMorse, " ") {
			fmt.Print(decode(word))
		}
		fmt.Println("")
	}

	if inputWords != "" {
		for _, ch := range strings.Split(inputWords, "") {
			fmt.Printf("%s ", encode(ch))
		}
		fmt.Println("")
	}
}

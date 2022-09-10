package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	option := os.Args[1]
	var err error
	var stdout string
	stdin := strings.Join(os.Args[2:], "+")
	switch option {
	case "-a":
		stdout, err = lookupAntonyms(stdin)
	case "-c":
		stdout, err = lookupCorrection(stdin)
	case "-d":
		stdout, err = lookupDefinition(stdin)
	case "-s":
		stdout, err = lookupSynonyms(stdin)
	case "-u":
		stdout, err = lookupUsage(stdin)
	case "-w":
		stdout, err = lookupWordOfTheDay()
	default:
		err = ErrInvalidOption
	}
	if err != nil {
		fmt.Println(err)
	}
	if stdout != "" {
		fmt.Println(stdout)
	}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	jsonparser "github.com/kolendnik/json-parser"
)

func main() {

	csvFlag := flag.Bool("csv", false, "export data to csv")
	flag.Parse()

	var stdin []byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stdin = append(stdin, scanner.Bytes()...)
	}
	if err := scanner.Err(); err != nil {
		printError(err)
		os.Exit(1)
	}

	w, err := jsonparser.ParseInput(stdin)
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	if *csvFlag {
		jsonparser.ToCsv(w)
	} else {
		jsonparser.ToStdout(w)
	}
}

func printError(err error) {
	fmt.Println(err.Error())
}

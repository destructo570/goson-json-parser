package main

import (
	"fmt"
	"goson/app"
)

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Please provide input")
	// 	return
	// }

	// input := os.Args[1]
	// fmt.Println("Input received:", input)
	// Parse(input)
	Parse(`{"name":69.60.}`)
}

func Error(line int, message string) {
	Report(line, "", message)
}

func Report(line int, where string, message string) {
	fmt.Printf("[line %d ] Error %s : %s", line, where, message)
}

func Parse(source string) {
	scanner := app.Scanner{
		Source: source,
		Line:   1,
	}

	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		println(token.ToString())
	}

	// return {"hello": "hello"}
}

func Stringify(object any) string {
	return "Hehehehe"
}

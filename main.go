package main

import (
	"encoding/json"
	"fmt"
	"goson/app"
	"goson/models"
)

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Please provide input")
	// 	return
	// }

	// input := os.Args[1]
	// fmt.Println("Input received:", input)
	// Parse(input)

	// Parse(`{
	// 	"name": 6990809.90
	// }`)
	Parse(`{
		"name": "Michael Smith",
		"age" : 50,
		"Net worth" : 6990809.90,
		"isAlive" : true,
		"hobbies" : null,
		"isLoggedIn" : false
	}`)

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

	var std any
	if err := json.Unmarshal([]byte(source), &std); err != nil {
		fmt.Println("std json error:", err)
	} else {
		fmt.Printf("std json: %#v\n", std)
	}

	parser := app.Parser{}

	tokens := scanner.ScanTokens()

	value, err := parser.Parse(tokens)

	if err != nil {
		fmt.Println("Error", err)
	}

	// fmt.Println("Parsed: ", value)

	// // 3. Convert your AST to std-shape and compare
	stdFromMine := models.ToStd(value)
	fmt.Printf("my std: %#v\n", stdFromMine)

	// for _, token := range tokens {
	// 	println(token.ToString())
	// }

	// return {"hello": "hello"}
}

func Stringify(object any) string {
	return "Hehehehe"
}

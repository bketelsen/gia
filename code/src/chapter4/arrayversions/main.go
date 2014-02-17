package main

import "fmt"

var versions [10]string
var currentVersion int

func main() {
	Changed("Welcome to Go Arrays")
	Changed("Welcome to Go Arrays 2!")
	Undo()

	fmt.Println(DocumentText())
}

func DocumentText() string {
	return versions[currentVersion] //<co id="retrievestring" />
}

func Changed(text string) {
	versions[currentVersion+1] = text  //<co id="updatenextindex" />
	currentVersion++  //<co id="setcurrentversion" />
}

func Undo() {
	currentVersion--  //<co id="decrementversion" />
}

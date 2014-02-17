package main

import "fmt"

var versions []string

func main() {

	versions = make([]string, 1)  //<co id="makeinitialslice" />

	Changed("Welcome to Go Arrays")
	Changed("Welcome to Go Arrays 2!")
	Undo()

	fmt.Println(DocumentText())
}

func DocumentText() string {

	return versions[len(versions)-1] //<co id="retrievelastindex" />
}


func Changed(text string) {
	versions = append(versions, text)  //<co id="appendversions" />
}

func Undo() {
	versions = versions[:len(versions)-1]  //<co id="undoversions" />
}

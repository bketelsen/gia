package main

import (
	"fmt"
)

//<start id="makeseven">
func makeSeven(key string, ageMap map[string]int) {
	ageMap[key] = 7
}

//<end id="makeseven">
func main() {
	//<start id="mapdeclaration">
	var ages map[string]int
	//<end id="mapdeclaration">
	//<start id="makemap">
	ages = make(map[string]int)
	//<end id="makemap">
	//<start id="simpledeclare">
	a := make(map[string]int)
	//<end id="simpledeclare">

	//<start id="assignmap">
	ages["Bobby"] = 5
	//<end id="assignmap">

	//<start id="changeinfunc">
	ages["Bobby"] = 5
	makeSeven("Bobby", ages)
	fmt.Printf("Bobby is %d years old.", ages["Bobby"])
	//<end id="changeinfunc">

	//<start id="retrievemap">
	bobbysAge := ages["Bobby"]
	//<end id="retrievemap">

	//<start id="deletemap">
	delete(ages, "Bobby")
	//<end id="deletemap">

	//<start id="blankident">
	if _, ok := ages["Susie"]; !ok {
		fmt.Printf("We don't know Susie")
	}
	//<end id="blankident">

	//<start id="retrievenil">
	garysAge := ages["Gary"]
	fmt.Printf("Gary is %d years old.", garysAge)
	//<end id="retrievenil">

	//<start id="commaok">
	samsAge, ok := ages["Sam"]
	if ok {
		fmt.Printf("Sam is %d years old.", samsAge)
	} else {
		fmt.Printf("We don't know Sam!")
	}
	//<end id="commaok">

	//<start id="bettercommaok">
	if samsAge, ok := ages["Sam"]; ok {
		fmt.Printf("Sam is %d years old.", samsAge)
	} else {
		fmt.Printf("We don't know Sam!")
	}
	//<end id="bettercommaok">

	fmt.Println(ages, a, bobbysAge, garysAge)

}

package main

import "fmt"

func main() {

	//<start id="forlen">
	students := []string{"Nathan", "Evelyn", "Lauren", "Anthony"}
	for i := 0; i < len(students); i++ {
		fmt.Printf("Child %d has name %s \n", i, students[i])
	}
	//<end id="forlen">
	//<start id="forrange">
	children := []string{"Nathan", "Evelyn", "Lauren", "Anthony"}
	for i, child := range children {
		fmt.Printf("Child %d has name %s \n", i, child)
	}
	//<end id="forrange">
	//<start id="blankindex">
	for _, child := range children {
		fmt.Println(child)
	}
	//<end id="blankindex">
	//<start id="blankdata">
	for i, _ := range children {
		fmt.Println("Child ", i)
	}
	//<end id="blankdata">

	//<start id="maprange">
	ages := make(map[string]int)
	ages["Bobby"] = 5
	ages["Susy"] = 7
	ages["Sammy"] = 6
	for key, value := range ages {
		fmt.Printf("%s is %d \n", key, value)
	}
	//<end id="maprange">

}

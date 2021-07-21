package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
)

func main() {
	test1()
}


type Person struct {
	ID			int			`json:"id"`
	Name		string
	Age			int
	Salary		float64		`json:"sal,omitempty"`
	managerId	int			
}

func test1() {

	p1 := Person{1001, "ali", 30, 2500.5, 101}

	bytes, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error marshaling.", err)
		return
	}
	fmt.Printf("%s\n", bytes)

	bytes, err = json.MarshalIndent(p1, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling.", err)
		return
	}
	fmt.Printf("%s\n", bytes)

	var p2 *Person = new(Person)
	err = json.Unmarshal(bytes, p2)
	if err != nil {
		fmt.Println("Error unmarshaling.", err)
		return
	}

	fmt.Println(p2)
}
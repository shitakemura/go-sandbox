package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

type Employee struct {
	Person
	ManagerID int
}

type Contractor struct {
	Person
	CompanyID int
}

func main() {
	employee := []Employee{
		Employee{
			Person: Person{
				LastName: "Doe", FirstName: "John", Address: "hoge",
			},
		},
		Employee{
			Person: Person{
				LastName: "Campbell", FirstName: "David",
			},
		},
	}

	data, _ := json.Marshal(employee)
	fmt.Printf("%s\n", data)

	var decoded []Employee
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v", decoded)
}

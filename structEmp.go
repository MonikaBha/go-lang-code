// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

func PrintEmployee(e Employee) {
	fmt.Printf("ID = %d\nName = %s\n", e.ID, e.Name)
}

func main() {
	var emp Employee
	emp.ID = 1
	emp.Name = "Petia Pyatochkin"
	PrintEmployee(emp)
}

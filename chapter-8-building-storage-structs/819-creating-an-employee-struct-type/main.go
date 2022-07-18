package main

import (
	"fmt"
	"github.com/lethang7794/magazine"
)

func main() {
	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000

	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
}

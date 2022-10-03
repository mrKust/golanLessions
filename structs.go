package main

import "fmt"

type employee struct{
	name string
	sex string
	age int
	salary int
}

func newEmployee(name, sex string, age, salary int) employee {
	return employee{
		name: name,
		sex: sex,
		age: age,
		salary: salary,
	}
}

func (e employee) employeeToString() string {
	return fmt.Sprintf("Employee: name %s, sex %s, age %d, salary %d", e.name, e.sex, e.age, e.salary)
}

func (e *employee)changeName(name string) {
	e.name = name
}

func main() {
	
	employee1 := employee{
		name: "Vasya",
		sex: "Male",
		age: 25,
		salary: 2,
	}

	employee2 := newEmployee("Vera", "Female", 24, 4)

	fmt.Printf("%s \n", employee1.employeeToString())
	employee2.changeName("Oleg")
	fmt.Printf("%s \n", employee2.employeeToString())


}

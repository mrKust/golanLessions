package main

import (
	"fmt"
	"time"
)

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

type storage interface {
	insert(e employee) error
	get(id int) (employee, error)
	delete(id int) error
}



func main() {
	
	employee1 := newEmployee("Vasya","Male",25,2)

	employee2 := newEmployee("Vera", "Female", 24, 4)

	var res uint8

	go count(&res)

	fmt.Printf("%s \n", employee1.employeeToString())
	employee2.changeName("Oleg")
	fmt.Printf("%s \n", employee2.employeeToString())

	fmt.Printf("result = %d", res)


}

func count(in *uint8) {
	time.Sleep(10)
	fmt.Println("Begin counting")
	*in = 2
	*in += 4
}

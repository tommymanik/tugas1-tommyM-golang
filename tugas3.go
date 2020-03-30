package main

import "fmt"

func main() {

	var s1 = Employee{"john wick", "Pratama", "Marketing", "2.000.000"}

	var employee = s1.getEmpl()
	fmt.Println("FirstName : ", employee.firstName)
	fmt.Println("Lastname : ", employee.lastName)
	fmt.Println("Job : ", employee.job)
	fmt.Println("Salary : ", employee.salary)

}

type Employee struct {
	firstName string
	lastName  string
	job       string
	salary    string
}

func (e Employee) setEmpl(fName string, lName string, joobs string, sal string) {
	e.firstName = fName
	e.lastName = lName
	e.job = joobs
	e.salary = sal
}

func (e Employee) getEmpl() Employee {
	return e
}

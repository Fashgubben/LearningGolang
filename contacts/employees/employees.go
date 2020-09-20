package employees

import "fmt"

// Define Employee struct
type Employee struct {
	FirstName   string
	LastName    string
	Department  string
	PhoneNumber string
	Email       string
}

// Returns a struct with requested vaulues
func CreateEmployee(firstName, lastName, department, phoneNumber, email string) Employee {
	return Employee{
		FirstName:   firstName,
		LastName:    lastName,
		Department:  department,
		PhoneNumber: phoneNumber,
		Email:       email,
	}
}

// Set new department method (pointer reciever)
func (emp *Employee) ChangeDepartment(newDepartment string) {
	emp.Department = newDepartment
	fmt.Println("New department set.")
}

// Set new phone number method (pointer reciever)
func (emp *Employee) ChangePhone(newNumber string) {
	emp.PhoneNumber = newNumber
	fmt.Println("New phone number set.")
}

// Set new email method (pointer reciever)
func (emp *Employee) ChangeEmail(newEmail string) {
	emp.Email = newEmail
	fmt.Println("New e-mail set.")
}

// Print selected employee (value reciever)
func (emp Employee) PrintEmployee() {
	fmt.Println("\n" + emp.FirstName + " " + emp.LastName)
	fmt.Println(emp.Department)
	fmt.Println(emp.PhoneNumber)
	fmt.Println(emp.Email)
}

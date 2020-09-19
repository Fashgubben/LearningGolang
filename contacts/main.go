package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fashgubben/LearningGolang/contacts/csvutils"
	"github.com/fashgubben/LearningGolang/contacts/employees"
)

// Requests specific input from user
func getInput(request string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n" + request)
	usrInput, _ := reader.ReadString('\n')
	return strings.TrimSpace(usrInput)
}

// Print all stored employees on terminal
func printEmployees(allContacts []employees.Employee) {
	for _, emp := range allContacts {
		fmt.Println("\n" + emp.FirstName + " " + emp.LastName)
		fmt.Println(emp.Department)
		fmt.Println(emp.PhoneNumber)
		fmt.Println(emp.Email)
	}
}

// Creates a new struct and appends it to slice
func addStruct(allContacts []employees.Employee, fn, ln, d, pn, em string) []employees.Employee {
	emp := employees.CreateEmployee(fn, ln, d, pn, em)
	allContacts = append(allContacts, emp)
	return allContacts
}

// Add new employee
func addEmployee(allContacts []employees.Employee) []employees.Employee {

	// Request
	fn := getInput("Enter first name:")
	ln := getInput("Enter last name:")
	d := getInput("Enter department:")
	pn := getInput("Enter phone number:")
	em := getInput("Enter e-mail:")

	// Add struct to list
	allContacts = addStruct(allContacts, fn, ln, d, pn, em)

	// TODO: Add row to csv
	csvutils.AddEmployeeToCsv(allContacts, "employees.csv")

	return allContacts
}

// Lets user choose from menu
func menu(allContacts []employees.Employee) {
	activeLoop := true
	for activeLoop == true {
		fmt.Println("\nEmployee catalouge")
		fmt.Println("\n[1] - Add new employee")
		fmt.Println("[2] - Print all employees")
		fmt.Println("[3] - Search for employee")
		fmt.Println("[4] - Exit")

		menuChoice := getInput("Enter a menu number:")
		switch menuChoice {
		case "1":
			allContacts = addEmployee(allContacts)
		case "2":
			printEmployees(allContacts)
		case "3":
			// Search for employee
		case "4":
			fmt.Println("Good bye.")
			activeLoop = false
		default:
			fmt.Println("\nPlease enter a valid number.")
		}
	}
}

func main() {
	allContacts := csvutils.GetAllContacts("employees.csv")
	menu(allContacts)
}

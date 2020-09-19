package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/fashgubben/LearningGolang/contacts/employees"
)

// Requests specific input from user
func getInput(request string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n" + request)
	usrInput, _ := reader.ReadString('\n')
	return strings.TrimSpace(usrInput)
}

// Reads from csv-file and stores all employees in slice
func getAllContacts(file string) []employees.Employee {
	csvFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	allContacts := []employees.Employee{}
	for _, line := range csvLines {

		emp := employees.CreateEmployee(line[0], line[1], line[2], line[3], line[4])
		allContacts = append(allContacts, emp)
	}
	return allContacts
}

// Print all stored employees
func printEmployees(allContacts []employees.Employee) {
	for i, emp := range allContacts {
		// Ignore csv titles
		if i == 0 {
			continue
		}
		fmt.Println("\n" + emp.FirstName + " " + emp.LastName)
		fmt.Println(emp.Department)
		fmt.Println(emp.PhoneNumber)
		fmt.Println(emp.Email)
	}
}

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
	allContacts := getAllContacts("employees.csv")
	menu(allContacts)
}

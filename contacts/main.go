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
func printallEmployeesSlice(allEmployeesSlice []employees.Employee) {
	for _, emp := range allEmployeesSlice {
		printEmployee(emp)
	}
}

// Print selected employee
func printEmployee(emp employees.Employee) {
	fmt.Println("\n" + emp.FirstName + " " + emp.LastName)
	fmt.Println(emp.Department)
	fmt.Println(emp.PhoneNumber)
	fmt.Println(emp.Email)
}

// Creates a new struct and appends it to slice
func addStruct(allEmployeesSlice []employees.Employee, fn, ln, d, pn, em string) []employees.Employee {
	emp := employees.CreateEmployee(fn, ln, d, pn, em)
	allEmployeesSlice = append(allEmployeesSlice, emp)
	return allEmployeesSlice
}

// Add new employee
func addEmployee(allEmployeesSlice []employees.Employee) []employees.Employee {

	// Request
	fn := getInput("Enter first name:")
	ln := getInput("Enter last name:")
	d := getInput("Enter department:")
	pn := getInput("Enter phone number:")
	em := getInput("Enter e-mail:")

	// Add struct to list
	allEmployeesSlice = addStruct(allEmployeesSlice, fn, ln, d, pn, em)

	// TODO: Add row to csv
	csvutils.AddEmployeeToCsv(allEmployeesSlice, "employees.csv")

	return allEmployeesSlice
}

// Checks if first string contains second string
func stringContains(first, second string) bool {
	if strings.Contains(strings.ToUpper(first), strings.ToUpper(second)) {
		return true
	} else {
		return false
	}
}

// Prints all employees where user input matches any struct value
func searchForEmployee(allEmployeesSlice []employees.Employee) {

	// Counter to keep track on hits
	counter := 0

	searchWord := getInput("Enter search word: ")
	for _, emp := range allEmployeesSlice {
		empAttributes := [5]string{emp.FirstName, emp.LastName, emp.Department, emp.PhoneNumber, emp.Email}
		for _, attribute := range empAttributes {
			if stringContains(attribute, searchWord) {
				printEmployee(emp)
				counter += 1
				break
			}
		}
	}
	if counter == 0 {
		fmt.Println("\nSorry, no matches on \"" + searchWord + "\".")
	} else {
		fmt.Println("\n", counter, " matches on \""+searchWord+"\".")
	}
}

// Lets user choose from menu
func menu(allEmployeesSlice []employees.Employee) {
	activeLoop := true
	for activeLoop == true {
		fmt.Println("\nEmployee catalouge")
		fmt.Println("\n[1] - Add new employee")
		fmt.Println("[2] - Print all employees")
		fmt.Println("[3] - Search for employee")
		fmt.Println("[4] - Edit employee")
		fmt.Println("[5] - Exit")

		menuChoice := getInput("Enter a menu number:")
		switch menuChoice {
		case "1":
			allEmployeesSlice = addEmployee(allEmployeesSlice)
		case "2":
			printallEmployeesSlice(allEmployeesSlice)
		case "3":
			searchForEmployee(allEmployeesSlice)
		case "4":
			//P
		case "5":
			fmt.Println("Good bye.")
			activeLoop = false
		default:
			fmt.Println("\nPlease enter a valid number.")
		}
	}
}

func main() {
	allEmployeesSlice := csvutils.GetallEmployeesSlice("employees.csv")
	menu(allEmployeesSlice)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fashgubben/LearningGolang/contacts/csvutils"
	"github.com/fashgubben/LearningGolang/contacts/employees"
)

// Creates a new struct and appends it to slice
func addStruct(allEmployeesSlice []employees.Employee, fn, ln, d, pn, em string) []employees.Employee {
	emp := employees.CreateEmployee(fn, ln, d, pn, em)
	allEmployeesSlice = append(allEmployeesSlice, emp)
	return allEmployeesSlice
}

// Returns message depending on the amount of occurrences
func returnMatchMessage(counter int, searchWord string) string {
	message := ""
	if counter == 0 {
		message = "\nSorry, no matches on \"" + searchWord + "\"."
	} else {
		strCounter := strconv.Itoa(counter)
		message = "\n" + strCounter + " matches on \"" + searchWord + "\"."
	}
	return message
}

// Checks if first string contains second string
func stringContains(first, second string) bool {
	if strings.Contains(strings.ToUpper(first), strings.ToUpper(second)) {
		return true
	} else {
		return false
	}
}

// Loops through employee and prints info on hit
func loopThroughEmps(searchWord string, emp employees.Employee, counter int) int {
	empAttributes := [5]string{emp.FirstName, emp.LastName, emp.Department,
		emp.PhoneNumber, emp.Email}
	for _, attribute := range empAttributes {
		if stringContains(attribute, searchWord) {
			emp.PrintEmployee()
			counter += 1
			return counter
		}
	}
	return counter
}

// Prints all employees where user input matches any struct value
func searchForEmployee(allEmployeesSlice []employees.Employee, searchWord string) string {
	// Counter to keep track on hits
	counter := 0
	for _, emp := range allEmployeesSlice {
		counter = loopThroughEmps(searchWord, emp, counter)
	}
	return returnMatchMessage(counter, searchWord)
}

// Print all stored employees on terminal
func printallEmployeesSlice(allEmployeesSlice []employees.Employee) {
	for _, emp := range allEmployeesSlice {
		emp.PrintEmployee()
	}
}

// Add new employee
func addEmployee(allEmployeesSlice []employees.Employee, newEmployee []string) []employees.Employee {
	// Add struct to list
	allEmployeesSlice = addStruct(allEmployeesSlice, newEmployee[0], newEmployee[1],
		newEmployee[2], newEmployee[3], newEmployee[4])
	// Add row to csv
	csvutils.AddEmployeeToCsv(allEmployeesSlice, "employees.csv")
	return allEmployeesSlice
}

// Makes requests to be inputed by user
func getNewEmployeeInfo() []string {

	fn := getInput("Enter first name:")
	ln := getInput("Enter last name:")
	d := getInput("Enter department:")
	pn := getInput("Enter phone number:")
	em := getInput("Enter e-mail:")
	return []string{fn, ln, d, pn, em}
}

// Requests specific input from user
func getInput(request string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n" + request)
	usrInput, _ := reader.ReadString('\n')
	return strings.TrimSpace(usrInput)
}

// Returns menu options
func printMenu() string {
	menu := "\nEmployee catalouge\n" +
		"\n[1] - Add new employee\n" +
		"[2] - Print all employees\n" +
		"[3] - Search for employee\n" +
		"[4] - Edit employee info\n" +
		"[5] - Exit"
	return menu
}

// Lets user choose from menu
func menu(allEmployeesSlice []employees.Employee) {
	activeLoop := true
	for activeLoop == true {
		fmt.Println(printMenu())
		menuChoice := getInput("Enter a menu number:")
		switch menuChoice {
		case "1":
			newEmployeeSlice := getNewEmployeeInfo()
			allEmployeesSlice = addEmployee(allEmployeesSlice, newEmployeeSlice)
		case "2":
			printallEmployeesSlice(allEmployeesSlice)
		case "3":
			searchWord := getInput("Enter search word: ")
			fmt.Println(searchForEmployee(allEmployeesSlice, searchWord))
		case "4":
			//Edit employee info
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fashgubben/LearningGolang/ContactBook/csvutils"
	"github.com/fashgubben/LearningGolang/ContactBook/employees"
)

// Main menu of the application
func mainMenu(allEmployeesSlice []employees.Employee) {
	activeLoop := true
	for activeLoop == true {
		fmt.Println(printMenu())
		menuChoice := getInput("Enter a mainMenu number:")
		switch menuChoice {
		case "1": // Add new employee
			newEmployeeSlice := getNewEmployeeInfo()
			allEmployeesSlice = addEmployee(allEmployeesSlice, newEmployeeSlice, "employees.csv")
		case "2": // Print all employees
			printallEmployeesSlice(allEmployeesSlice)
		case "3": // Search for employee
			searchWord := getInput("Enter search word: ")
			fmt.Println(searchForEmployee(allEmployeesSlice, searchWord))
		case "4": // Edit employee info
			allEmployeesSlice = editEmployee(allEmployeesSlice)
		case "5": // Exit
			fmt.Println("Good bye.")
			activeLoop = false
		default:
			fmt.Println("\nPlease enter a valid number.")
		}
	}
}

// Returns mainMenu options
func printMenu() string {
	mainMenu := "\nEmployee catalouge\n" +
		"\n[1] - Add new employee\n" +
		"[2] - Print all employees\n" +
		"[3] - Search for employee\n" +
		"[4] - Edit employee info\n" +
		"[5] - Exit"
	return mainMenu
}

// Requests specific input from user
func getInput(request string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n" + request)
	usrInput, _ := reader.ReadString('\n')
	return strings.TrimSpace(usrInput)
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

// Add new employee
func addEmployee(allEmployeesSlice []employees.Employee, newEmployee []string, csvFile string) []employees.Employee {
	// Add struct to list
	allEmployeesSlice = addStruct(allEmployeesSlice, newEmployee[0], newEmployee[1],
		newEmployee[2], newEmployee[3], newEmployee[4])
	// Add row to csv
	csvutils.AddEmployeeToCsv(allEmployeesSlice, csvFile)
	return allEmployeesSlice
}

// Creates a new struct and appends it to slice
func addStruct(allEmployeesSlice []employees.Employee, fn, ln, d, pn, em string) []employees.Employee {
	emp := employees.CreateEmployee(fn, ln, d, pn, em)
	allEmployeesSlice = append(allEmployeesSlice, emp)
	return allEmployeesSlice
}

// Print all stored employees on terminal
func printallEmployeesSlice(allEmployeesSlice []employees.Employee) {
	for _, emp := range allEmployeesSlice {
		emp.PrintEmployee()
	}
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

// Checks if first string contains second string
func stringContains(first, second string) bool {
	if strings.Contains(strings.ToUpper(first), strings.ToUpper(second)) {
		return true
	} else {
		return false
	}
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

// Lets user select user and category to edit
func editEmployee(allEmployeesSlice []employees.Employee) []employees.Employee {
	empChoice := selectEmployee(allEmployeesSlice)
	catChoice := selectCategory()
	csvutils.AddEmployeeToCsv(allEmployeesSlice, "employees.csv")
	return updateEmployee(allEmployeesSlice, catChoice, empChoice)

}

// Lets user select employee from a list
func selectEmployee(allEmployeesSlice []employees.Employee) int {
	fmt.Println(printAllNames(allEmployeesSlice))
	request := "\nChoose employee to edit:"
	amountOfChoices := len(allEmployeesSlice) - 1
	strEmpChoice := getValidInput(request, amountOfChoices)
	return convertStringToInt(strEmpChoice)
}

// Returns a mainMenu with names to choose from
func printAllNames(allEmployeesSlice []employees.Employee) string {
	nameList := ""
	strNum := ""
	fmt.Println("\nStored employees:")
	for i, emp := range allEmployeesSlice {
		strNum = strconv.Itoa(i)
		nameList += "[" + strNum + "] - " + emp.FirstName + " " + emp.LastName + "\n"
	}
	return nameList
}

// Asks for input from user until it's valid
func getValidInput(request string, amountOfChoices int) string {
	menuChoice := ""
	choiceIsValid := false
	for choiceIsValid == false {
		menuChoice = getInput(request)
		for i := 0; i <= amountOfChoices; i++ {
			if menuChoice == strconv.Itoa(i) {
				choiceIsValid = true
			}
		}
		if choiceIsValid == false {
			fmt.Println("Please enter a valid mainMenu number.")
		}
	}
	return menuChoice
}

func convertStringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// Lets user select what category to edit
func selectCategory() string {
	fmt.Println(printAllCategories())
	request := "\nChoose category to edit:"
	amountOfChoices := 4
	return getValidInput(request, amountOfChoices)
}

// Returns a mainMenu with attributes
func printAllCategories() string {
	fmt.Println("\nEmployee(s) you're able to edit:")
	return "\n[0] - First name\n[1] - Last name\n[2] - Department\n[3] - Phone number\n[4] - E-mail\n"
}

// Updates the employees information (struct and csv)
func updateEmployee(allEmployeesSlice []employees.Employee, catChoice string, empChoice int) []employees.Employee {
	request := "\nEnter new value:"
	newValue := getInput(request)
	return changeEmployeeInfo(allEmployeesSlice, catChoice, newValue, empChoice)
}

// Replaces right category with new value
func changeEmployeeInfo(allEmployeesSlice []employees.Employee, catChoice, newValue string, empChoice int) []employees.Employee {
	switch catChoice {
	case "0": // First name
		allEmployeesSlice[empChoice].ChangeFirstName(newValue)
	case "1": // Last name
		allEmployeesSlice[empChoice].ChangeLastName(newValue)
	case "2": // Department
		allEmployeesSlice[empChoice].ChangeDepartment(newValue)
	case "3": // Phone number
		allEmployeesSlice[empChoice].ChangePhone(newValue)
	case "4": // E-mail
		allEmployeesSlice[empChoice].ChangeEmail(newValue)
	}
	return allEmployeesSlice
}

func main() {
	allEmployeesSlice := csvutils.GetallEmployeesSlice("employees.csv")
	mainMenu(allEmployeesSlice)
}

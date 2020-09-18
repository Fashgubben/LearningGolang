package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

// Define Employee struct
type Employee struct {
	firstName   string
	lastName    string
	department  string
	phoneNumber string
	email       string
}

// Confirmation method (value reciever)
func (e Employee) saveConfirm() string {
	return "Employee " + e.firstName + " " + e.lastName + " is now saved."
}

// Set new phone number method (pointer reciever)
func (e *Employee) changePhone(newNumber string) {
	e.phoneNumber = newNumber
	fmt.Println("New phone number set.")
}

// Reads from csv-file and stores all employees in slice
func getStoredContacts() []Employee {

	csvFile, err := os.Open("employees.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	allContacts := []Employee{}

	for _, line := range csvLines {
		employee := Employee{
			firstName:   line[0],
			lastName:    line[1],
			department:  line[2],
			phoneNumber: line[3],
			email:       line[4],
		}
		allContacts = append(allContacts, employee)
	}

	return allContacts
}

// Lets user choose from menu
func menu() {

	activeLoop := true
	for activeLoop == true {

		fmt.Println("Employee catalouge")
		fmt.Println("\n[1] - Add new employee")
		fmt.Println("[2] - Print all info")
		fmt.Println("[3] - Search for employee")
		fmt.Println("[4] - Exit")

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("\nEnter a menu number: ")
		usrInput, _ := reader.ReadString('\n')

		switch usrInput {
		case "1":
			// Add new employee
		case "2":
			// Print all info
		case "3":
			// Search for employee
		case "4":
			// Exit
			fmt.Println("Good bye.")
			activeLoop = false
		default:
			fmt.Println("\nPlease enter a valid number.")

		}

	}
}

func main() {

	menu()

}

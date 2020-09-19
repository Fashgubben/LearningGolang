package csvutils

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/fashgubben/LearningGolang/contacts/employees"
)

// Reads from csv-file and stores all employees in slice
func GetAllContacts(file string) []employees.Employee {
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

// Writes row to csv-file
func AddEmployeeToCsv(allContacts []employees.Employee, file string) {

	csvFile, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	for _, emp := range allContacts {

		empSlice := []string{emp.FirstName, emp.LastName, emp.Department, emp.PhoneNumber, emp.Email}
		err := writer.Write(empSlice)
		if err != nil {
			fmt.Println(err)
		}
	}
}

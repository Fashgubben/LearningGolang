package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/fashgubben/LearningGolang/ContactBook/csvutils"
	"github.com/fashgubben/LearningGolang/ContactBook/employees"
)

// TESTING EMPLOYEE PACKAGE
func TestCreateEmployee(t *testing.T) {

	firstName := "testFirst"
	lastName := "testLast"
	department := "testDepartment"
	phoneNumber := "testPhone"
	email := "testEmail"

	testEmployee := employees.CreateEmployee(firstName, lastName, department, phoneNumber, email)

	if firstName != testEmployee.FirstName {
		t.Fail()
	} else if lastName != testEmployee.LastName {
		t.Fail()
	} else if department != testEmployee.Department {
		t.Fail()
	} else if phoneNumber != testEmployee.PhoneNumber {
		t.Fail()
	} else if email != testEmployee.Email {
		t.Fail()
	}
}

func TestChangeFirstName(t *testing.T) {
	testEmployee := employees.CreateEmployee("firstName", "lastName", "Stockholm", "101010", "email@test.com")
	testEmployee.ChangeFirstName("NewFirstName")

	if testEmployee.FirstName != "NewFirstName" {
		t.Fail()
	}
}

func TestChangeLastName(t *testing.T) {
	testEmployee := employees.CreateEmployee("firstName", "lastName", "Stockholm", "101010", "email@test.com")
	testEmployee.ChangeLastName("NewLastName")

	if testEmployee.LastName != "NewLastName" {
		t.Fail()
	}
}

func TestChangeDepartment(t *testing.T) {
	testEmployee := employees.CreateEmployee("firstName", "lastName", "Stockholm", "101010", "email@test.com")
	testEmployee.ChangeDepartment("Dublin")

	if testEmployee.Department != "Dublin" {
		t.Fail()
	}
}

func TestChangePhone(t *testing.T) {
	testEmployee := employees.CreateEmployee("firstName", "lastName", "Stockholm", "101010", "email@test.com")
	testEmployee.ChangePhone("202020")

	if testEmployee.PhoneNumber != "202020" {
		t.Fail()
	}
}

func TestChangeEmail(t *testing.T) {
	testEmployee := employees.CreateEmployee("firstName", "lastName", "Stockholm", "101010", "email@test.com")
	testEmployee.ChangeEmail("email@test.se")

	if testEmployee.Email != "email@test.se" {
		t.Fail()
	}
}

// TESTING CSVUTILS PACKAGE

func TestAddEmployeeToCsv(t *testing.T) {

	// Keepng an original slice to restore csv-file at end
	originalEployees := csvutils.GetallEmployeesSlice("employees_test.csv")

	// Append a new employee to slice
	testEmployees := csvutils.GetallEmployeesSlice("employees_test.csv")
	testEmployees = addStruct(testEmployees, "f3", "l3", "d3", "pn3", "em3")

	// Run the function we're testing
	csvutils.AddEmployeeToCsv(testEmployees, "employees_test.csv")

	// Check if contact was written to file
	resultEmployees := csvutils.GetallEmployeesSlice("employees_test.csv")
	lastEmployee := resultEmployees[len(resultEmployees)-1]
	if lastEmployee.FirstName != "f3" {
		t.Fail()
	}
	// Restore csv file
	csvutils.AddEmployeeToCsv(originalEployees, "employees_test.csv")
}

func TestGetallEmployeesSlice(t *testing.T) {

	allEmployeesSlice := csvutils.GetallEmployeesSlice("employees_test.csv")
	expectationArr1 := [5]string{"f1", "l1", "d1", "p1", "e1"}
	expectationArr2 := [5]string{"f2", "l2", "d2", "p2", "e2"}

	emp1 := allEmployeesSlice[0]
	emp2 := allEmployeesSlice[1]

	if emp1.FirstName != expectationArr1[0] {
		t.Fail()
	}

	if emp2.FirstName != expectationArr2[0] {
		t.Fail()
	}
}

// TESTING MAIN PACKAGE

func TestPrintMenu(t *testing.T) {
	expectedResult := "\nEmployee catalouge\n" +
		"\n[1] - Add new employee\n" +
		"[2] - Print all employees\n" +
		"[3] - Search for employee\n" +
		"[4] - Edit employee info\n" +
		"[5] - Exit"
	actualResult := printMenu()
	if expectedResult != actualResult {
		t.Fail()
	}
}

func TestAddEmployee(t *testing.T) {

	// Keepng an original slice to restore csv-file at end
	originalEployees := csvutils.GetallEmployeesSlice("employees_test.csv")

	// Run test
	allEmployeesSlice := csvutils.GetallEmployeesSlice("employees_test.csv")
	newEmployee := []string{"f3", "l3", "d3", "pn3", "em3"}
	allEmployeesSlice = addEmployee(allEmployeesSlice, newEmployee, "employees_test.csv")
	lastAddedEmployee := allEmployeesSlice[len(allEmployeesSlice)-1]
	if lastAddedEmployee.FirstName != "f3" {
		t.Fail()
	}
	// Restore csv to original
	csvutils.AddEmployeeToCsv(originalEployees, "employees_test.csv")
}

func TestSearchForEmployee(t *testing.T) {

	allEmployeesSlice := csvutils.GetallEmployeesSlice("employees_test.csv")

	// Test1
	searchWord := "test"
	expectedResult := "\nSorry, no matches on \"" + searchWord + "\"."
	actualResult := searchForEmployee(allEmployeesSlice, searchWord)
	if expectedResult != actualResult {
		t.Fail()
	}

	// Test 2
	intCounter := 1
	strCounter := strconv.Itoa(intCounter)
	searchWord = "f1"
	expectedResult = "\n" + strCounter + " matches on \"" + searchWord + "\"."
	actualResult = searchForEmployee(allEmployeesSlice, searchWord)
	if expectedResult != actualResult {
		t.Fail()
	}
}

func TestLoopThroughEmps(t *testing.T) {
	testEmployee := employees.CreateEmployee("a", "b", "c", "d", "e")
	counter := 0

	counter = loopThroughEmps("123", testEmployee, counter)
	if counter != 0 {
		t.Fail()
	}

	counter = loopThroughEmps("a", testEmployee, counter)
	if counter != 1 {
		t.Fail()
	}
}

func TestStringContains(t *testing.T) {

	//True
	firstTest := stringContains("Foobar", "foo")
	//False
	secondTest := stringContains("Foobar", "123")

	if firstTest != true {
		t.Fail()
	}
	if secondTest != false {
		t.Fail()
	}
}

func TestReturnMatchMessage(t *testing.T) {
	searchWord := "test"

	// Test 1
	intCounter := 0
	expectedResult1 := "\nSorry, no matches on \"" + searchWord + "\"."
	actualResult1 := returnMatchMessage(intCounter, searchWord)
	if actualResult1 != expectedResult1 {
		t.Fail()
	}

	// Test 2
	intCounter = 3
	strCounter := strconv.Itoa(intCounter)
	expectedResult2 := "\n" + strCounter + " matches on \"" + searchWord + "\"."
	actualResult2 := returnMatchMessage(intCounter, searchWord)
	if actualResult2 != expectedResult2 {
		t.Fail()
	}
}

func TestAddStruct(t *testing.T) {
	allEmployeesSlice := csvutils.GetallEmployeesSlice("employees_test.csv")

	if len(allEmployeesSlice) != 3 {
		t.Fail()
	}

	allEmployeesSlice = addStruct(allEmployeesSlice, "f3", "l3", "d3", "pn3", "em3")
	if len(allEmployeesSlice) != 4 {
		t.Fail()
	}
}

func TestPrintAllNames(t *testing.T) {
	allEmployeesSlice := csvutils.GetallEmployeesSlice("employees_test.csv")
	expectedResult := "[0] - f1 l1\n[1] - f2 l2\n[2] - f3 l3\n"
	actualResult := printAllNames(allEmployeesSlice)

	if expectedResult != actualResult {
		t.Fail()
	}
}

func TestPrintAllCategories(t *testing.T) {
	expectedResult := "\n[0] - First name\n[1] - Last name\n[2] - Department\n[3] - Phone number\n[4] - E-mail\n"
	actualResult := printAllCategories()
	if expectedResult != actualResult {
		t.Fail()
	}
}

func TestConvertStringToInt(t *testing.T) {
	s := "10"
	expectedResult := 10
	actualResult := convertStringToInt(s)
	if expectedResult != actualResult {
		t.Fail()
	}
}

func TestCangeEmployeeInfo(t *testing.T) {
	allEmployeesSlice := csvutils.GetallEmployeesSlice("employees_test.csv")
	expectedResult := "test"

	allEmployeesSlice = changeEmployeeInfo(allEmployeesSlice, "0", "test", 0)
	testEmp := allEmployeesSlice[0]
	actualResult := testEmp.FirstName

	fmt.Println(actualResult)
	fmt.Println(expectedResult)

	if expectedResult != actualResult {
		t.Fail()
	}
}

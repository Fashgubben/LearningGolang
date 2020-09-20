package main

import (
	"testing"

	"github.com/fashgubben/LearningGolang/contacts/csvutils"
	"github.com/fashgubben/LearningGolang/contacts/employees"
)

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

func TestGetallEmployeesSlice(t *testing.T) {

	expectationArr1 := [5]string{"f1", "l1", "d1", "p1", "e1"}
	expectationArr2 := [5]string{"f2", "l2", "d2", "p2", "e2"}
	allEmployeesSlice := csvutils.GetallEmployeesSlice("employees_test.csv")

	emp1 := allEmployeesSlice[0]
	emp2 := allEmployeesSlice[1]

	if emp1.FirstName != expectationArr1[0] {
		t.Fail()
	} else if emp1.LastName != expectationArr1[1] {
		t.Fail()
	} else if emp1.Department != expectationArr1[2] {
		t.Fail()
	} else if emp1.PhoneNumber != expectationArr1[3] {
		t.Fail()
	} else if emp1.Email != expectationArr1[4] {
		t.Fail()
	}

	if emp2.FirstName != expectationArr2[0] {
		t.Fail()
	} else if emp2.LastName != expectationArr2[1] {
		t.Fail()
	} else if emp2.Department != expectationArr2[2] {
		t.Fail()
	} else if emp2.PhoneNumber != expectationArr2[3] {
		t.Fail()
	} else if emp2.Email != expectationArr2[4] {
		t.Fail()
	}
}

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

func TestAddStruct(t *testing.T) {
	allEmployeesSlice := csvutils.GetallEmployeesSlice("employees_test.csv")

	if len(allEmployeesSlice) != 2 {
		t.Fail()
	}

	allEmployeesSlice = addStruct(allEmployeesSlice, "f3", "l3", "d3", "pn3", "em3")
	if len(allEmployeesSlice) != 3 {
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

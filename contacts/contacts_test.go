package main

import (
	"testing"

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

func TestGetAllContacts(t *testing.T) {

	expectationArr1 := [5]string{"f1", "l1", "d1", "p1", "e1"}
	expectationArr2 := [5]string{"f2", "l2", "d2", "p2", "e2"}
	allEmployees := getAllContacts("employees_test.csv")

	// Index 1 & 2 because 0 is title
	emp1 := allEmployees[1]
	emp2 := allEmployees[2]

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

func TestAddStruct(t *testing.T) {
	allContacts := getAllContacts("employees_test.csv")

	if len(allContacts) != 3 {
		t.Fail()
	}

	allContacts = addStruct(allContacts, "f3", "l3", "d3", "pn3", "em3")
	if len(allContacts) != 4 {
		t.Fail()
	}
}

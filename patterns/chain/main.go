package main

import (
	"fmt"
	"main/pkg/chainPkg"
)

func main() {

	cashier := &chainPkg.Cashier{}

	//Set next for medical department
	medical := &chainPkg.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &chainPkg.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &chainPkg.Reception{}
	reception.SetNext(doctor)

	patient := &chainPkg.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
	fmt.Printf("Status: %v \n", *patient)
}

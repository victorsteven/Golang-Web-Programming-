package main

import "fmt"

//An interface allow polymorphism
type human interface {
	//what this means is that both the "person" and the "secretAgent" implement the "human" interface, because of the speak method that both have which is define below. We means both are of type human
	speak()
}

func saySomething(h human) {
	h.speak()
}

//Struct is a data structure
//we can make "fname" available outside this package by using "Fname"
type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func main() {

	m := map[string]int{
		"Ojo": 39,
		"Job": 20,
	}
	fmt.Println(m)
	fmt.Println(m["Ojo"])

	p1 := person{
		fname: "Steven",
		lname: "Mike",
	}
	// fmt.Println(p1)
	// fmt.Println(p1.fname)
	// p1.speak()

	sa1 := secretAgent{
		person{
			"Steven",
			"Victor",
		},
		false,
	}
	// sa1.person.speak()
	// sa1.speak()

	saySomething(p1)  //this will print out what "p1.speak()" will print
	saySomething(sa1) //this will print out what "sa1.speak()" will print

}

//the receiver is used to attach a function of some type
// func (receiver) identifier(parameter) (returns) { <code> }

//this function could as well be defined inside the "person" type when defining it, but we chose to do it here
//we call the function using the "person" struct assignment dot the function name.
func (p person) speak() {
	fmt.Println(p.fname, `says, "Good morning Steve."`)
}

func (sa secretAgent) speak() {
	//lname and fname got promoted that is why we could access them directly
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

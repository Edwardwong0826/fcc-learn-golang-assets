package main

import (
	"fmt"
)

// & operator generates a pointer
// * dereferences a pointer to gain access to the value
// this is method, a method is just a function with a receiver argument
// type struct is similiar to class in java, so this setMessage is the method on its\
// for method, the pointer receiver is more common than value receiver
func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

// don't edit below this line

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func test(e *email, newMessage string) {
	fmt.Println("-- before --")
	e.print()
	fmt.Println("-- end before --")
	e.setMessage("this is my second draft")
	fmt.Println("-- after --")
	e.print()
	fmt.Println("-- end after --")
	fmt.Println("==========================")
}

func (e email) print() {
	fmt.Println("message:", e.message)
	fmt.Println("fromAddress:", e.fromAddress)
	fmt.Println("toAddress:", e.toAddress)
}

func main() {
	test(&email{
		message:     "this is my first draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my second draft")

	test(&email{
		message:     "this is my third draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my fourth draft")

}

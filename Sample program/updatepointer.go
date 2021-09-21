package main

import "fmt"

type contactInfo struct {
	name    string
	rollno  int
	pincode int
}

func main() {
	alex := contactInfo{
		name:    "Logesh",
		rollno:  46,
		pincode: 638502,
	}
	jimpointer := &alex
	jimpointer.updateName("Logeshwaran")
	alex.print()

}

func (pointerToperson *contactInfo) updateName(updatename string) {
	(pointerToperson).name = updatename

}

func (p contactInfo) print() {

	fmt.Printf("%+v", p)
}

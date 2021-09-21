package main

import "fmt"

type person struct {
	fname string
	lanme string
}

type employee struct {
	id     int
	person person
}

type human interface {
	speak()
}

func (e employee) speak() {
	fmt.Print(e.id)
}
func (p person) speak() {
	fmt.Print(p.fname)
}
func bot(h human) {
	fmt.Print(h)
}

func main() {
	s := employee{id: 7, person: person{fname: "Logesh", lanme: "N"}}
	s.speak()
	p := person{fname: "Loki", lanme: "N"}
	p.speak()

	bot(s)

}

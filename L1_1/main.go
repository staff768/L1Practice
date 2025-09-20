package main

import "fmt"


type Human struct {
	Name string
	LastName string
	Age int
	Height, Weight int
}

func NewHuman(name string, last string, age int, height, weight int) *Human{
	return &Human{
		Name: name,
		LastName: last,
		Age: age,
		Height: height,
		Weight: weight,
	}
}

func (h *Human) SayHello(){
	fmt.Printf("Hello, my name is %s %s, I am %d years old. My height is %d sm and weight is %dkg\n",h.Name,h.LastName,h.Age,h.Height,h.Weight)
}

type Action struct {
	Human
	Street string
} 
func (a * Action) SayHello() {
	fmt.Printf("Hello, my name is %s %s, I am %d years old. My height is %d sm and weight is %dkg and my addr is %s\n",a.Name,a.LastName,a.Age,a.Height,a.Weight,a.Street)
}

func main() {
	Tom := NewHuman("Tom","Bobson",23,172,60)
	act := Action{
		Human: Human{
			Name: "Bib",
			LastName: "Porkson",
			Age: 12,
			Height: 160,
			Weight: 102,
		},
		Street: "London CapeTown",
	}
	act.SayHello()
	Tom.SayHello()
	act.Human.SayHello()
}
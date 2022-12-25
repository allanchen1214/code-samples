package main

import "fmt"

type Vehicle interface {
	Alert() string
}

type Car struct{}

func (c Car) Alert() string {
	return "Honk! Honk!"
}

type Bicycle struct{}

func (b Bicycle) Alert() string {
	return "Ring! Ring!"
}

func main() {
	c := Car{}
	b := Bicycle{}

	vehicles := []Vehicle{c, b}
	for _, v := range vehicles {
		fmt.Println(v.Alert())
	}
}

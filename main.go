package main

import "fmt"

type Component interface {
	getPrice() float64
}

type ConcreteMethod struct {
}

func (s *ConcreteMethod) getPrice() float64 {
	return 3.0
}

type Decorator struct {
	additional Component
}

func (p *Decorator) getPrice() float64 {
	totalPrice := p.additional.getPrice()
	return totalPrice + 3
}

type Decorator2 struct {
	additional2 Component
}

func (p *Decorator2) getPrice() float64 {
	totalPrice := p.additional2.getPrice()
	return totalPrice + 4
}

func main() {

	something := &ConcreteMethod{}

	
	addSomething := &Decorator2{
		additional2: something,
	}

	
	addSomething2 := &Decorator{
		additional: addSomething,
	}

	fmt.Println(addSomething2.getPrice())
}

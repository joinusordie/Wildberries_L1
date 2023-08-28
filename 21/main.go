package main

import "fmt"

type Target interface {
	Request() string
}

type Adaptee struct {
}

func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

func (a *Adaptee) SpecificRequest() string {
	return "Request"
}

type Adapter struct {
	*Adaptee
}

func (a *Adapter) Request() string {
	return a.SpecificRequest()
}

func main() {

	adapter := NewAdapter(&Adaptee{})

	fmt.Println(adapter.Request())
}

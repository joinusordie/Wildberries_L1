package main

import "fmt"

type Human struct {
	Name string
	Age  int
}
type Action struct {
	Human
	ActionName string
}

func main() {
	actionHuman := Action{
		Human{"Yaroslave", 23}, "Play",
	}

	fmt.Printf("%s performs a %s action", actionHuman.GetName(), actionHuman.ActionName)
}

func (h *Human) GetName() string {
	return h.Name
}
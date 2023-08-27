package main

import (
	"fmt"
	"reflect"
)

func typeof2(t interface{}) string {
	return reflect.TypeOf(t).String()
}

func typeof1(t interface{}) string {
	switch t.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan string:
		return "chan int"
	default:
		return "unknown"
	}
}

func main() {
	var a int
	var b float64
	var c string
	var d bool
	var e chan int

	fmt.Println("Ver. 1")
	fmt.Println(typeof1(a))
	fmt.Println(typeof1(b))
	fmt.Println(typeof1(c))
	fmt.Println(typeof1(d))
	fmt.Println(typeof1(e))

	fmt.Println("Ver. 2")
	fmt.Println(typeof2(a))
	fmt.Println(typeof2(b))
	fmt.Println(typeof2(c))
	fmt.Println(typeof2(d))
	fmt.Println(typeof2(e))
}

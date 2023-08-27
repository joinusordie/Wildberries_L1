package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var number int64
	var bit int
	fmt.Println("Select the number")
	fmt.Fscan(os.Stdin, &number)
	fmt.Println("Select the bit")
	fmt.Fscan(os.Stdin, &bit)

	val := fmt.Sprintf("%064b", number)
	fmt.Println(val)
	bits := strings.Split(val, "")

	if bits[len(bits)-bit] == "0" {
		bits[len(bits)-bit] = "1"
	} else {
		bits[len(bits)-bit] = "0"
	}
	val = strings.Join(bits, "")
	fmt.Println(val)
	res, _ := strconv.ParseInt(val, 2, 64)
	fmt.Println(res)

}

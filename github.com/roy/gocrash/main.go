package main

import (
	"fmt"
	"math"

	"go_workspace/github.com/roy/gocrash/util"
)

func main() {
	//var name string = "roy"
	name := "roy"
	var age = 37
	var isCool bool = true
	fmt.Println(name, age, isCool)
	//fmt.Printf("%T \t %T \t %T", name, age, isCool)

	//var vf = math.Floor(2.7)
	fmt.Println(math.Floor(2.7))

	fmt.Println(util.Reverse("hello"))
}

package main

import "fmt"

type Int int32
type String string

func main() {
	var i Int = 5
	fmt.Println(i.whoAreYou())

	var s String = "heyy, naber?"
	fmt.Println(s.whoAreYou())
}

//func (i int) whoAreYou() string {
//	return "I am a number nd my value is " + fmt.Sprintf(i)
//}
//
//func (s string) whoAreYou() string {
//	return "I am a number nd my value is " + fmt.Sprintf(i)
//}

func (i Int) whoAreYou() string {
	return "I am a number and my value is " + fmt.Sprintf("%d.\n", i)
}

func (s String) whoAreYou() string {
	return "I am a number and my value is " + string(s) + ".\n"
}

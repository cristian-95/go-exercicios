package main

import (
	"fmt"
	"unicode/utf8"
)

func f1() {
	str := "a"
	for i := 0; i < 100; i++ {
		fmt.Println(str)
		str += "a"
	}
}
func f2() {
	str := []byte("asSASA ddd dsjkdsjs dkᅖ")
	fmt.Println(len(str))
}

func f2b() {
	str := []byte("asSASA ddd dsjkdsjs dkᅖ")
	l := 0
	for _, value := range str {
		l += utf8.RuneLen(rune(value))
	}
	fmt.Println(l)
}

func f3() {
	str := "asSASA ddd dsjkdsjs dkᅖ"
	abc := []rune{'a', 'b', 'c'}
	str2 := []rune(str)
	str2[4] = abc[0]
	str2[5] = abc[1]
	str2[6] = abc[2]

	fmt.Println(string(str2))
}
func f4() {
	str := "foobar"
	inversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		inversed += string(str[i])
	}
	fmt.Println(inversed)
}

func main() {
	f4()
}

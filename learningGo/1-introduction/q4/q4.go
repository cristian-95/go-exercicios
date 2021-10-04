/*
Q4. (1) Strings
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

/*	1. Create a Go program that prints the following (up to 100 characters):
A
AA
AAA
AAAA
AAAAA
AAAAAA
AAAAAAA
...*/
func f1() {
	str := "a"
	for i := 0; i < 100; i++ {
		fmt.Println(str)
		str += "a"
	}
}

/*	2. Create a program that counts the number of characters in this string:
asSASA ddd dsjkdsjs dk*/
func f2() {
	str := []byte("asSASA ddd dsjkdsjs dkᅖ")
	fmt.Println(len(str))
}

/*	2b. In addition, make it output the number of bytes in that string. Hint: Check out the
utf8 package.*/
func f2b() {
	str := []byte("asSASA ddd dsjkdsjs dkᅖ")
	l := 0
	for _, value := range str {
		l += utf8.RuneLen(rune(value))
	}
	fmt.Println(l)
}

/*	3. Extend/change the program from the previous question to replace the three runes
at position 4 with ’abc’. */
func f3() {
	str := "asSASA ddd dsjkdsjs dkᅖ"
	abc := []rune{'a', 'b', 'c'}
	str2 := []rune(str)
	str2[4] = abc[0]
	str2[5] = abc[1]
	str2[6] = abc[2]

	fmt.Println(string(str2))
}

/*	4. Write a Go program that reverses a string, so “foobar” is printed as “raboof”. Hint:
You will need to know about conversion; skip ahead to section “Conversions” on
page 56.*/
func f4() {
	str := "foobar"
	inversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		inversed += string(str[i])
	}
	fmt.Println(inversed)
}

func main() {
	// f1()
	// f2()
	// f2b()
	// f3()
	f4()
}

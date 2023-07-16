package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "This is a single, line string"
	str2 := `This is a  multiple
	line string and this feature is going
	to be make a fun of golang in programming
	community.`

	fmt.Println(str1)
	fmt.Println(str2)

	str := str1 + "\t" + str2

	fmt.Println(str)

	if str1 == str2 {
		fmt.Println("Both str1 and str2 are equal")
	} else {
		fmt.Println("str1 and str2 are not equal")
	}

	if strings.Contains(str, str1) {
		fmt.Println("contains success!!")
	}

	fmt.Println("Split - ", strings.Split(str1, ","))

	fmt.Println("Replace - ", strings.Replace(str1, "is", "was", 1))

	newStr := "ankesh"
	fmt.Println(len(newStr), newStr[0])
	fmt.Println(65 + 26)

	fmt.Println(strings.Index("Ankesh is", "s"))
}

func IndexOf(c string) int {
	// Accepts character as argument, Returns index of
	// the first occurrence of specified character

	// str = "ankesh"

	return 0
}

package main

import (
	"fmt"
)

func is_vowel(char rune) bool {

	if (char == 'a') || (char == 'e') || (char == 'i') ||
		(char == 'o') || (char == 'u') {

		return true

	} else {

		return false

	}

}

func count_vowels(str string) string {
	var str1 string
	for _, char := range str {
		if is_vowel(char) {
			str1 += string(char)
		}
	}
	return str1
}

func check(s string) string {
    m := make(map[rune]int, len(s)) 
	fmt.Println("m:" , m)
    for _, r := range s {
        m[r]++
fmt.Println("hdfdjkdfffffffdfdfkjdfjkkjfdkjdkjdfjkdfkjfd:" , m[r])
    }
fmt.Println("String.................:" , s)
    for _, r := range s {
fmt.Println("String.................:" , s)
	fmt.Println("hdfdjkdfffffffdfdfkjd:" , m)
        if m[r] == 1 {
            return string(r)
        }
    }
    return ""
}

func main() {

	x := count_vowels("Hello Goo in delhi beach")

	w2 := check(x)

	fmt.Println("w2:" , w2)

}
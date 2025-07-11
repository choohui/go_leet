package main

/*

 */

import (
	"fmt"
)

func main() {
	input1 := "aaa"
	input2 := "a*a"
	fmt.Println(isMatch(input1, input2))
}

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	firstMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}

/*
func isMatch(s string, p string) bool {
	i := 0
	var instance rune
	for _, v := range s {
		if i >= len(p) {
			return false
		}
		if p[i] == '*' {
			if instance == v {
				continue
			} else if instance == '.' {
				if unicode.IsLetter(v) {
					continue
				} else {
					return false
				}
			} else {
				if i < len(p)-1 {
					i++
					if p[i] == byte(v) {
						i++
						instance = v
					} else {
						return false
					}
				} else {
					return false
				}
			}
		} else if p[i] == '.' {
			i++
			if unicode.IsLetter(v) {
				continue
			} else {
				return false
			}
		} else {
			if p[i] == byte(v) {
				i++
				instance = v
			} else {
				return false
			}
		}
	}
	return true
}
*/

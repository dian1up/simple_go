package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	input, hasil := short("osama")
	fmt.Println("Input = ", input)
	fmt.Println("Hasil = ", hasil)
}

func short(x string) (string, string) {
	abjad := "abcdefghijklmnopqrstuvwxyz"
	hidup := ""
	mati := ""
	apha := regexp.MustCompile("[aiueo]")
	x = strings.ToLower(x)

	for i := 0; i < len(abjad); i++ {
		for j := 0; j < len(x); j++ {
			if string(abjad[i]) == string(x[j]) {
				if apha.MatchString(string(x[j])) {
					hidup += string(abjad[i])
				} else {
					mati += string(abjad[i])
				}

			} else {
				continue
			}
		}
	}

	return x, hidup + mati
}

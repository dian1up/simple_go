package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	input, hidup, mati := count("omama")
	fmt.Println("Input = ", input)
	fmt.Println("Huruf Hidup = ", hidup)
	fmt.Println("Huruf Mati = ", mati)
}

func count(x string) (string, int, int) {
	abjad := "abcdefghijklmnopqrstuvwxyz"
	urutan := ""
	hidup := 0
	mati := 0
	apha := regexp.MustCompile("[aiueo]")
	x = strings.ToLower(x)

	for i := 0; i < len(abjad); i++ {
		for j := 0; j < len(x); j++ {
			if string(abjad[i]) == string(x[j]) {
				if apha.MatchString(string(x[j])) {
					urutan += string(abjad[i])
					hidup = hidup + 1
					break
				} else {
					urutan += string(abjad[i])
					mati = mati + 1
				}

			} else {
				continue
			}
		}
	}
	return x, hidup, mati
}

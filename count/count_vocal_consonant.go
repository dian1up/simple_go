package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	hidup, mati := urut("omama")
	fmt.Println("Huruf Hidup = ", hidup)
	fmt.Println("Huruf Mati = ", mati)
}

func urut(x string) (int, int) {
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
	return hidup, mati
}

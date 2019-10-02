package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type countStruct struct {
	Hidup int
	Mati  int
}

type shortStruct struct {
	Hasil string
}

func main() {
	http.HandleFunc("/count", CountControllers)
	http.HandleFunc("/short", ShortControllers)
	fmt.Println("run server pada port 8080")
	http.ListenAndServe(":8080", nil)
}

func ShortControllers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		input := r.FormValue("Input")
		hasil := ShortModel(input)
		data := []shortStruct{
			shortStruct{hasil},
		}
		result, err := json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func CountControllers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		input := r.FormValue("Input")
		hidup, mati := CountModel(input)
		data := []countStruct{
			countStruct{hidup, mati},
		}
		result, err := json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func CountModel(x string) (int, int) {
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

func ShortModel(x string) string {
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

	return hidup + mati
}

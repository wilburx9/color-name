package main

import "testing"

func TestNormalized(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"333", "FF333333"},
		{"#333", "FF333333"},
		{"FFFF", "FFFFFFFF"},
		{"#FFFF", "FFFFFFFF"},
		{"000000", "FF000000"},
		{"#000000", "FF000000"},
		{"AABBCCDD", "AABBCCDD"},
		{"#AABBCCDD", "AABBCCDD"},
	}

	for _, test := range tests {
		s, _ := normalize(test.input)
		if got := s; got != test.want {
			t.Errorf(`normalize("%s") == %s`, test.input, got)
		}
	}

}


func TestNotNormalized(t *testing.T) {
	var tests = []string {
		"2", "33", "#EE", "66666", "#4342f", "544ZiTotr", "09876tghf",
	}

	for _, test := range tests {
		s, e := normalize(test)
		if e == nil {
			t.Errorf(`normalize("%s") == %s`, test, s)
		}
	}

}

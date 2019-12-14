package main

import (
	"image/color"
	"math"
	"testing"
)

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
	var tests = []string{
		"2", "33", "#EE", "66666", "#4342f", "544ZiTotr", "09876tghf",
	}

	for _, test := range tests {
		s, e := normalize(test)
		if e == nil {
			t.Errorf(`normalize("%s") == %s`, test, s)
		}
	}

}

func TestRGBParsedToValidToHSL(t *testing.T) {
	var tests = []struct {
		input   color.RGBA
		expects HSL
	}{
		{color.RGBA{R: 255, G: 255, B: 255}, HSL{H: 0, S: 0, L: 100}}, // White
		{color.RGBA{R: 0, G: 0, B: 0}, HSL{H: 0, S: 0, L: 0}},         // Black
		{color.RGBA{R: 255, G: 0, B: 0}, HSL{H: 0, S: 100, L: 50}},    // Red
		{color.RGBA{R: 0, G: 255, B: 0}, HSL{H: 120, S: 100, L: 50}},  // Lime
		{color.RGBA{R: 0, G: 0, B: 255}, HSL{H: 240, S: 100, L: 50}},  // Blue
		{color.RGBA{R: 255, G: 255, B: 0}, HSL{H: 60, S: 100, L: 50}}, // Yellow
		{color.RGBA{R: 0, G: 0, B: 128}, HSL{H: 240, S: 100, L: 25}},  // Navy
		{color.RGBA{R: 128, G: 0, B: 0}, HSL{H: 0, S: 100, L: 25}},    // Maroon
		{color.RGBA{R: 191, G: 191, B: 191}, HSL{H: 0, S: 0, L: 75}},  // Sliver
	}

	const minDiff = 0.5

	eq := func(a HSL, b HSL) bool {
		if math.Abs(a.H-b.H) > minDiff {
			return false
		}
		if math.Abs(a.S-b.S) > minDiff {
			return false
		}
		if math.Abs(a.L-b.L) > minDiff {
			return false
		}
		return true
	}

	for _, test := range tests {
		if hsl := toHsl(test.input); !eq(hsl, test.expects) {
			t.Errorf(`toHsl("%+v") == %+v. Expects %+v`, test.input, hsl, test.expects)
		}
	}
}


func TestValidStringColorsParsedToRGBA(t *testing.T) {
	var tests = []string{
		"FFFFFF", "333333", "EEEEEE", "66666", "4342FF", "544092", "098765",
	}

	for _, test := range tests {
		_, e := strToRGBA(test)
		if e != nil {
			t.Errorf(`strToRGBA("%v") == error: %v`, test, e)
		}
	}
}

func TestInValidStringColorsNotParsedToRGBA(t *testing.T) {
	var tests = []string{
		"FFFF", "#333333", "0987657EEEEEE", "iuewkkd", "iwd23uiw", "4311194", "09yhh9392-ir",
	}

	for _, test := range tests {
		v, e := strToRGBA(test)
		if e == nil {
			t.Errorf(`strToRGBA("%v") == %v`, test, v)
		}
	}
}
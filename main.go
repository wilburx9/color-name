package main

import (
	"bytes"
	"flag"
	"fmt"
	"image/color"
	"math"
	"os"
	"strconv"
	"strings"
)

type HSL struct {
	H, S, L float64
}

const (
	warningColor = "\033[33m%s\033[0m\n"
)

func main() {
	hexPtr := flag.String("h", "",
		"The hex value of the color whose colorName you want. Accepted formats: FFF, #FFF, FFFF, #FFFF, FFFFFF, #FFFFFF, FFFFFFFF, #FFFFFFFF")
	flag.Parse()

	c := *hexPtr

	if c == "" {
		fmt.Print("Usage: color-name -h 000")
		os.Exit(1)
		return
	}
	normalized, err := normalize(c)
	if err != nil {
		fmt.Printf("%v is not a supported format", c)
		os.Exit(1)
		return
	}

	rgb, err := strToRGBA(normalized)
	if err != nil {
		fmt.Printf("%v Could not parse color to RGB", c)
		os.Exit(1)
		return
	}

	item, exactMatch, err := colorName(normalized, rgb)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
		return
	}

	if exactMatch {
		fmt.Println(item.name)
	} else {
		fmt.Printf(warningColor, item.name)
	}
}


func normalize(color string) (string, error) {

	// Remove leading '#'
	color = strings.TrimPrefix(color, "#")

	errMsg := fmt.Errorf("#%v appears to be an invalid color\n", color)

	bytesLen := len(color)
	if bytesLen < 3 || bytesLen > 8 || bytesLen == 5 {
		return "", errMsg
	}

	// Converting the passed hex to uppercase
	color = strings.ToUpper(color)

	i := len(color)
	if i == 6 {
		return color, nil
	}
	var buffer bytes.Buffer

	repeat := func() {
		for _, i := range color {
			str := fmt.Sprintf("%c", i)
			buffer.WriteString(strings.Repeat(str, 2))
		}
	}

	switch i {
	case 3:
		repeat()
	case 4:
		color = color[1:]
		repeat()
	case 8:
		for i := 2; i < 7; i += 2 {
			buffer.WriteString(color[i : i+2])
		}
	}

	str := buffer.String()
	if str == "" {
		return "", errMsg
	}
	return str, nil
}

func rgbToHsl(rgba color.RGBA) HSL {
	r, g, b := float64(rgba.R), float64(rgba.G), float64(rgba.B)
	r /= 255
	g /= 255
	b /= 255
	min := math.Min(r, math.Min(g, b))
	max := math.Max(r, math.Max(g, b))
	delta := max - min
	l := (min + max) / 2

	var s float64
	if l > 0 && l < 1 {
		if l < 0.5 {
			s = delta / (2 * l)
		} else {
			s = delta / (2 - 2*l)
		}
	}

	var h float64
	if delta > 0 {
		if max == r && max != g {
			h += (g - b) / delta
		}
		if max == g && max != b {
			h += 2 + (b-r)/delta
		}
		if max == b && max != r {
			h += 4 + (r-g)/delta
		}
		h /= 6
	}
	return HSL{
		H: h * 255,
		S: s * 255,
		L: l * 255,
	}
}

func strToRGBA(str string) (color.RGBA, error) {
	rStr := fmt.Sprintf("0x%v", str[0:2])
	gStr := fmt.Sprintf("0x%v", str[2:4])
	bStr := fmt.Sprintf("0x%v", str[4:])

	r, err := strconv.ParseUint(rStr, 0, 8)
	if err != nil {
		return color.RGBA{}, err
	}

	g, err := strconv.ParseUint(gStr, 0, 8)
	if err != nil {
		return color.RGBA{}, err
	}

	b, err := strconv.ParseUint(bStr, 0, 8)
	if err != nil {
		return color.RGBA{}, err
	}

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
	}, nil
}

func colorName(str string, rgb color.RGBA) (item, bool, error) {
	var hsl = rgbToHsl(rgb)
	var ndf, ndf1, ndf2 float64
	var cl = -1
	var df float64 = -1
	for i, v := range colorItems {
		if v.color == str {
			return v, true, nil
		}

		rbg2, _ := strToRGBA(v.color)
		hsl2 := rgbToHsl(rbg2)

		ndf1 = math.Pow(float64(rgb.R-rbg2.R), 2) +
			math.Pow(float64(rgb.G-rbg2.G), 2) +
			math.Pow(float64(rgb.B-rbg2.B), 2)

		ndf2 = math.Pow(hsl.H-hsl2.H, 2) +
			math.Pow(hsl.S-hsl2.S, 2) +
			math.Pow(hsl.L-hsl2.L, 2)

		ndf = ndf1 + (ndf2 * 2)
		if df < 0 || df > ndf {
			df = ndf
			cl = i
		}
	}

	if cl < 0 {
		return item{}, false, fmt.Errorf("#%s is an invalid color", str)
	}

	return colorItems[cl], false, nil
}

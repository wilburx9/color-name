package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

}


func normalize(color string) (string, error) {

	// Remove leading '#'
	color = strings.TrimPrefix(color, "#")

	// Converting the passed hex to uppercase
	color = strings.ToUpper(color)

	i := len(color)
	if i == 8 {
		return color, nil
	}
	var buffer bytes.Buffer

	pad := func() {
		for _, i := range color {
			str := fmt.Sprintf("%c", i)
			buffer.WriteString(strings.Repeat(str, 2))
		}
	}

	prepend := func() {
		buffer.WriteString("FF")
	}
	switch i {
	case 3:
		prepend()
		pad()
	case 4:
		pad()
	case 6:
		prepend()
		buffer.WriteString(color)
	}

	str := buffer.String()
	if str == "" {
		return "", fmt.Errorf("#%v appears to be an invalid colorStr\n", color)
	}
	return str, nil
}

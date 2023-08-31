package format

import (
	"reflect"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func NumWithSpaces[T ~int | ~uint | ~int64 | ~float32 | ~float64](number T) string {
	rounding := 2

	p := message.NewPrinter(language.Russian)

	var numberWithSpaces string

	if number == T(int(number)) {
		rounding = 0
	}

	switch reflect.TypeOf(number).String() {
	case "int", "int64", "uint":
		numberWithSpaces = p.Sprintf("%d", number)
	case "float32", "float64":
		if rounding == 0 {
			numberWithSpaces = p.Sprintf("%.0f", number)
		} else {
			numberWithSpaces = p.Sprintf("%.2f", number)
		}
	}

	return numberWithSpaces
}

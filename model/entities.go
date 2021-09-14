package model

import (
	"errors"
	"strconv"
	"unicode"
)

type Result struct {
	Type   string
	Length int
	Value  string
}

func NewResult(c string) (Result, error) {
	if len(c) >= 4 {
		resultado := SeteoValores(c)
		if checkLongitud(resultado) && checkType(resultado) {
			return resultado, nil
		}
	}
	return Result{}, errors.New("cadena invalida")
}
func checkLongitud(resultado Result) bool {
	return len(resultado.Value) == resultado.Length
}
func checkType(resultado Result) bool {
	if resultado.Type == "NN" {
		return isInt(resultado.Value)
	} else {
		return isLetter(resultado.Value)
	}
}
func SeteoValores(c string) Result {
	tipo := c[0:2]
	largo := c[2:4]
	valor := c[4:]
	largoParseado, _ := strconv.ParseInt(largo, 0, 8)
	return Result{tipo, int(largoParseado), valor}
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func isLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

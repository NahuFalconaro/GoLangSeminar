package main

import (
	"EntregableSeminarioGo/model"
	"fmt"
)

func convertirCadena(c string) (model.Result, error) {
	retorno, err := model.NewResult(c)
	return retorno, err
}
func main() {
	cadena := "TX03ABC"
	result, err := convertirCadena(cadena)
	if err == nil {
		fmt.Println(result)
	} else {
		panic(err)
	}
}

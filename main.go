package main

import (
	"EntregableSeminarioGo/model"
	"fmt"
)

func main() {
	cadena := "TX03ABC"
	result, err := model.NewResult(cadena)
	if err == nil {
		fmt.Println(result)
	} else {
		panic(err)
	}
}

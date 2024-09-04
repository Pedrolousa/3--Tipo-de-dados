package main

import ("fmt"
		"errors"
)

func main() {
	var numero int64 = 1000000
	fmt.Println(numero)

	var numero2 uint32 = 32000
	fmt.Println(numero2)
// unsygned int
//INT32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

// BYTE = UINT8
	var numero4 byte = 12
	fmt.Println(numero4)

	var numeroreal1 float32 = 123.45
	fmt.Println(numeroreal1)



	// FIM NUMEROS REAIS

	var str string = "texto"
	fmt.Println(str)

	str2:= "texto2"
	fmt.Println(str2)

	char := 'a'

	fmt.Println(char)

	// FIM STRINGS

	var texto string
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano bool
	fmt.Println(booleano)

	var erro error = errors.New("erro interno")
	fmt.Println(erro)

	




}
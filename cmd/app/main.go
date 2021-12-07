package main

import (
	"fmt"
	"go_compiler/internal/lexer"
)

func main(){
	j := lexer.Lexer("int x := 12;")
	for i := 0; i < len(j); i++{
		fmt.Println(j[i].Token, "_", j[i].Word, "_")
	}
}
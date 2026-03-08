package main

import (
	"fmt"
	"go_for_dummies_borshch/utils"
	"strings"
)

func main() {
{	
	fmt.Println("Hellow")
	utils.Test()
}

{
	fmt.Println("Повторить символ i раз")
	for i := range 8 {
	x := strings.Repeat("*", i)
	fmt.Println(x)
	}
}
}
package main

import (
	"fmt"
	"unsafe"
	// "bufio"
	// "encoding/csv"
	// "encoding/json"
	// "go_for_dummies_borshch/utils"
	// "maps"
	// "os"
	// "slices"
	// "sort"
	// "strings"
	// "unicode/utf8"
)

const (
	monday  = iota + 1
	tuesday = iota * 2
	wednesday
	thursday

)

const (
	_    = 6
	A, _ = iota, iota + 10
	_, _
	_, B
)

func main() {
	fmt.Println("const и iota")
{
	fmt.Println(monday, tuesday, wednesday, thursday)
	fmt.Println(A, B)
}

{
	var a int64 = 100  
	b := 33.5 
	c := ""
	fmt.Printf("a (int64): %d байт\n", unsafe.Sizeof(a))
	fmt.Printf("b (float64): %d байт\n", unsafe.Sizeof(b))
	fmt.Printf("c (string): %d байт\n", unsafe.Sizeof(c))
}

fmt.Println("Определение типа переменной")

{
	name := "Alice"
	fmt.Printf("%T", name) // string
	fmt.Println()
}

fmt.Println("Округление")

{
    var n float64 = 1.1234
    
    fmt.Printf("%.2f\n", n)
}



}
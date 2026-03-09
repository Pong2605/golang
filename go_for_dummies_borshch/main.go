package main

import (
	"fmt"
	"go_for_dummies_borshch/utils"
	"strings"
	// "bufio"
	// "os"
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

{
	fmt.Println("Повторить символ i раз")
	for i := 1; i < 8; i++ {
	x := strings.Repeat("*", i)
	fmt.Println(x)
	}
}

{
	var someText string = ""
    fmt.Println("some text: ", someText)

    someText = "new text"
    fmt.Println("some text: ", someText)

    moreText := someText
    fmt.Println("more text: ", moreText)
}

// 5.3 Ввод пользовательских данных в консоли
fmt.Println("5.3 Ввод пользовательских данных в консоли")
fmt.Println("bufio.NewReader и ReadString")

{

	// Чтобы считать всю строку целиком, игнорируя пробелы, нам нужно использовать другой способ. 
	// Вот тут на помощь приходит пакет bufio, который предоставляет более гибкие инструменты для чтения пользовательского ввода

	// reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Введите предложение:")

	// Считываем строку до символа первого нажатия Enter
	// input, _ := reader.ReadString('\n')

	input := "test"

	// Выводим считанную строку
	fmt.Println("Вы ввели:", input)
}

fmt.Println("Функция вывода Printf")

{
	a, b, c := "One", "Two", "Three"
	fmt.Printf("%s\n%s\n%s\n", a, b, c)
}


// 6.1 Арифметические операции над числами
fmt.Println("6.1 Арифметические операции над числами")
fmt.Println("Преобразование типов")

{
    var a int = 10
    var b float64 = 3.5

    // Преобразуем целое число `а` в число с плавающей точкой
    sum := float64(a) + b
    difference := float64(a) - b
    product := float64(a) * b
    quotient := float64(a) / b
    
    fmt.Println("Cумма:", sum)
    fmt.Println("Разность:", difference)
    fmt.Println("Произведение:", product)
    fmt.Println("Частное:", quotient)	
}

{
    var x, y float32 = 4.0, 5.0
    fmt.Printf(`Сложение: %v
Вычитание: %v
Умножение: %v
Деление: %v`, x+y, x-y, x*y, x/y)
}

}


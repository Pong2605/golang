package main

import (
	"fmt"
	"go_for_dummies_borshch/utils"
	"strings"
	"unicode/utf8"
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

// 6.2 Операции над строками и символами
fmt.Println("6.2 Операции над строками и символами")
fmt.Println("Длина строки")

{
    text := "Этот текст состоит из 74 байт и 43 символов"
    var lengthInBytes int = len(text)
    var lengthInRunes int = utf8.RuneCountInString(text)
    fmt.Println(lengthInBytes, lengthInRunes)

// Если нужно вычислить длину строки, содержащей только английские буквы, символы и цифры, достаточно использовать функцию len(), так как каждый из этих символов занимает один байт.
// Если требуется узнать длину строки, в которой присутствуют буквы других языков, такие как русские или арабские, следует применять функцию utf8.RuneCountInString() из пакета "unicode/utf8".	
}

fmt.Println("Операции над типом rune")

{
    var r rune = 'B'   // 'B' имеет код 66 в Unicode
    fmt.Println(r + 1) // 'C' имеет код 67 в Unicode 
    fmt.Println(r - 1) // 'A' имеет код 65 в Unicode 
}

fmt.Println("Вывод rune как int32 и string")

{
    var r rune = 'A'  // 'A' имеет код 65 в Unicode

    fmt.Println(r)    // выведет 65 на консоль

    fmt.Println(string(r)) // выведет "B"
}

// 7.2 Конструкция switch
fmt.Println("7.2 Конструкция switch")

{
    const (
        Morning = 1
        Afternoon = 2
        Evening = 3
        Night = 4
    )

    timeOfDay := 2

    switch timeOfDay {
    case Morning:
        fmt.Println("Сейчас утро!")
    case Afternoon:
        fmt.Println("Сейчас обед!")
    case Evening:
        fmt.Println("Сейчас вечер!")
    case Night:
        fmt.Println("Сейчас ночь!")
    default:
        fmt.Println("Сейчас неизвестное время года...")
    }
}

}


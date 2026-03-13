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

fmt.Println("Калькулятор")

{
const (
		Addition        = "+"
		Subtraction     = "-"
		Multiplication  = "*"
		Division        = "/"
		Modulus         = "%"
	)

	var firstNumber, secondNumber float64
	var operation string

	// fmt.Scan(&firstNumber, &secondNumber, &operation)

    firstNumber, secondNumber, operation = 4, 2, "/"

	// Проверяем, является ли firstNumber целым числом
	if firstNumber != float64(int(firstNumber)) {
		fmt.Println("Первое число не является целым, операция остатка от деления (%) может работать некорректно.")
	}

	switch operation {
	case Addition:
		fmt.Println(firstNumber + secondNumber)
	case Subtraction:
		fmt.Println(firstNumber - secondNumber)
	case Multiplication:
		fmt.Println(firstNumber * secondNumber)
	case Division:
		if secondNumber == 0 {
			fmt.Println("Делить на ноль нельзя!")
		} else {
			fmt.Println(firstNumber / secondNumber)
		}
	case Modulus:
		// Проверяем, что оба числа фактически целые, перед выполнением %
		if firstNumber != float64(int(firstNumber)) || secondNumber != float64(int(secondNumber)) {
			fmt.Println("Для операции остатка от деления оба числа должны быть целыми.")
		} else if secondNumber == 0 {
			fmt.Println("Делить на ноль нельзя!")
		} else {
			fmt.Println(int(firstNumber) % int(secondNumber))
		}
	default:
		fmt.Println("Неподдерживаемая операция")
	}
}

// 9.1 Массивы
fmt.Println("9.1 Массивы")

{
    // 1. Объявление массива с нулевыми значениями
    var arr [5]int // Обьявили переменную arr типа: Массив чисел int с размером 5
    fmt.Println(arr)
}

{
    // 2. Объявление и инициализация значениями
    arr := [5]int{10, 20, 30, 40, 50} // Массив на 5 элементов, значения заданы явно
    fmt.Println(arr)
}

{
    // 3. Частичная инициализация
    arr := [5]int{1, 2} // Массив: [1, 2, 0, 0, 0]
    fmt.Println(arr)
}

{
    arr := [20]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3}

    max1 := arr[0]
    index1 := 0
    for i := 1; i < len(arr); i++ {
        if arr[i] > max1 {
            max1 = arr[i]
            index1 = i
        }
    }

    max2 := -1 << 31
    index2 := -1
    for i := 0; i < len(arr); i++ {

        if i != index1 && arr[i] > max2 {
            max2 = arr[i]
            index2 = i
        }
    }


    if index1 < index2 {
        fmt.Println(max1, max2)
    } else {
        fmt.Println(max2, max1)
    }
}

// 9.2 Срезы (slice)
fmt.Println("9.2 Срезы (slice)")

{
    // Динамическая структура данных — это структура, которая может изменять свой размер во время работы программы. 

    // Упрощенно:

    // Если данных мало — она использует меньше памяти.
    // Если данных много — она автоматически "растет", чтобы вместить всё, что нужно.


    // Слайсы можно создавать тремя способами:

    // 1. Создание пустого слайса

    var slice1 []int // Создан пустой слайс
    fmt.Println(slice1)
                    
    // 2. Создание слайса с начальными значениями:

    slice2 := []int{10, 20, 40, 50, 100} // Слайс с 5 элементами
    fmt.Println(slice2)
                    
    // 3. Создание пустого слайса через make (наиболее распространенный) потому-что:

    // Функция make позволяет задавать длину и емкость слайса:

    slice3 := make([]int, 3, 5) // Длина 3, емкость 5
    fmt.Println(slice3)

    // Функция make принимает три аргумента:

    // Тип слайса (первый аргумент): Тип слайса который вы хотите создать.
    // Длина (второй аргумент): Начальное количество элементов в слайсе.
    // Емкость (третий аргумент необязательный): Максимальное количество элементов, 
    // которые слайс может хранить без выделения новой памяти.
}

{
    slice := []int{10, 20, 30}

    slice = append(slice, 40)  // Добавляем элемент 40
    fmt.Println(slice)         // [10, 20, 30, 40]

    slice = append(slice, 100) // Добавляем элемент 100
    fmt.Println(slice)         // [10, 20, 30, 40, 100]
}

{
    slice := []int{}
    fmt.Println(slice) // []

    slice = append(slice, 10, 20, 30)
    fmt.Println(slice) // [10, 20, 30]

    slice = append(slice, 40, 100)
    fmt.Println(slice) // [10, 20, 30, 40, 100]
}

{
    slice := []int{1, 2, 3, 4}
    fmt.Println(slice)

    subSlice := slice[1:3]
    fmt.Println(subSlice)

    subSlice[1] = 5

    fmt.Println(subSlice)
}

{
    var slice []int = []int{1, 2, 3}
    fmt.Println(slice)

    slice2 := append(slice, 4, 5, 6)
    fmt.Println(slice)
    fmt.Println(slice2)

    slice2[0] = 100
    fmt.Println(slice)
    fmt.Println(slice2)
}

// 9.3 Конструкция for := range
fmt.Println("9.3 Конструкция for := range")

{
	// import (
	// 	"bufio"
	// 	"fmt"
	// 	"os"
	// )

	// input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	input := "Привет, мир!"
	for _, char := range input {
		fmt.Println(string(char))
	}
}

{
	// import (
	// 	"bufio"
	// 	"fmt"
	// 	"os"
	// )

	// input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input := "Hello, world!"

    var count int
    
    for _, char := range input {
        switch string(char){
            case "a", "e", "i", "o", "u":
            count++
        }
    }
    fmt.Print(count)
}

{
	// import (
	//     "bufio"
	//     "fmt"
	//     "os"
	//     "strings"
	// )


    // reader := bufio.NewReader(os.Stdin)
    // input, _ := reader.ReadString('\n')
	input := "Hello, world!"
	
    vowels := "aeiouAEIOU"
    count := 0

    for _, char := range input {
        if strings.ContainsRune(vowels, char) {
            count++
        }
    }

    fmt.Println(count)
}

// 10  Хеш-таблицы
fmt.Println("10  Хеш-таблицы")

{
    fmt.Println("Map (мапа)")
    
    books := map[string]string {
    "Гарри Поттер": "Дж. К. Роулинг",
    "Властелин Колец": "Дж. Р. Р. Толкин",
    "1984":  "Джордж Оруэлл",
    }
    fmt.Println(books["Властелин Колец"])  // Дж. Р. Р. Толкин

}

{
    // Как создавать map?
    // Создать map можно тремя способами:

    // 1. Создание пустого map c помощью make():

    books := make(map[int]string)
    fmt.Println(books)

    // 2. Создание map с начальными значениями: 

    isMale := map[string]bool {
    "Илья": true,
    "Вика": false,
    "Василий": true,
    }
    fmt.Println(isMale)

    // 3. Создание пустого map через make с заданной емкостью (редко используется):

    someMap := make(map[float32]float64, 10) // Создаем map с емкостью для 10 элементов
    fmt.Println(someMap)
}

{
    fmt.Println("Проверка на наличие ключа в мапе")

    database := map[int]string {
        1: "Illia Borshch",
        2: "Petya Pupkin",
    }

    name, ok := database[123]

    if ok {
        fmt.Println("Значение:", name)
    } else {
        fmt.Println("Ключ не найден!")  // это будет выведено
    }
}

{
    database := map[int]string {
        1: "Illia Borshch",
        2: "Petya Pupkin",
    }
    fmt.Println("Удаление пар в map")

    delete(database, 1)  // удалили Илью 🥲

    _, ok := database[1]
    if ok {
        fmt.Println("delete не сработала и Илья остался в базе данных!")
    } else {
        fmt.Println("delete сработал как надо!")
    }
}

{
    fmt.Println("Итерация по map")

    database := map[int]string{
        1: "Illia Borshch",
        2: "Petya Pupkin",
    }

    for id, name := range database {
        fmt.Printf("ID: %d, Name: %s\n", id, name)
    }
}

{
    slice1 := []int{}
    fmt.Println(slice1)
}
}


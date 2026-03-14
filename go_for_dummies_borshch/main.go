package main

import (
	// "bufio"
	// "os"
	"fmt"
	"go_for_dummies_borshch/utils"
	"maps"
	"slices"
	"sort"
	"strings"
	"unicode/utf8"
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
// Напишите программу, которая подсчитывает количество уникальных символов в строке (включая пробелы и знаки пунктуации).

// Программа должна:

// Прочитать строку, введённую пользователем.
// Найти все уникальные символы в строке.
// Вывести их в порядке возрастания ASCII-кодов.

    // scanner := bufio.NewScanner(os.Stdin)

	// Считываем строку от пользователя
	// scanner.Scan()
	// input := scanner.Text()

    input := "Hello, World!"

	// Создаём карту для хранения уникальных символов (ключ — rune, значение — bool)
	dict := make(map[rune]bool)

	// Проходим по всем символам строки и добавляем их в карту
	for _, char := range input {
		dict[char] = true
	}

	// Преобразуем ключи карты (символы) в срез для сортировки
	var chars []rune
	for char := range dict {
		chars = append(chars, char)
	}

	// Сортируем символы по ASCII‑коду (по возрастанию)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	// Выводим результат: символы через пробел
	for _, char := range chars {
		// if i > 0 {
		// 	fmt.Print(" ") // пробел между символами
		// }
		fmt.Printf("%c", char) // выводим символ как rune
	}
    fmt.Println()
}

{
    // reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
    input := "Hello, World!"

    dict := make(map[rune]int)
    
    for _, key := range input {
        dict[key]++
    }

    arr := slices.Sorted(maps.Keys(dict))
    
    for _, char := range arr {
        fmt.Print(string(char))
    }
    fmt.Println()    
}

{
    // dict := make(map[string]int)
    // var key string
    
    // for n, _ := fmt.Scan(&key); n > 0; n, _ = fmt.Scan(&key) {
    //     dict[key]++
    // }

    // arr := slices.Sorted(maps.Keys(dict))
    
    // fmt.Println(arr)

    // for _, el := range arr {
    //     fmt.Print(el, ": ", dict[el], "\n")
    // }
}

{
    // Проект – Конвертер валют
    // Курс валют относительно 1 USD
    var rates = map[string]float64{
        "USD": 1.0,   // Доллар США
        "EUR": 0.92,  // Евро
        "RUB": 90.0,  // Российский рубль
        "JPY": 157.0, // Японская иена
        "CNY": 7.25,  // Китайский юань
        "GBP": 0.78,  // Британский фунт
        "KZT": 460.0, // Казахстанский тенге
        "TRY": 32.5,  // Турецкая лира
        "INR": 83.0,  // Индийская рупия
        "BRL": 5.12,  // Бразильский реал
        "AUD": 1.50,  // Австралийский доллар
        "CAD": 1.36,  // Канадский доллар
        "CHF": 0.89,  // Швейцарский франк
        "SEK": 10.8,  // Шведская крона
        "NOK": 10.5,  // Норвежская крона
    }

    // генерируем слайс валют для удобства
    currencies := make([]string, 0)
    for currency := range rates {
        currencies = append(currencies, currency)
    }

    fmt.Println("Добро пожаловать в конвертер валют!")
    fmt.Println("Доступные валюты для конвертации:")
    for i, currency := range currencies {
        fmt.Printf("%d. %s\n", i+1, currency)
    }

    fmt.Println()

    var amount float64
    fmt.Print("Введите сумму в USD: ")
    // fmt.Scan(&amount)
    amount = 5.0

    if amount <= 0 {
        fmt.Println("Сумма должна превышать 0!")
    }

    fmt.Println("Выберите номер валюты для конвертации из списка выше:")
    var choice int
    // fmt.Scan(&choice)
    choice = 2

    if choice <= 0 || choice > len(currencies) {
        fmt.Println("Неправильный выбор валюты!")
    }

    selectedCurrency := currencies[choice-1]

    conversionRate := rates[selectedCurrency]
    convertedAmount := amount * conversionRate
    fmt.Printf("%.2f USD = %.2f %s\n", amount, convertedAmount, selectedCurrency) 
}

{
	// reader := bufio.NewReader(os.Stdin)
    input := "abcdaabcde"
	// Считываем строку от пользователя
	// input, _ := reader.ReadString('\n')
	input = input[:len(input)-1] // Убираем символ \n

	if len(input) == 0 {
		fmt.Println("0 ")
		return
	}

	charIndex := make(map[rune]int) // Карта: символ → последний индекс
	maxLen := 0                    // Максимальная длина подстроки
	maxStart := 0                 // Начальная позиция максимальной подстроки
	start := 0                    // Начало текущей подстроки без повторов


	for end, char := range input {
		// Если символ уже встречался И его последний индекс >= start
		if lastIndex, ok := charIndex[char]; ok && lastIndex >= start {
			start = lastIndex + 1 // Сдвигаем начало подстроки после последнего вхождения
		}

		charIndex[char] = end // Обновляем последний индекс символа


		// Проверяем, стала ли текущая подстрока длиннее максимальной
		currentLen := end - start + 1
		if currentLen > maxLen { // Строгое сравнение: только если длиннее
			maxLen = currentLen
			maxStart = start
		}
	}

	// Формируем результат
	longestSubstring := input[maxStart : maxStart+maxLen]
	fmt.Printf("%d %s\n", maxLen, longestSubstring)
}

// Функции с неограниченным количеством параметров
fmt.Println("Функции с неограниченным количеством параметров")

{
    // func sum(nums ...int) int {
    // total := 0
    // for _, num := range nums {
    //     total += num
    // }
    // return total


// func main() {
    result := utils.Sum(1, 2, 3, 4, 5)
    fmt.Println("Сумма:", result) // 15

    result = utils.Sum(1, 2, 3, 4, 5, 10, 100)
    fmt.Println("Сумма:", result) // 125
// }
}

fmt.Println("Распаковка слайса")

{
// func sum(numbers ...int) int {
//     total := 0
//     for _, num := range numbers {
//         total += num
//     }
//     return total
// }

// func main() {
    nums := []int{10, 20, 30}

    // result := sum(nums[0], nums[1], nums[2])  <- о таком ужасе мы забываем

    result := utils.Sum(nums...)        // Распаковка слайса nums
    fmt.Println("Сумма:", result) // 60
// }
}

fmt.Println("Тип данных any")

{
// Тип данных any – появился в Go 1.18 как новый универсальный тип. 
// Для знатоков – это просто псевдоним для уже существующего типа interface{}.
// Но что он означает? any буквально означает "любой тип данных". 
// Это позволяет функциям работать с данными любого типа без необходимости уточнять их на этапе компиляции.  

// func print(args ...any) {
//     for _, arg := range args {
//         fmt.Print(arg, " ")
//     }
// }

// func main() {
    utils.Print("Привет", 42, 3.14, true, []int{1, 2, 3})
// }

}

fmt.Println("Возврат нескольких значений")

{

// Функции с несколькими возвращаемыми значениями полезны в ситуациях, 
// когда нужно вернуть основной результат и дополнительную информацию, 
// например, статус выполнения или ошибку

// func divide(a, b int) (int, bool) {
//     if b == 0 {
//         return 0, false
//     }
//     return a / b, true
// }

// func main() {
    res, ok := utils.Divide(10, 2)
    if ok {
        fmt.Println(res)  // 5
    } else {
        fmt.Println("Ошибка!")
    }
// }
}

fmt.Println("Именованный возврат")

{
// С именованными возвращаемыми значениями мы можем явно задать имена для значений, которые функция возвращает. 
// Это позволяет упростить код, так как можно избежать явного указания возвращаемых переменных в return. 

// func findMax(numbers []int) (max int, ok bool) {
//     if len(numbers) == 0 {
//         max, ok = 0, false  // Присваиваем значения именованным возвращаемым переменным
//         return              // Возвращаем без указания переменных max и ok явно
//     }
    
//     max = numbers[0]   // max уже объявлен как возвращаемая переменная
//     for _, num := range numbers {
//         if num > max {
//             max = num
//         }
//     }
//     ok = true  // Устанавливаем флаг
//     return     // Возвращаем автоматически
// }

// func main() {
    nums := []int{3, 5, 7, 2, 8}
    max, ok := utils.FindMax(nums)
    
    if ok {
        fmt.Println(max)  // 8
    } else {
        fmt.Println("Массив пуст!")
    }


    // Пример с пустым массивом
    emptyNums := []int{}
    max, ok = utils.FindMax(emptyNums)
    if ok {
        fmt.Println(max)
    } else {
        fmt.Println("Массив пуст!")  // Массив пуст!
    }
// }
}

fmt.Println("Оператор defer")

{
// defer регистрирует функции для выполнения в обратном порядке их вызова!   

    defer fmt.Println("A")
    defer fmt.Println("B")
    defer fmt.Println("C")
    fmt.Println("D")

// D
// C
// B
// A
// После выполнения всех строчек кода всей программы!
}

{
// Анонимные функции

// Анонимная функция — это функция, у которой нет имени. 
// Она создаётся "на месте" и может быть сразу вызвана или присвоена переменной для последующего использования.

// Такие функции полезны, когда:

// Нужно использовать функцию только один раз!
// Функция передаётся как аргумент в другую функцию (например, для обработки данных).

    sayHello := func(name string) {
        fmt.Println("Привет,", name)
    }


    // Вызываем функцию через переменную
    sayHello("Мир")  // Привет, Мир
    sayHello("Go")   // Привет, Go
}

{
    fmt.Println("Как сразу вызвать анонимную функцию")

    // Сразу вызываем анонимную функцию

    func(a, b int) {
        fmt.Println("Сумма:", a+b)
    }(3, 7)  // Передаём аргументы 3 и 7

    // Выведет: 10
}

{
    // Присваиваем анонимную функцию переменной
    multiply := func(a, b int) int {
        return a * b
    }

    // Вызываем функцию
    result := multiply(4, 5)
    fmt.Println("Результат:", result) // Результат: 20
}
}


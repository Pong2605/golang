package main

import "fmt"

// 3.1  Переменные
// 3.2  Типы данных
// 3.3  Константы

func main() {
	var x int8 = 2
	var t bool = true
	s := "Hello"
	const pi = 3.14
	fmt.Println(s + "!", x, t)
	fmt.Println(pi)
	fmt.Println()

	{
		ru := "Привет"  // Кириллица (каждый символ = 2 байта)
		en := "Hello"    // Латиница (каждый символ = 1 байт)
		
		fmt.Println("len() возвращает количество байт, а не символов!")

		// len() возвращает количество байт, а не символов!
		fmt.Println(len(ru), len(en)) // 12 (6 символов × 2 байта), 5
		
		// Обращение по индексу даёт байт, а не символ:
		fmt.Println("Обращение по индексу даёт байт, а не символ:")

		fmt.Println(ru[0], string(ru[0])) // 208, "Ð" (это половинка 'П')
		fmt.Println(en[0], string(en[0])) // 72, "H" (корректно, т.к. ASCII)

		// Правильный способ: конвертация в руны
		fmt.Println("Правильный способ: конвертация в руны")
		runes := []rune(ru)
		fmt.Println(runes[0], string(runes[0])) // 1055, "П" (теперь верно)
		fmt.Println(len(runes))                 // 6 (реальное кол-во символов)

		// Итерироваться по строке лучше так:
		fmt.Println("Итерироваться по строке лучше так:")
		for idx, r := range ru { // range автоматически работает с рунами
			fmt.Printf("Символ %d: %c (кодовая точка: %d)\n", idx, r, r)
	}
	}

// 3.4 Арифметические операторы

	{
		var x float32 = 13.0 / 4
		fmt.Println(x)

		var y float32 = 13 / 4
		fmt.Println(y)

		var a int = 2
		a++            // аналог выражения a = a + 1
		fmt.Println(a) // выведет 3
	}

// 3.5 Операторы сравнения и логические операторы

	{
		var a int = 31
		var b int = 14
			
		// логическое И
		fmt.Println(a != b && a >= b)  // выведет true
			
		// логическое ИЛИ
		fmt.Println(a == b || a > b)   // выведет true
			
		// логическое НЕ
		fmt.Println(!(a < b))          // выведет true

		fmt.Println(10 > 7 && 5 < 6 || 5 == 6)
	}

// 3.6  Принимаем ввод

	{
		// var number int
		// fmt.Scan(&number)

		// var a, b, c int
		// fmt.Scan(&a, &b, &c)

		// подаем на вход:
		// 4 5
		// 6
		// 7
		// получаем вывод: 4 5 6 7
		// все отлично считалось, разбивка по пробелам и переносам строк
		// var a, b, c, d int
		// fmt.Scan(&a, &b, &c, &d)
		// fmt.Print(a, b, c, d)

		// подаем на вход:
		// 4 5
		// 6
		// 7
		// получаем вывод: 4 5 0 0
		// считалась первая строка в первый и второй аргументы, 
		// остальные строки не считались
		// var a, b, c, d int
		// fmt.Scanln(&a, &b, &c, &d)
		// fmt.Print(a, b, c, d)
	}

// 3.7 Выводим информацию

	{
		a := 5
		b := 11
		fmt.Print(a, b) 
		// выведет: 5 11 

		fmt.Print("Привет,", "Степик")
		// выведет: Привет,Степик

		fmt.Print(a, b)
		fmt.Print("Привет,", "Степик")
		// выведет: 5 11Привет,Степик

		a = 5
		b = 11
		fmt.Println(a, b)
		// выведет 5 11

		fmt.Println("Привет,", "Степик")
		// выведет: Привет, Степик

		fmt.Println(a, b)
		fmt.Println("Привет,", "Степик")
		/* выведет:
		5 11
		Привет, Степик
		*/
	}

// 4.1 Условный оператор if/else

	{
		pass := 7055
		if pass == 7055 { 
			fmt.Println("Пароль верный")
		}
		// пароль верный, условие равно true
		// код будет выполнен и выведет: Пароль верный

		pass = 7055
		if pass == 3142 { 
			fmt.Println("Пароль неверный")
		}
		// пароль неверный, условие равно false
		// код не будет выполнен и ничего не будет выведено

		iq := 70
		if iq > 120 {
			fmt.Println("Вы гений!")
		} else {
			fmt.Println("Увы, но вы не гений")
		}

		var a, b int
		a, b = 1, 2
		if a > b {
			fmt.Println("a больше чем b")
		} else if a == b {
			fmt.Println("a равно b")
		} else {
			fmt.Println("a меньше чем b")
		}

		// Иногда может случиться так, что вам понадобится переменная только для условного оператора.
		// В такой ситуации оператор if в Go может начинаться с объявления переменной перед условием:

		if iq := 125; iq > 120 {
			fmt.Println("Гений!")
		} else {
			fmt.Println("Увы(")
		}

		if x := 33; x % 2 == 1 {
			fmt.Println(x + 9)
		} else {
			fmt.Println(x - 3)
		}

		height := 177
		if height <= 165 {
			fmt.Println("Человек невысокого роста")
		} else if height > 165 && height <= 185 {
			fmt.Println("Человек среднего роста")
		} else {
			fmt.Println("Высокий человек")
		}
		// выведет: Человек среднего роста

		var balance int = 1100
		
		if balance > 1000 {
			fmt.Println("Apple")
		} else if balance < 500 {
			fmt.Println("Nokia с фонариком")
		} else {
			fmt.Println("Samsung")
		}
	}

// 4.2  Оператор switch

	{
		var balance int = 1100

		var out string
		switch {
			case balance > 1000: out = "Apple"
			case balance >= 500 && balance <= 1000: out = "Samsung"
			case balance < 500: out = "Nokia с фонариком"
		}	
		fmt.Println(out)

		x := 2

		switch x {
			case 1:
				fmt.Println("Один")
			case 2:
				fmt.Println("Два")
			case 3:
				fmt.Println("Три")
			default:
				fmt.Println(x)
		}

		month := "Январь"

		switch {
			case month == "Декабрь" || month == "Январь" || month == "Февраль":
				fmt.Println("Зима")
			case month == "Март" || month == "Апрель" || month == "Май":
				fmt.Println("Весна")
		}

/* Использование в switch ключевого слова fallthrough. 
Если в конце блока case написать fallthrough, 
то тело следующего case будет выполнено независимо от того 
истинно ли его (следующего за ним case) условие:
*/

		day := 3

		switch day {
			case 1:
				fmt.Println("Понедельник")
				fallthrough
			case 2:
				fmt.Println("Вторник")
				fallthrough
			case 3:
				fmt.Println("Среда")
				fallthrough
			case 4:
				fmt.Println("Четверг")
				fallthrough
			case 5:
				fmt.Println("Пятница")
			default:
				fmt.Println("Неправильный день")
		}

		/* выведет
		Среда
		Четверг
		Пятница
		*/
	}

// 4.3 Циклы

// Цикл for

	{
		for i := 0; i <= 5; i++ {
		fmt.Println(i)
		}

// объявление переменной для счетчика можно вынести за пределы цикла и
// изменение счетчика можно переместить в тело цикла:

		var j = 0
		for ; j < 5;{
			fmt.Println(j)
			j++
		}

// Если цикл использует только условие, то его можно сократить следующим образом (убрать разделяющие точки с запятой):
// В таком виде это получается аналог цикла while, который используется в других языках программирования.

		var k = 0
		for k < 5 {
			fmt.Println(k)
			k++
		}

// Оператор continue	

		sum_even := 0
		
		for i := 1; i <= 10; i++ {
			if i % 2 != 0 {
				continue             // число нечетное, то пропускаем его 
			}                        // и переходим к следующей итерации
			sum_even += i
		}
		fmt.Println(sum_even)        // выведет: 30

// Оператор break

		for i := 1; i <= 9; i++{
			if i % 5 == 0 {
				break            // если число кратно 5, то выходим из цикла
			}
			fmt.Println(i)
		}
		/* выведет:
		1
		2
		3
		4
		*/
	}
}

// 5.1 Введение в функции

// func helloStepik() {
// 	fmt.Println("Привет, Степик!") // тело функции
// }
// func main() {
// 	helloStepik()
// }

// 5.2 Аргументы

// func mult(x, y int) {
//     fmt.Println(x * y)
// }

// func main() {
// 	var a, b int
// 	fmt.Scan(&a, &b)
// 	mult(a, b)
// }

// 5.3 Возврат из функции

// func concat(a, b string) string {
//     res := a + b
//     return res
// }

// func plusMinus(x, y int) (int, int) {
//     return x + y, x - y
// }

// 5.4 Defer

// {

// package main

// import "fmt"

// func welcome() {
//     fmt.Println("Добро пожаловать")
// }

// func main() {
//     defer welcome()
//     fmt.Println("Привет!")
// }

                  
// Приведенный выше код сначала выведет Привет! и только после этого выведет результат функции welcome().
// Это происходит потому, что вызов функции welcome() отложен, то есть она ждет, пока main() закончит выполнение, и только после этого происходит её вызов.

// defer часто используется для очистки, например, для освобождения ресурсов, используемых кодом, таких как файлы, соединения и т.д.

// }

// func main() {
//     fmt.Println("начало")

//     for i := 0; i < 5; i++ {
//         defer fmt.Println(i)
//     }
//     fmt.Println("конец")
// }

// Вывод:
// начало
// конец
// 4
// 3
// 2
// 1
// 0

// 5.5 Область видимости

// Локальные переменные всегда получают приоритет над глобальными переменными с тем же именем.
// Но именование глобальных и локальных переменных одинаковыми именами является плохой практикой.

// {
// package main

// import "fmt"

// var s = "глобальная переменная"

// func scope() {
//     var s = "локальная переменная"
// }

// func main() {
//     fmt.Println(s) // глобальная переменная
// }
// }

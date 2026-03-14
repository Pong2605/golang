package utils

import "fmt"

func Test() {
	fmt.Println("Hellow from another package")
}

func Sum(nums ...int) int {
total := 0
for _, num := range nums {
	total += num
}
return total
}

func Print(args ...any) {
    for _, arg := range args {
        fmt.Print(arg, " ")
    }
}

func Divide(a, b int) (int, bool) {
    if b == 0 {
        return 0, false
    }
    return a / b, true
}

func FindMax(numbers []int) (max int, ok bool) {
    if len(numbers) == 0 {
        max, ok = 0, false  // Присваиваем значения именованным возвращаемым переменным
        return              // Возвращаем без указания переменных max и ok явно
    }
    
    max = numbers[0]   // max уже объявлен как возвращаемая переменная
    for _, num := range numbers {
        if num > max {
            max = num
        }
    }
    ok = true  // Устанавливаем флаг
    return     // Возвращаем автоматически
}
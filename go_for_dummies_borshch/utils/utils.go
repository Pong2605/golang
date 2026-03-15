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

func FindMax2(arr []int) int {
    max := arr[0]
    for i := 1; i < len(arr); i++ {
        if arr[i] > max {
            max = arr[i]
        }
    }
    return max
}

func FindMin2(arr []int) int {
    min := arr[0]
    for i := 1; i < len(arr); i++ {
        if arr[i] < min {
            min = arr[i]
        }
    }
    return min
}


func IsPalindrome(input string) bool {
    n := len(input) / 2
    for i := 0; i < n; i++ {
        if input[i] != input[len(input) - i - 1] {
        return false
        }
    }
    return true
}

func LongestUniqueSubstring(s string) (int, string) {
	// Карта для хранения последней позиции каждого символа
	charMap := make(map[rune]int)
	
	maxLength := 0
	start := 0          // Начало текущей подстроки
	maxStart := 0       // Начало самой длинной подстроки
	
	// Проходим по строке
	for i, char := range s {
		// Если символ уже встречался и его позиция >= start
		if pos, exists := charMap[char]; exists && pos >= start {
			start = pos + 1 // Сдвигаем начало подстроки
		}
		
		// Обновляем позицию символа
		charMap[char] = i
		
		// Проверяем, не стала ли текущая подстрока длиннее максимальной
		currentLength := i - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
			maxStart = start
		}
	}
	
	// Извлекаем подстроку
	if maxLength > 0 {
		substring := s[maxStart : maxStart+maxLength]
		return maxLength, substring
	}
	
	return 0, ""
}
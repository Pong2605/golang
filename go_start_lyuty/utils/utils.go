package utils

import (
	"fmt"
	"time"
)

var D = 4

func GetValue() string{
    return "Hello from this another package"
}

func Out(from, to int) {
  for i := from; i <= to; i++ {
    time.Sleep(50 * time.Millisecond)
    fmt.Println(i)
  }
}

func OutCh(from, to int, ch chan bool) {
  for i := from; i <= to; i++ {
      time.Sleep(50 * time.Millisecond)
      fmt.Println(i)
  }
  ch <- true
}

// Как видите, наши функции передают результат в каналы. 
// Теперь мы можем вызвать их как горутины в main() и вывести полученную сумму.

// Мы используем каналы для получения результата каждой горутины и вывода их суммы.

// Если вам больше не нужно отправлять данные в канал, вы можете закрыть его, используя close(ch), 
// где ch - имя канала. Это делается в отправителе.

func EvenSum(from, to int, ch chan int) {
  result := 0
  for i := from; i <= to; i++ {
    if i % 2 == 0 {
      result += i
    }    
  }
  ch <- result
}

func SquareSum(from, to int, ch chan int) {
  result := 0
  for i := from; i <= to; i++ {
    if i % 2 == 0 {
      result += i*i
    }    
  }
  ch <- result
}

func Download(size int, ch chan int) {
  result := 0
  for i := 0; i <= size; i++ {
      result += i
  }
  ch <- result
}
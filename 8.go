package main

import (
	"fmt"
)

func setBit(n int64, i int, value int) int64 {
	i -= 1
	if i < 0 || i > 63 {
		fmt.Println("Incorrect bit number")
		return n
	}
	if value == 1 {
		//создаем число у которого только i-ый бит 1, а остальное 0.
		// Применяем побитовое ИЛИ к исходному числу и сформированному
		n |= (1 << i)
	} else if value == 0 {
		//создаем число у которого только i-ый бит 1, а остальное 0.
		// инвертируем его, теперь там только i-ый бит 0, а остальное 1.
		// Применяем побитовое И к исходному числу и сформированному
		n &= ^(1 << i)
	} else {
		fmt.Println("Value must be 0 or 1")
	}

	return n
}

func main() {
	number := int64(20)
	bitNumber := 3
	bitValue := 0
	fmt.Printf("if in %d put %d to %d bit we will get %d\n", number, bitValue, bitNumber, setBit(number, bitNumber, bitValue))
	fmt.Println("Completed")
}

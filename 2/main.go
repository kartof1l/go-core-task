package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var originalSlice []int = randomSlice()
	fmt.Println("Исходный слайс:", originalSlice)
	fmt.Println("Только четные:", sliceExample(originalSlice))
	fmt.Println("Добавил элемент:", addElements(originalSlice, 5))
	fmt.Println("Копия:", copySlice(originalSlice))
	fmt.Println("Удаление элемента:", removeElement(originalSlice, 5))

	fmt.Println("\nПроверка оригинального слайса:", originalSlice)
}

func randomSlice() []int {
	slice := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		slice = append(slice, rand.IntN(100)) 
	}
	return slice
}

func sliceExample(slice []int) []int {
	result := make([]int, 0)
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func addElements(slice []int, value int) []int {
	biggerSlice := make([]int, len(slice))
	copy(biggerSlice, slice)
	return append(biggerSlice, value)
}

func copySlice(slice []int) []int {
	if len(slice) == 0 {
		return []int{}
	}
	return append([]int{slice[0]}, copySlice(slice[1:])...)
}

func removeElement(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return copySlice(slice)
	}
	result := make([]int, 0, len(slice)-1)
	result = append(result, slice[:i]...)
	result = append(result, slice[i+1:]...)
	return result
}
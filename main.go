package main

import (
	"fmt"
)

func main() {
	var arr []string
	arr = append(arr, "a", "b", "c", "d", "e", "f", "g", "h", "i")
	arr2 := []string{"1", "2", "3"}
	fmt.Println(arr)
	fmt.Println(arr2)
	for a, b := range arr {
		fmt.Println(a, b)
	}
	for a, b := range arr2 {
		fmt.Println(a, b)
	}
	fmt.Println(arr[:1])
	fmt.Println(arr[1:])
	fmt.Println(arr[1:2])
	i := 0
	arr = append(arr[:i], arr[i+1:]...)
	fmt.Println(arr)

	arr = append(arr[:i], arr[i+1:]...)
	fmt.Println(arr)

	arr3 := make([]string, len(arr))
	copy(arr3, arr)
	// or
	arr4 := append([]string(nil), arr...)
	fmt.Println(arr3, arr4)
	reverse(arr)

	str := "abcde"
	fmt.Println(str)
}

func reverse(numbers []string) []string {
	fmt.Println(numbers)
	for i := len(numbers)/2 - 1; i >= 0; i-- {
		opp := len(numbers) - 1 - i
		fmt.Println(numbers[i], numbers[opp])
		temp := numbers[i]
		numbers[i] = numbers[opp]
		numbers[opp] = temp
	}
	fmt.Println(numbers)
	return numbers
}

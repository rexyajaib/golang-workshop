package main

import "fmt"

func Add(slice []int) {
	slice = append(slice, 10) // new value will not included
	fmt.Printf("[Add] %v, len:%d, cap: %d\n", slice, len(slice), cap(slice))
}

func Mutate(slice []int) {
	slice[0] = 1
	fmt.Printf("[Mutate] %v, len:%d, cap: %d\n", slice, len(slice), cap(slice))
}

func main() {
	sliceOfInt := make([]int, 0, 10) // define length and capacity
	fmt.Printf("[Main] %v, len:%d, cap: %d\n", sliceOfInt, len(sliceOfInt), cap(sliceOfInt))
	Add(sliceOfInt)
	fmt.Printf("[Main] %v, len:%d, cap: %d\n", sliceOfInt, len(sliceOfInt), cap(sliceOfInt))
	sliceOfInt = append(sliceOfInt, 10)
	Mutate(sliceOfInt)
	fmt.Printf("[Main] %v, len:%d, cap: %d\n", sliceOfInt, len(sliceOfInt), cap(sliceOfInt))

	mapOfInt := make(map[int]int)
	mapOfInt[1] = 1
	mapOfInt[2] = 2
	fmt.Println(mapOfInt)
	delete(mapOfInt, 1)
	fmt.Println(mapOfInt)
}

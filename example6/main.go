package main

import "fmt"

// 1) Generic type

type Box[T any] struct {
	value T
}

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func (b *Box[T]) Set(v T) {
	b.value = v
}

func (b *Box[T]) Get() T {
	return b.value
}

//func (b *Box[T]) Test[E any]() E { // generic method is not supported
//
//}

func Add[T Numeric](a, b T) T {
	return a + b
}

func main() {
	intBox := Box[int]{}
	intBox.Set(42)
	fmt.Println(intBox.Get())

	strBox := Box[string]{}
	strBox.Set("hello")
	fmt.Println(strBox.Get())

	//a, b := "10", "20"
	//fmt.Println(Add[string](a, b)) // "1020" // Error Compile

	// Compile-time safety
	// intBox.Set("oops") // ❌ compile-time error
}

package main

import "fmt"

type Box struct {
	value int
}

func change(b Box) {
	b.value = 10 // copy b and original box remains unchanged
}

func changePtr(b *Box) {
	b.value = 10 // modify the original box via pointer
}

func main() {
	box := Box{value: 1}

	change(box)
	fmt.Println(box.value) // 1 ❌

	changePtr(&box)
	fmt.Println(box.value) // 10 ✅
}

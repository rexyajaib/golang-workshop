package main

import (
	"errors"
	"fmt"
)

func riskyOperation() error {
	return errors.New("something went wrong")
}

func process() error {
	fmt.Println("Start process")
	defer fmt.Println("Defer: cleanup always runs")

	if err := riskyOperation(); err != nil {
		fmt.Println("Got error:", err)
		return err
	}

	fmt.Println("After risky operation")
	return nil
}

func main() {
	err := process()
	if err != nil {
		fmt.Println("Process failed with error:", err)
	}
}

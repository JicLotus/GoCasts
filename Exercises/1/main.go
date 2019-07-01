package main

import "fmt"

type numbers []int

func main() {
	n := getNumbers(11)
	n.printEvenOdd()
}

func (n numbers) printEvenOdd() {
	for _, number := range n {
		if number%2 == 0 {
			fmt.Println(number, "is", "even")
		} else {
			fmt.Println(number, "is", "odd")
		}
	}
}

func getNumbers(amount int) numbers {
	newNumbers := numbers{}
	for i := 0; i < amount; i++ {
		newNumbers = append(newNumbers, i)
	}
	return newNumbers
}

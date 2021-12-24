package main

import (
	"fmt"
)

func greetings(msg string) string {
	return msg
}

func getSum(num1, num2 int64) int64 {
	return int64(num1) + int64(num2)
}

func main() {
	fmt.Println(greetings("Hello world..."), getSum(3,5))
}

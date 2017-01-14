package main

import (
	"fmt"
)

func main() {
	var result int64 = 0
	for i := 1; i < 1000; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			result += int64(i)
		}
	}
	fmt.Println(result)
}

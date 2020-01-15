package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := []int{}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < 10; i++ {
		num := r.Intn(100)
		nums = append(nums)
		if (num % 2) == 0 {
			fmt.Println(num, "even")
		} else {
			fmt.Println(num, "odd")
		}
	}
}

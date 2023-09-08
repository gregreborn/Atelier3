// Atelier3.go
package Atelier3

import (
	"math/rand"
	"time"
)

func init() {
	time.Now().UnixNano()
}

func ArrayGenerate(size int) []int {
	if size <= 0 {
		return nil
	}

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(10000) + 1
	}

	return arr
}

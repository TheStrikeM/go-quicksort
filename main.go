package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func swap(a, b *int) {
	*a, *b = *b, *a
}

func partitionX(predicate int, slice []int, left, right int) (int, int, int) {
	E := left
	G := left

	for i := left; i < right; i++ {
		if slice[i] < predicate {
			y := slice[i]
			slice[i] = slice[G]
			G++
			slice[G-1] = slice[E]
			slice[E] = y
			E++
		} else if slice[i] == predicate {
			swap(&slice[i], &slice[G])
			G++
		}
	}

	return E, right - E, G
}

func merge(slice1, slice2 []int) []int {
	merged := make([]int, len(slice1)+len(slice2))
	return merged
}

func Qsort(slice []int, left, right int) {
	if right-left < 2 {
		return
	}
	p, _, G := partitionX(slice[rand.Intn(right-left)+left], slice, left, right)
	Qsort(slice, left, p)
	Qsort(slice, G, right)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var lenNums int
	fmt.Fscan(reader, &lenNums)

	slice := make([]int, lenNums)
	for i := 0; i < lenNums; i++ {
		fmt.Fscan(reader, &slice[i])
	}

	Qsort(slice, 0, lenNums)

	writer := bufio.NewWriter(os.Stdout)
	for _, item := range slice {
		fmt.Fprint(writer, item)
		fmt.Fprint(writer, " ")
	}
	writer.Flush()
}

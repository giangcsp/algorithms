package main

import (
	"fmt"

	"github.com/giangcsp/algorithms/go/pkg/countingsort"
	"github.com/giangcsp/algorithms/go/pkg/mergesort"
	"github.com/giangcsp/algorithms/go/pkg/quicksort"
)

func main() {
  a := []int{9, 12, 4, 7, 1, 0, 39}
  fmt.Println("Quicksort before", a)
  quicksort.Quicksort(a)
  fmt.Println("Quicksort after", a)

  b := []int{9, 12, 4, 7, 1, 0, 39}
  fmt.Println("Mergesort before", b)
  mergesort.Mergesort(b)
  fmt.Println("Mergesort after", b)

  c := []int{9, 12, 4, 7, 1, 0, 39, -10}
  fmt.Println("CountingSort before", c)
  countingsort.CountingSort(c)
  fmt.Println("CountingSort after", c)
}

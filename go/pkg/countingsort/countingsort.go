package countingsort

import (
	"golang.org/x/exp/constraints"
)

func CountingSort[N constraints.Integer](a []N) {
  if len(a) < 1 {
    return
  }

  maxVal := a[0]
  minVal := a[0]
  for _, n := range a {
    if maxVal < n {
      maxVal = n
    }

    if minVal > n {
      minVal = n
    }
  }

  counts := make([]int, maxVal - minVal + 1)
  for _, n := range a {
    counts[int(n - minVal)] += 1
  }

  for i, n := 0, 0; n < len(counts); n++ {
    for j := 0; j < counts[n]; j++ {
      a[i] = N(n) + minVal
      i++
    }
  }
}

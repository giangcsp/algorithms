package quicksort

import "cmp"

func Quicksort[N cmp.Ordered](a []N) {
  partition := func(start int, stop int) int {
    j := start
    for i := start; i < stop; i++ {
      if a[i] < a[stop] {
        a[i], a[j] = a[j], a[i]
        j++
      }
    }
    
    a[j], a[stop] = a[stop], a[j]
    return j
  }

  var helper func(start int, stop int)
  helper = func(start int, stop int) {
    if start >= stop {
      return
    }

    pivot := partition(start, stop)
    helper(start, pivot - 1)
    helper(pivot + 1, stop)
  }

  helper(0, len(a) - 1)
}

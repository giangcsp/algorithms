package mergesort

import "cmp"

func merge[N cmp.Ordered](a []N, start int, mid int, stop int) {
  i, j, n := start, mid+1, 0
  temp := make([]N, stop - start + 1)

  for ; i <= mid && j <= stop; n++ {
    if a[i] < a[j] {
      temp[n] = a[i]
      i++
    } else {
      temp[n] = a[j]
      j++
    }
  }

  for ; i <= mid; i, n = i + 1, n + 1 {
    temp[n] = a[i]
  }

  for ; j <= stop; j, n = j + 1, n + 1 {
    temp[n] = a[j]
  }

  for i := 0; i < stop - start + 1; i++ {
    a[i+start] = temp[i]
  }
}

func Mergesort[N cmp.Ordered](a []N) {
  var helper func(start int, stop int)
  helper = func(start int, stop int) {
    if start >= stop {
      return
    }

    mid := (start + stop) / 2
    helper(start, mid)
    helper(mid+1, stop)
    merge(a, start, mid, stop)
  }

  helper(0, len(a) - 1)
}

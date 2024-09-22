package quicksort

func Quicksort(items QuickSortType) {
  partition := func(start int, stop int) int {
    j := start
    for i := start; i < stop; i++ {
      if items.LessThan(i, stop) {
        items.Swap(i, j)
        j++
      }
    }
    
    items.Swap(j, stop)
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

  helper(0, items.Size() - 1)
}

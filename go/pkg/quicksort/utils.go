package quicksort

type quicksortInt []int

func (a quicksortInt) LessThan(i int, j int) bool {
  return a[i] < a[j]
}

func (a quicksortInt) Swap(i int, j int) {
  a[i], a[j] = a[j], a[i]
}

func (a quicksortInt) Size() int {
  return len(a)
}

func SortInt(a []int) {
  Quicksort(quicksortInt(a))
}

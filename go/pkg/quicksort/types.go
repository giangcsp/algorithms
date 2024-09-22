package quicksort

type QuickSortType interface {
  LessThan(i int, j int) bool
  Swap(i int, j int)
  Size() int
}

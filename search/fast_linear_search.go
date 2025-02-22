package search


func LinearSearch(arr []int, x int) int {
	i := 0
	count := len(arr)

	for i < count{
		if arr[i] == x {
			return i
		}
		i++
	}
	return -1
}
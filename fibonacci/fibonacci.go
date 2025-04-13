package fibonacci

// Factorial menghitung faktorial dari n.
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// Fibonacci menghasilkan urutan Fibonacci hingga batas n.
func Fibonacci(n int) []int {
	sequence := []int{}
	a, b := 0, 1
	for a <= n {
		sequence = append(sequence, a)
		a, b = b, a+b
	}
	return sequence
}

// BinarySearch mencari target dalam array yang diurutkan arr.
// Mengembalikan indeks target jika ditemukan, atau -1 jika tidak ditemukan.
func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

<<<<<<< HEAD
// Package main code
=======
// Package main code go-algoritm
>>>>>>> eb21fb6 (add comment)
package main

import (
	"fmt"

	"github.com/RajaSunrise/go-algorithm/search"
	"github.com/RajaSunrise/go-algorithm/short"
)


func main(){
<<<<<<< HEAD
<<<<<<< HEAD
=======
	// array
>>>>>>> 13a77bf (add comment)
=======
	// list array
>>>>>>> eb21fb6 (add comment)
	arr := []int{12, 23, 39, 12, 21, 12, 6, 76, 15}
	
	// test binary seacrh
	fmt.Println("ini adalah binary search", search.BinarySearch(arr, 12))
	// test fast linear search
	fmt.Println("ini adalah fast linear search", search.LinearSearch(arr, 12))
	// test interpolination search
	fmt.Println("Ini adalah interpolination search", search.InterPolinationSearch(arr, 12))

	// test bubble short
	fmt.Println("Ini adalah Bubble Short", short.BubbleShort(arr))
	// test counting short
	fmt.Println("Ini adalah counting Short", short.CountingShort(arr))

<<<<<<< HEAD
<<<<<<< HEAD
}
=======
}
>>>>>>> 13a77bf (add comment)
=======
}
>>>>>>> eb21fb6 (add comment)

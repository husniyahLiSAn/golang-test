package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Enter size of your array: ")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("\nYour array is: ", arr)

	max := findMax(arr)
	printout(arr, max)

	// Sorting the array
	fmt.Println("\nSorting Array by Ascending...")
	bubbleSort(arr, max)

	fmt.Println("\nSorting Array by Descending...")
	reverseSort(arr, max)
}

func findMax(a []int) (max int) {
	max = a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

func bubbleSort(a []int, max int) {
	n := len(a)
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				fmt.Println("\nSwapping")
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
				printout(a, max)
			}
		}
	}
}

func reverseSort(a []int, max int) {
	n := len(a)
	swapped := true

	for swapped {
		swapped = false
		for i := n - 1; i > 0; i-- {
			if a[i] > a[i-1] {
				fmt.Println("\nSwapping")
				a[i], a[i-1] = a[i-1], a[i]
				swapped = true
				printout(a, max)
			}
		}
	}
}

func printout(a []int, max int) {
	for i := max; i >= 0; i-- {
		for j := 0; j < len(a); j++ {
			if a[j] <= i {
				fmt.Print(" ")
			} else {
				fmt.Print("|")
			}
			fmt.Print("  ")
		}
		fmt.Println()
	}
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], "  ")
	}
	fmt.Println()
}

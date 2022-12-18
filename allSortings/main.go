package main

import "fmt"

//					best			avg				worst
//Selection Sort	Ω(n^2)			θ(n^2)			O(n^2)
//Bubble Sort		Ω(n)			θ(n^2)			O(n^2)
//Insertion Sort	Ω(n)			θ(n^2)			O(n^2)
//Heap Sort			Ω(n log(n))		θ(n log(n))		O(n log(n))
//Quick Sort		Ω(n log(n))		θ(n log(n))		O(n^2)
//Merge Sort		Ω(n log(n))		θ(n log(n))		O(n log(n))
//Bucket Sort		Ω(n+k)	θ(n+k)	O(n^2)
//Radix Sort		Ω(nk)	θ(nk)	O(nk)

func main() {

	x := []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	bubblesort(x)
	selectionsort(x) //selection sort (create a min as min=a[0]…..find the min and then finally swap with with a[0] )
	insertionsort(x) //insertion (split the array between sorted and unsorted then algo )
}

func bubblesort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println("Bubble sort :          ", a)
}
func selectionsort(a []int) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				a[j], a[min] = a[min], a[j]
			}
		}
	}
	fmt.Println("selection sort :          ", a)

}
func insertionsort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < i; j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println("insertion sort :           ", a)

}

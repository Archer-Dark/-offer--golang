package main

import "fmt"

func main()  {
	x := []int{3,5,1,2,4,9,6,8,7}
	y := QuickSort(x)
	fmt.Println(y)
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	i,j := 1,len(arr)-1
	for i <= j {
		if arr[i] > arr[0] {
			arr[i],arr[j] = arr[j],arr[i]
			j--
		} else {
			arr[i-1],arr[i] = arr[i],arr[i-1]
			i++
		}
	}
	QuickSort(arr[:i-1])
	QuickSort(arr[i:])
	return arr
}

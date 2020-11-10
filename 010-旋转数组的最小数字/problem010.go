package main

import "fmt"

// func minArray(numbers []int) int {
//     for i:=0;i<len(numbers)-1;i++{
//         if numbers[i]>numbers[i+1]{
//             return numbers[i+1]
//         }
//     }
//     return numbers[0]
// }

// func minArray(numbers []int) int {
//     for i:=len(numbers)-1;i>0;i--{
//         if numbers[i] < numbers[i-1]{
//             return numbers[i]
//         }
//     }
//     return numbers[0]
// }

func minArray(numbers []int) int {
	l,r := 0,len(numbers)-1
	for l<r{
		m := l+((r-l)>>1)
		if numbers[m] > numbers[r]{
			l = m+1
		}else if numbers[m] < numbers[r] {
			r = m
		}else {
			r--
		}
	}
	return numbers[l]
}

func main()  {
	arr := []int{3,4,5,1,2}
	fmt.Println(minArray(arr))
}
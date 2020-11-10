package main

import "fmt"

func numWays(n int) int {
	s1,s2:=1,1
	for i:=0;i<=n;i++{
		if i == n {
			return s1
		}
		s1,s2 = s2,(s1+s2)%(1e9+7)
	}
	return -1
}

func main()  {
	fmt.Println(numWays(6))
}
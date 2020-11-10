package main

import "fmt"

func fib(n int) int {
	f1,f2 := 0,1
	for i := 0;i<=n;i++{
		if i==n{
			return f1
		}
		f1,f2 = f2,(f1+f2)%(1e9+7)
	}
	return -1
}

func main()  {
	fmt.Println(fib(6))
}
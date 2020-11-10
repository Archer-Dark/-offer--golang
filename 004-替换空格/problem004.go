package main

import (
	"fmt"
)

func replaceSpace(s string) string {
	count := 0
	for i := 0;i<len(s);i++{
		if s[i] == ' ' {
			count++
		}
	}
	ns := make([]byte,len(s)+count*2)
	//j := 0
	for i,j := 0,0;i<len(s);i++ {
		if s[i] == ' ' {
			ns[j] = '%'
			j++
			ns[j] = '2'
			j++
			ns[j] = '0'
			j++
		}else {
			ns[j] = s[i]
			j++
		}
	}
	return string(ns)

}

func main()  {
	str:= "We Are Happy"
	nstr := replaceSpace(str)
	fmt.Println(nstr)
}

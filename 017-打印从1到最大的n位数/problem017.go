package main

import (
	"fmt"
	"strconv"
)

//func printNumbers(n int) []int {
//	max := int(math.Pow10(n))-1
//	res := make([]int,0)
//	for i := 1;i <= max;i++{
//		res = append(res,i)
//	}
//	return res
//}

func printNumbers(n int)  {
	base :=[]string{}
	for i:=0;i<10;i++{
		base = append(base,strconv.Itoa(i))
	}
	for l:=1;l<=n-1;l++ {
		tmp:=make([]string,len(base))
		copy(tmp,base)
		digit :=l+1 //补齐位数
		for i:=1;i<10;i++{
			for _,v:=range tmp{
				diff_len := digit-len(v)
				tmp_0:= ""
				for x:=1;x<diff_len;x++ {
					tmp_0+="0"
				}
				v = tmp_0+v
				item :=strconv.Itoa(i)+v
				base = append(base,item)
			}
		}
	}
	base =base[1:]
	fmt.Println((base))
}
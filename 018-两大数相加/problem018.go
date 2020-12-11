package main

import "strconv"

func mySum(x,y string) string {
	if len(x) < len(y) {
		x,y = y,x
	}
	difference := len(x) - len(y)
	yhead := ""
	for i:= 0;i<difference;i++{
		yhead = yhead+"0"
	}
	y = yhead+y
	carry := 0
	newNum := ""
	for i := len(x)-1;i>=0;i-- {
		a,_ := strconv.Atoi(x[i:i+1])
		b,_ := strconv.Atoi(y[i:i+1])
		c := a + b + carry
		if c >= 10 {
			newNum = strconv.Itoa(c - 10) + newNum
			carry = 1
		}else if i > 0{
			carry = 0
		}
	}
	if carry > 0 {
		newNum = "1"+newNum
	}
	return newNum
}

func main()  {
	println(mySum("123456789","987654321"))
}
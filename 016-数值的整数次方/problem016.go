package main

//快速幂解析（二分法角度）：
//快速幂实际上是二分思想的一种应用。
//算法流程：
//	当 x=0 时：直接返回 0（避免后续 x=1/x 操作报错）。
//	初始化 res=1 ；
//	当 n<0 时：把问题转化至 n \geq 0n≥0 的范围内，即执行x=1/x ，n=−n ；
//	循环计算：当 n = 0n=0 时跳出；
//		当 n&1=1 时：将当前 xx 乘入 resres （即 res *= xres∗=x ）；
//		执行 x = x^2（即 x∗=x ）；
//		执行 n 右移一位（即 n >>= 1）。
//	返回 res

//def myPow(x: float, n: int) -> float:
//	if x == 0: return 0
//	res = 1
//	if n < 0: x, n = 1 / x, -n
//	while n:
//		if n & 1: res *= x
//		x *= x
//		n >>= 1
//	return res

func myPow(x float64,n int) float64 {
	if x==0{
		return 0
	}
	res := 1.0
	if n < 0 {
		x,n = 1/x,-n
	}
	for n > 0{
		if (n&1) == 1{
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}
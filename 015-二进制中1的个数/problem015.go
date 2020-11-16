package main

import (
	"fmt"
	"strings"
)

//方法一：巧用 n&(n−1)
//	(n−1) 解析： 二进制数字 n 最右边的 1 变成 0 ，此 1 右边的 0 都变成 1 。
//	n&(n−1) 解析： 二进制数字 n 最右边的 1 变成 0 ，其余不变。
//算法流程：
//	初始化数量统计变量 resres 。
//	循环消去最右边的 1 ：当 n=0 时跳出。
//	res += 1 ： 统计变量加 1 ；
//	n &= n - 1 ： 消去数字 n 最右边的 1 。
//	返回统计数量 resres 。

func hammingWeight(n uint32) int {
	res := 0
	for n > 0{
		res += 1
		n &= (n-1)
	}
	return res
}


//方法2：逐位判断
//	根据 与运算 定义，设二进制数字 n ，则有：
//	若 n&1=0 ，则 n 二进制 最右一位 为 0 ；
//	若 n&1=1 ，则 n 二进制 最右一位 为 1 。
//	根据以上特点，考虑以下 循环判断 ：
//	判断 n 最右一位是否为 1，根据结果计数。
//	将 n 右移一位（本题要求把数字 n 看作无符号数，因此使用 无符号右移 操作）。
//算法流程：
//	初始化数量统计变量 res=0 。
//	循环逐位判断： 当 n=0 时跳出。
//	res += n & 1 ： 若 n&1=1 ，则统计数 res 加一。
//	n >>= 1 ： 将二进制数字 n 无符号右移一位（ Java 中无符号右移为 ">>>" ） 。
//	返回统计数量 res 。
func hammingWeight2(n uint32) int {
	res := 0
	for n > 0{
		res += int(n&1)
		n >>= 1
	}
	return res
}

//无脑字符串统计法
func hammingWeight3(n uint32) int {
	return strings.Count(fmt.Sprintf("%b",n),"1")
}

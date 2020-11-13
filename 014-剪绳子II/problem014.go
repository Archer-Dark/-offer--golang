package main


//数论法1：
//	绳子段切分的越多，乘积越大，且3做因数比2更优。
//	不断剪去 长度3 并用sum累乘
//	把绳子切为多个长度为 3 的片段，则留下的最后一段绳子的长度可能为 0,1,2 三种情况。
//	循环结束的三种情况：
//	    n=n-3==2
//	n长度剩下2，因 n > 4 跳出循环，return 乘积为sum*2。
//	    n=n-3==3
//	长度刚好可以被剪完，因 n > 4 跳出循环，return 乘积为sum*3。
//	    n=n-3==4
//	再剪一次的话，长度剩下1，则最后一段绳子长度为 1； 需要把最后的 3 和 1 替换为 2 和 2，因为 2 * 2 > 3 * 1； 因 n > 4 跳出循环，return 乘积 sum*4 即可。
//  此题唯一与problem013不同的地方是需要考虑取模(1e9+7) 即：1000000007
func cuttingRope(n int) int {
	if n ==0 {
		return 0
	}
	if n <=2{
		return 1
	}
	if n == 3{
		return 2
	}
	var res = 1
	for n > 4{
		res *= 3
		if res > (1e9+7){
			res = res%(1e9+7)
		}
		n -= 3
	}
	return int(res*n % (1e9+7))
}

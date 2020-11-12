package main

func get(x int) int {
	res := 0
	for x != 0 {
		res += x % 10
		x /= 10
	}
	return res
}

func movingCount(m,n,k int) int {
	if k == 0 {
		return 1
	}
	vis := make([][]bool,m)
	for i := range vis{
		vis[i] = make([]bool,n)
	}
	ans := 1
	vis[0][0] = true
	for i := 0;i < m;i++{
		for j := 0;j<n;j++{
			if (i==0 && j==0) || (get(i)+get(j)>k){
				continue
			}
			if i-1 >= 0{
				vis[i][j] = vis[i][j] || vis[i-1][j]
			}
			if j-1 >= 0 {
				vis[i][j] = vis[i][j] || vis[i][j-1]
			}
			if vis[i][j] {
				ans += 1
			}else {
				ans += 0
			}
		}
	}
	return ans
}
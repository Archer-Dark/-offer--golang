package main

func get(x int) int {
	res := 0
	for x != 0 {
		res += x % 10
		x /= 10
	}
	return res
}

//解法1：广度优先
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
			}
		}
	}
	return ans
}

//解法2：深度优先
func movingCount2(m,n,k int) int {
	if k == 0 {
		return 1
	}
	vis := make([][]bool,m)
	for i := range vis{
		vis[i] = make([]bool,n)
	}
	return dfs(0,0,m,n,k,vis)
}

func dfs(i,j,m,n,k int,vis [][]bool) int {
	if i < 0 || i >= m || j < 0 || j >= n || get(i)+get(j) > k || vis[i][j]{
		return 0
	}
	vis[i][j] = true
	return 1 + dfs(i+1,j,m,n,k,vis) + dfs(i-1,j,m,n,k,vis) +
		dfs(i,j+1,m,n,k,vis) + dfs(i,j-1,m,n,k,vis)
}
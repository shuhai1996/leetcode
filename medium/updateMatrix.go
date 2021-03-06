package main
//NO.542
func main() {

}

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	out := make([][]int, m)
	for i := 0; i < m; i++ {
		out[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] != 0 {
				out[i][j] = 10000
				if i != 0 {
					out[i][j] = min(out[i][j], out[i-1][j]+1)
				}
				if j != 0 {
					out[i][j] = min(out[i][j], out[i][j-1]+1)
				}
			} else {
				out[i][j] = 0
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if out[i][j] > 1 {
				if i < m-1 {
					out[i][j] = min(out[i][j], out[i+1][j]+1)
				}
				if j < n-1 {
					out[i][j] = min(out[i][j], out[i][j+1]+1)
				}
			}
		}
	}
	return out
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

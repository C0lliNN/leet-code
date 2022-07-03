package main

func main() {

}

func combine(n int, k int) [][]int {
	combs := make([][]int, 0)
	comb := make([]int, 0)

	backtracking(&combs, &comb, 1, n, k)
	return combs
}

func backtracking(combs *[][]int, comb *[]int, start, n, k int) {
	if len(*comb) == k {
		newSlice := make([]int, len(*comb))
		for i := range newSlice {
			newSlice[i] = (*comb)[i]
		}
		*combs = append(*combs, newSlice)
		return
	}

	for i := start; i <= n; i++ {
		*comb = append(*comb, i)
		backtracking(combs, comb, i+1, n, k)
		*comb = (*comb)[:len(*comb)-1]
	}
}

package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(5))
}

// the search finds the largest value of x for which !isBadVersion is false. Thus, the next value k = x + 1 is the
// smallest possible value for which isBadVersion is true
func firstBadVersion(n int) int {
	largestGoodVersion := -1
	for i := n; i >= 1; i /= 2 {
		for !isBadVersion(largestGoodVersion + i) {
			largestGoodVersion += i
		}
	}

	return largestGoodVersion + 1
}

func isBadVersion(version int) bool {
	if version >= 4 {
		return true
	}
	return false
}

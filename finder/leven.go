package finder

import "fmt"

// So if the query exists we would not use to Levenschtein Distance
// because of the direct look-up, otherwise, we can run all the search term
// across the look-up dictionary. Now just for the sake direction,
// we are not going to load up the whole dictionary. Instead,
// we could use a data-structure that can keeps the Soundex-encoding of the name
// of each file inside a tree. So if a use search contains R when encoded, we could just
// parse the children of R instead, similar to sharding. And then based on the edit distance
// of each encoding in relation to the search term, we'll show the possible docs

func Levenschtein(s1, s2 string) {
	fmt.Println("Levenschteing distance")

	// initialising 2D slice -> rows = len(s1) + 1, cols = len(s2) + 1
	dp := make([][]int, len(s1)+1)
	fmt.Println("[?] initialised matrix\n\n", dp)
	for i := range dp {
		fmt.Printf("[?] new row -> %d\n", dp[i])
		dp[i] = make([]int, len(s2)+1)
	}

	// first row/column
	l1 := len(s1)
	for i := 0; i <= l1; i++ {
		dp[i][0] = i
	}

	l2 := len(s2)
	for j := 0; j < l2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost = 1
			}

			dp[i][j] = min(
				dp[i-1][j]+1,
				dp[i][j-1]+1,
				dp[i-1][j-1]+cost,
			)

		}
	}

	fmt.Println("[?] finalised matrix", dp)
	fmt.Println("{**} final answer>>", dp[l1][l2])
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}


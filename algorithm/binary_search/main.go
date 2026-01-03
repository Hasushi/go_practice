package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(in, &N, &K)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}

	// P(i) (A(i) >= K)
	lo, hi := 0, N
	// lo: 条件に満たさないことが確定したときの最大位置
	// hi: 条件を満たしたことが確定したときの最小位置
	for lo < hi {
		mid := (lo + hi) / 2
		if A[mid] >= K {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	if lo == N {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, hi)
	}
}

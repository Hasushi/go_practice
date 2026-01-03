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

	lo, hi := 0, N
	for lo < hi {
		mid := (lo + hi) / 2
		if A[mid] < K {
			lo = mid
		} else {
			hi = mid
		}
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	if lo == N {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, lo)
	}
}

package main

import (
	"math"
)

const MOD int64 = 1_000_000_007

func modPow(a, e int64) int64 {
	res := int64(1)
	a %= MOD

	for e > 0 {
		if e&1 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		e >>= 1
	}
	return res
}

func xorAfterQueries(nums []int, queries [][]int) int {
	n := len(nums)

	B := int(math.Sqrt(float64(n))) + 1

	mul := make([]int64, n)
	for i := range mul {
		mul[i] = 1
	}

	// small[k] stores queries with step k
	small := make([][][]int, B+1)

	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]

		if k > B {
			// brute force
			for i := l; i <= r; i += k {
				mul[i] = mul[i] * int64(v) % MOD
			}
		} else {
			small[k] = append(small[k], q)
		}
	}

	for k := 1; k <= B; k++ {
		if len(small[k]) == 0 {
			continue
		}

		// group by remainder
		byRem := make([][][]int, k)
		for _, q := range small[k] {
			rem := q[0] % k
			byRem[rem] = append(byRem[rem], q)
		}

		for rem := 0; rem < k; rem++ {
			// count length of this sequence
			length := 0
			for idx := rem; idx < n; idx += k {
				length++
			}

			diff := make([]int64, length+1)
			for i := range diff {
				diff[i] = 1
			}

			for _, q := range byRem[rem] {
				l, r, v := q[0], q[1], int64(q[3])

				start := (l - rem) / k
				end := (r - rem) / k

				diff[start] = diff[start] * v % MOD

				if end+1 < length {
					inv := modPow(v, MOD-2)
					diff[end+1] = diff[end+1] * inv % MOD
				}
			}

			cur := int64(1)
			pos := 0

			for idx := rem; idx < n; idx += k {
				cur = cur * diff[pos] % MOD
				mul[idx] = mul[idx] * cur % MOD
				pos++
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		val := int((int64(nums[i]) * mul[i]) % MOD)
		ans ^= val
	}

	return ans
}
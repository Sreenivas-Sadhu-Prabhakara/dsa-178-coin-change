package main

import "fmt"

// MinCoins returns the fewest coins (unlimited supply of each denomination)
// summing to amount, or -1 when it cannot be made. Unbounded knapsack DP:
// amounts swept upward so a coin may be reused.
func MinCoins(coins []int, amount int) int {
	const inf = int(^uint(0) >> 1)
	dp := make([]int, amount+1)
	for a := 1; a <= amount; a++ {
		dp[a] = inf
		for _, c := range coins {
			if c <= a && dp[a-c] != inf && dp[a-c]+1 < dp[a] {
				dp[a] = dp[a-c] + 1
			}
		}
	}
	if dp[amount] == inf {
		return -1
	}
	return dp[amount]
}

func main() {
	fmt.Println(MinCoins([]int{1, 2, 5}, 11))
	fmt.Println(MinCoins([]int{1, 3, 4}, 6)) // greedy would say 3; DP says 2
}

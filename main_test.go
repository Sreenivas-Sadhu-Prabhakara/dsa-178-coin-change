package main

import "testing"

func TestMinCoins(t *testing.T) {
	cases := []struct {
		coins  []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{1, 3, 4}, 6, 2}, // 3+3: the greedy counterexample
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{5, 10}, 7, -1},
		{[]int{186, 419, 83, 408}, 6249, 20},
	}
	for _, c := range cases {
		if got := MinCoins(c.coins, c.amount); got != c.want {
			t.Errorf("MinCoins(%v,%d) = %d, want %d", c.coins, c.amount, got, c.want)
		}
	}
}

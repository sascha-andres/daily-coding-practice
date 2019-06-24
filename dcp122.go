package daily_coding_practice

import (
	"math"
)

// Run122 You are given a 2-d matrix where each cell represents number of coins in that cell. Assuming we start at matrix[0][0], and can only move right or down, find the maximum number of coins you can collect by the bottom right corner.
func Run122(coins [][]int) int {
	if 0 == len(coins) {
		return 0
	}

	width := len(coins[0])
	if 0 == width {
		return 0
	}

	for _, val := range coins {
		if len(val) != width {
			return 0
		}
	}

	return int(math.Max(getSumOfRight(coins, 0, 0),
		getSumOfDown(coins, 0, 0)))
}

func getSumOfRight(coins [][]int, x, y int) float64 {
	if y >= len(coins[0])-1 {
		return float64(coins[x][y])
	}
	return float64(coins[x][y]) + math.Max(getSumOfRight(coins, x, y+1),
		getSumOfDown(coins, x, y+1))
}

func getSumOfDown(coins [][]int, x, y int) float64 {
	if x >= len(coins)-1 {
		return float64(coins[x][y])
	}
	return float64(coins[x][y]) + math.Max(getSumOfRight(coins, x+1, y),
		getSumOfDown(coins, x+1, y))
}

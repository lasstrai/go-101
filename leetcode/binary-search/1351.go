package main

func countNegatives(grid [][]int) int {
	count := 0
	for i, _ := range grid {
		current := grid[i]
		flag := true
		j := len(current) - 1
		for j >= 0 && flag {
			if current[j] < 0 {
				count++
			} else {
				flag = false
			}
			j--
		}
	}
	return count
}

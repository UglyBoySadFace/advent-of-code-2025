package main

import (
	"fmt"
)

// 987654321111111
// 811111111111119
// 234234234234278
// 818181911112111
// The batteries are arranged into banks; each line of digits in your input corresponds to a single bank of batteries. Within each bank, you need to turn on exactly two batteries; the joltage that the bank produces is equal to the number formed by the digits on the batteries you've turned on. For example, if you have a bank like 12345 and you turn on batteries 2 and 4, the bank would produce 24 jolts. (You cannot rearrange batteries.)
func main() {

	jolts := 0
	var line string
	var lines []string
	for true {
		_, err := fmt.Scanf("%s", &line)
		// save the line
		lines = append(lines, line)
		if err != nil {
			break
		}
	}
	for _, line := range lines {
		fmt.Printf("Processing line: %s\n", line)
		result := getMaxJolts(line)
		fmt.Printf("Max jolts for this line: %d\n", result)
		jolts += result
	}

	fmt.Println("\n", jolts)
}

func getMaxJolts(line string) int {
	n := len(line)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 13)
	}

	for i := 1; i <= n; i++ {
		digit := int(line[i-1] - '0')
		for j := 0; j <= 12; j++ {
			// Skip current battery
			dp[i][j] = dp[i-1][j]

			// Pick current battery
			if j > 0 {
				newValue := dp[i-1][j-1]*10 + digit
				if newValue > dp[i][j] {
					dp[i][j] = newValue
				}
			}
		}
	}

	return dp[n][12]
}

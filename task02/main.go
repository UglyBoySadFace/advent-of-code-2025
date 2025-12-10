package main

import (
	"fmt"
	"strconv"
)

func main() {
	var from, to int
	total := 0
	for true {
		_, err := fmt.Scanf("%d-%d", &from, &to)
		if err != nil {
			break
		}

		for from <= to {
			if hasMatchingSequences(from) {
				fmt.Println("Adding:", from)
				total += from
			}
			from++
		}
		fmt.Println("Current total:", total)
	}
}

func hasMatchingSequences(num int) bool {
	s := strconv.Itoa(num)
	len := len(s)
	for seq_len := 1; seq_len <= len/2; seq_len++ {
		if len%seq_len != 0 {
			continue
		}
		matches := true
		for j := 0; j < len-seq_len; j++ {
			if s[j] != s[j+seq_len] {
				matches = false
				break
			}
		}
		if matches {
			return true
		}
	}
	return false
}

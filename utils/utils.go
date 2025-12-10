package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadLines reads a file line by line and returns a slice of strings
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// ReadFile reads entire file content as a string
func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// MustReadLines reads lines or panics on error
func MustReadLines(filename string) []string {
	lines, err := ReadLines(filename)
	if err != nil {
		panic(fmt.Sprintf("failed to read file %s: %v", filename, err))
	}
	return lines
}

// MustReadFile reads file content or panics on error
func MustReadFile(filename string) string {
	content, err := ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("failed to read file %s: %v", filename, err))
	}
	return content
}

// ToIntSlice converts a slice of strings to a slice of integers
func ToIntSlice(strs []string) ([]int, error) {
	ints := make([]int, len(strs))
	for i, s := range strs {
		num, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return nil, err
		}
		ints[i] = num
	}
	return ints, nil
}

// MustToIntSlice converts strings to ints or panics on error
func MustToIntSlice(strs []string) []int {
	ints, err := ToIntSlice(strs)
	if err != nil {
		panic(fmt.Sprintf("failed to convert to int slice: %v", err))
	}
	return ints
}

// Abs returns absolute value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Sum returns the sum of all integers in a slice
func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

package main

import "fmt"

func main() {
	dial := 50
	counter := 0
	var dir uint8
	var count int
	for true {
		_, err := fmt.Scanf("%c%d", &dir, &count)
		if err != nil {
			break
		}
		switch dir {
		case 'L':
			dial = (dial - count) % 100
		case 'R':
			dial = (dial + count) % 100
		}

		if dial == 0 {
			counter++
		}
	}

	fmt.Println("Final counter:", counter)
}

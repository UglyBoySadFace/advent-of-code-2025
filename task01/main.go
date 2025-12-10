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

		counter += count / 100
		count = count % 100

		if meanwhileClicked(dial, dir, count) {
			counter++
		}

		switch dir {
		case 'L':
			dial = (dial - count + 100) % 100
		case 'R':
			dial = (dial + count) % 100
		}

		fmt.Println("Dial:", dial, "Counter:", counter)

	}

	fmt.Println("Final counter:", counter)
}

func meanwhileClicked(dial int, dir uint8, count int) bool {
	if dial == 0 {
		return false
	}
	switch dir {
	case 'L':
		if dial-count <= 0 {
			return true
		}
	case 'R':
		if dial+count >= 100 {
			return true
		}
	}
	return false
}

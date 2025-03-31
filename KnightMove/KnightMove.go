package main

import "fmt"

// isValidKnightMove checks if the move from (x1, y1) to (x2, y2) is a valid knight move.
func isValidKnightMove(x1, y1, x2, y2 int) bool {
	dx, dy := x2-x1, y2-y1
	return (dx == 2 && dy == 1) || (dx == 2 && dy == -1) ||
		(dx == -2 && dy == 1) || (dx == -2 && dy == -1) ||
		(dx == 1 && dy == 2) || (dx == 1 && dy == -2) ||
		(dx == -1 && dy == 2) || (dx == -1 && dy == -2)
}

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)

	if isValidKnightMove(x1, y1, x2, y2) {
		fmt.Println("true")
		return
	}
	fmt.Println("false")
}

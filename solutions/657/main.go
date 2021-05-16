package main

func judgeCircle(moves string) bool {
	// x and y position
	x, y := 0, 0

	// go thru moves
	for i := 0; i < len(moves); i++ {
		// switch character
		switch moves[i] {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		default:
			x++
		}
	}

	// check if back at origin
	return x == 0 && y == 0
}

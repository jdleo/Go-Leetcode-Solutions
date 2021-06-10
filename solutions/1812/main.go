package main

func squareIsWhite(coordinates string) bool {
	i, j := coordinates[0]-'a', coordinates[1]-'1'
	return (i&1 == 0 && j&1 == 1) || (i&1 == 1 && j&1 == 0)
}

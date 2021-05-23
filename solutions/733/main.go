package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	// make sure old color != newColor
	if image[sr][sc] != newColor {
		// start fill
		dfsFill(image, sr, sc, newColor, image[sr][sc])
	}
	return image
}

// helper method to recursively fill
func dfsFill(image [][]int, i, j, newColor, oldColor int) {
	// bounds check
	if i >= 0 && i < len(image) && j >= 0 && j < len(image[i]) && image[i][j] == oldColor {
		// fill
		image[i][j] = newColor
		// up, down, left, right
		dfsFill(image, i-1, j, newColor, oldColor)
		dfsFill(image, i+1, j, newColor, oldColor)
		dfsFill(image, i, j-1, newColor, oldColor)
		dfsFill(image, i, j+1, newColor, oldColor)
	}
}

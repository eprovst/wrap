package pdf

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func getHeight(lines []styledLine) int {
	height := 0
	for _, line := range lines {
		height += line.leading() + 1
	}
	return height
}

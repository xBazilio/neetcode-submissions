func isValidSudoku(board [][]byte) bool {
	var squareMems [3]map[byte]int
	for i, line := range board {
		// create new lineMem each row
		lineMem := make(map[byte]int)
		// create new squareMems each 3 rows
		if i == 0 || (i-1)/3 != i/3 {
			squareMems = [3]map[byte]int{
				make(map[byte]int),
				make(map[byte]int),
				make(map[byte]int),
			}
		}
		for j, b := range line {

			if b != '.' {
				// check line
				lineMem[b]++
				if lineMem[b] > 1 {
					return false
				}
				// check squares
				sqi := j / 3
				squareMems[sqi][b]++
				if squareMems[sqi][b] > 1 {
					return false
				}
			}

		}
	}

	// check columns
	for i := range 9 {
		// create new lineMem each column
		lineMem := make(map[byte]int)
		for j := range 9 {
			b := board[j][i]
			if b != '.' {
				lineMem[b]++
				if lineMem[b] > 1 {
					return false
				}
			}
		}
	}
	return true
}

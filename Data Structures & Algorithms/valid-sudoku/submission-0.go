func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]struct{}, 9)
	cols := make([]map[byte]struct{}, 9)
	boxes := make([]map[byte]struct{}, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]struct{})
		cols[i] = make(map[byte]struct{})
		boxes[i] = make(map[byte]struct{})
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			item := board[i][j]
			if item == '.' {
				continue
			}

			box := (i/3)*3 + j/3

			if _, ok := rows[i][item]; ok {
				return false
			}
			if _, ok := cols[j][item]; ok {
				return false
			}
			if _, ok := boxes[box][item]; ok {
				return false
			}

			rows[i][item] = struct{}{}
			cols[j][item] = struct{}{}
			boxes[box][item] = struct{}{}
		}
	}

	return true
}

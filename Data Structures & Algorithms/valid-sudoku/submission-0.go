import "slices"

type set struct {
	arr []byte
	hash map[byte]struct{}
}

func (s *set) Insert(b byte) bool {
	if b == '.' {
		return true
	}
	_, ok := s.hash[b]
	if ok {
		return false
	}
	if s.hash == nil {
		s.hash = make(map[byte]struct{})	
	}
	
	s.hash[b] = struct{}{}
	s.arr = append(s.arr, b)

	return true
}

func isValidSudoku(board [][]byte) bool {
	rows := make([]set, 9)
	cols := make([]set, 9)
	grid := make([]set, 9)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			value := board[row][col]
			
			gridindex := row / 3 * 3 + col / 3
			results := []bool{
				grid[gridindex].Insert(value),
				rows[row].Insert(value),
				cols[col].Insert(value),
			}

			if slices.Contains(results, false) {
				return false
			}
		}
	}

	return true
}

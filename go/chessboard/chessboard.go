package chessboard

// Rank stores if a square is occupied by a piece
type Rank []bool

// Chessboard contains a map of eight Ranks
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) (count int) {
	if squares, ok := cb[rank]; ok {
		for _, square := range squares {
			if square {
				count++
			}
		}
	}
	return
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (count int) {
	index := file - 1
	for _, rank := range cb {
		if index < len(rank) && rank[index] {
			count++
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) (count int) {
	for _, rank := range cb {
		count += len(rank)
	}
	return
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (count int) {
	for rank := range cb {
		count += CountInRank(cb, rank)
	}
	return
}

package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from 1 to 8
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	counter := 0
	for _, v := range cb[rank] {
		if v {
			counter++
		}
	}
	return counter
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	counter := 0
	for _, rank := range cb {
		if file <= len(rank) && rank[file-1] {
			counter++
		}
	}
	return counter
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	counter := 0
	for range cb {
		counter += len(cb)
	}
	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	counter := 0
	for _, rank := range cb {
		for _, square := range rank {
			if square {
				counter++
			}
		}
	}
	return counter
}

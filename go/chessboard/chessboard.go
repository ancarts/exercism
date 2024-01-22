package chessboard

type File []bool

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type Chessboard map[string]File

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
	_, ok := cb[file]
	if !ok {
		return 0
	}
	for _, v := range cb[file] {
		if v == true {
			count++

		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	if rank < 1 || rank > 8 {
		return 0
	}
	for _, j := range cb {
		if j[rank-1] == true {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, j := range cb {
		count = len(cb) * len(j)
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		for _, v := range file {
			if v == true {
				count++
			}
		}
	}
	return count
}

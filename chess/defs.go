package chess

const (
	// PIECE TYPES
	empty = 0
	wpawn byte = 1
	bpawn byte = 2
	king byte = 3
	knight byte = 4
	bishop byte = 5
	rook byte = 6
	queen byte = 7

	// SIDES
	white int8 = 1
	black int8 = -1

	// PIECE MOVE OFFSETS REPRESENTATION:
	up        = -16 // WE RESERVED 8 BITS FOR OFFBOARD DETECTION FUNCTION
	down      = 16
	left      = -1
	right     = 1
	upRight   = -15 // UP + RIGHT
	downRight = 17 // DOWN + RIGHT
	downLeft  = 15 // DOWN + LEFT
	upLeft    = -17 // UP + LEFT
)

var (
	moveVectors = map[byte][]int {
		king: []int{upRight, right, downRight, down, downLeft, left, upLeft, up},
		knight: []int{(up << 1) + right, (right << 1) + up, (right << 1) + down, (down << 1) + right,
			(down << 1) + left, (left << 1) + down, (left << 1) + up, (up << 1) + left},
		bishop: []int{upRight, downRight, downLeft, upLeft},
		rook: []int{right, down, left, up},
		queen: []int{upRight, right, downRight, down, downLeft, left, upLeft, up},
	}
	pawnCaptureVectors = map[byte][]int {
		wpawn: []int{upLeft, upRight},
		bpawn: []int{downLeft, downRight},
	}
	pawnMoveVectors = map[byte][]int {
		wpawn: []int{up, up<<1},
		bpawn: []int{down, down<<1},
	}
)

// SQUARE REPRESENTATION: 1[OFFBOARD] 111[RANK] 1[OFFBOARD] 111[FILE]
func isOnBoard(square int) bool {
	return square&0x88 == 0
}

// PIECE REPRESENTATION: 0000 1[SIDE] 111[TYPE]
func getPieceSide(piece byte) int8 {
	if piece >> 3 == 0 {
		return 1
	}
	return -1
}

// TYPE IS REPRESENTED BY 3 LAST BITS, 7 IS A MAX VALUE 
func getPieceType(piece byte) byte {
	return piece & 7 
}

// PIECE GROUP REPRESENTATION: LEAPER = FALSE, RANGE = TRUE
func isRangePiece(piece byte) bool {
	return getPieceType(piece) > 4 // [1-4] LEAPERS, [5-7] RANGE PIECES 
}
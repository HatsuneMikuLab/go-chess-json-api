package chess

type Square int16
type Piece byte
type Side int8 // 1 FOR WHITE AND -1 FOR BLACK
type Move [2]Square
type PiecesMap map[Square]Piece
type KingsMap map[Side]Square

const (
	// PIECE TYPES
	empty = 0

	pawn Piece = 0b1
	king Piece = 0b10
	knight Piece = 0b11
	bishop Piece = 0b100
	rook Piece = 0b101
	queen Piece = 0b110

	wpawn Piece = 0b1001
	wking Piece = 0b1010
	wknight Piece = 0b1011
	wbishop Piece = 0b1100
	wrook Piece = 0b1101
	wqueen Piece = 0b1110

	bpawn Piece = 0b10001
	bking Piece = 0b10010
	bknight Piece = 0b10011
	bbishop Piece = 0b10100
	brook Piece = 0b10101
	bqueen Piece = 0b10110

	// SIDES
	white Side = 1
	black Side = -1

	// PIECE MOVE OFFSETS REPRESENTATION:
	up Square        = -16 // WE RESERVED 8 BITS FOR OFFBOARD DETECTION FUNCTION. WHEN WE MOVE FORWARD OR BACKWARD WE SKIP ROW
	down Square      = 16
	left Square     = -1
	right Square     = 1
	upRight Square   = -15 // UP + RIGHT
	downRight Square = 17 // DOWN + RIGHT
	downLeft Square  = 15 // DOWN + LEFT
	upLeft Square    = -17 // UP + LEFT
)

var (
	rangePieces = map[Piece]bool{ queen: true, rook: true, bishop: true }
	// BLACK AND WHITE USE SAME VECTORS FOR ALL PIECES EXCEPT PAWNS
	moveVectors = map[Piece][]Square {
		king: []Square{upRight, right, downRight, down, downLeft, left, upLeft, up},
		knight: []Square{(up << 1) + right, (right << 1) + up, (right << 1) + down, (down << 1) + right,
			(down << 1) + left, (left << 1) + down, (left << 1) + up, (up << 1) + left},
		bishop: []Square{upRight, downRight, downLeft, upLeft},
		rook: []Square{right, down, left, up},
		queen: []Square{upRight, right, downRight, down, downLeft, left, upLeft, up},
	}
	pawnCaptureVectors = map[Piece][]Square {
		wpawn: []Square{upLeft, upRight},
		bpawn: []Square{downLeft, downRight},
	}
	pawnMoveVectors = map[Piece][]Square {
		wpawn: []Square{up, up<<1},
		bpawn: []Square{down, down<<1},
	}
	startPosition = PiecesMap{
		0: brook, 1: bknight, 2: bbishop, 3: bqueen, 4: bking, 5: bbishop, 6: bknight, 7: brook,
		16: bpawn, 17: bpawn, 18: bpawn, 19: bpawn, 20: bpawn, 21: bpawn, 22: bpawn, 23: bpawn,
		96: wpawn, 97: wpawn, 98: wpawn, 99: wpawn, 100: wpawn, 101: wpawn, 102: wpawn, 103: wpawn,
		112: wrook, 113: wknight, 114: wbishop, 115: wqueen, 116: wking, 117: wbishop, 118: wknight, 119: wrook,
	}
	pieceNames = map[Piece]string {
		empty: "--",
		wpawn: "WP",
		wking: "WK",
		wknight: "WN",
		wbishop: "WB",
		wrook: "WR",
		wqueen: "WQ",
		bpawn: "BP",
		bking: "BK",
		bknight: "BN",
		bbishop: "BB",
		brook: "BR",
		bqueen: "BQ",
	}
)

// SQUARE REPRESENTATION: 1[OFFBOARD] 111[RANK] 1[OFFBOARD] 111[FILE] (TOTAL: 8 bits)
func isOnBoard(square Square) bool {
	return square&0x88 == 0
}

func getSquareIndex(file int, rank int) Square {
	return Square((7 - rank) << 4 + file)
}

// CONVERT SQUARE INDEX TO NAME
func index2name(square Square) string {
	return string([]rune{rune(square % 8 + 'A'), 8 - rune('8' - square / 16)})
}

func name2index(square string) Square {
	return Square(0 - (square[1] - '8') * 16 + (square[0] - 'a') % 16)
}

// PIECE REPRESENTATION: 000 11[SIDE] 111[TYPE] (TOTAL: 5 bits)
func getPieceSide(piece Piece) Side {
	if piece >> 3 == 1 {
		return 1
	} else if piece >> 3 == 2 {
		return -1
	} else {
		return 0
	}
}

// PIECE REPRESENTATION: 000 11[SIDE] 111[TYPE] (TOTAL: 5 bits)
func getPieceType(piece Piece) Piece {
	return piece & 0b111 
}




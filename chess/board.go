package chess

type PiecesMap map[int]byte
type KingsMap map[int8]int
//type CastlePerm map[int8]bool
type Move [2]int

type Board struct {
	Pieces        PiecesMap
	Kings         KingsMap
	MovesNext     int8
	CastlePerm    byte
	EnPassant     byte
	HalfmoveClock byte
}

func SetupStartPosition() *Board {
	board := &Board{}
	board.Pieces = PiecesMap{
		112: brook, 113: bknight, 114: bbishop, 115: bqueen, 116: bking, 117: bbishop, 118: bknight, 119: brook,
		96: bpawn, 97: bpawn, 98: bpawn, 99: bpawn, 100: bpawn, 101: bpawn, 102: bpawn, 103: bpawn,
		16: wpawn, 17: wpawn, 18: wpawn, 19: wpawn, 20: wpawn, 21: wpawn, 22: wpawn, 23: wpawn,
		0: wrook, 1: wknight, 2: wbishop, 3: wqueen, 4: wking, 5: wbishop, 6: wknight, 7: wrook,
	}
	board.Kings[white] = 4
	board.Kings[black] = 116
	board.MovesNext = white
	board.CastlePerm = 0b1111
	return board
}

func (b *Board) forwardMove(move Move) byte { // [0]FROM [1]TO
	if b.Pieces[move[0]] == king {
		b.Kings[b.MovesNext] = move[1]
	}
	capturedPiece := b.Pieces[move[1]]
	b.Pieces[move[1]] = b.Pieces[move[0]]
	b.Pieces[move[0]] = empty
	b.MovesNext = -1 * b.MovesNext
	return capturedPiece 
}

func (b *Board) undoMove(move Move, capturedPiece byte) { // [0]FROM [1]TO
	if b.Pieces[move[1]] == king {
		b.Kings[b.MovesNext] = move[0]
	}
	b.Pieces[move[0]] = b.Pieces[move[1]]
	b.MovesNext = -1 * b.MovesNext
	if capturedPiece == empty {
		delete(b.Pieces, move[1])
	} else {
		b.Pieces[move[1]] = capturedPiece
	}
}

func (b *Board) isAttacked(square int) bool {
	for pieceType, offsetSlice := range moveVectors {
		for _, offset := range offsetSlice {
			targetSquare := square + offset
			for isOnBoard(targetSquare) && b.Pieces[targetSquare] == empty {
				targetSquare += offset
			}
			if isOnBoard(targetSquare) && 
			getPieceSide(b.Pieces[targetSquare]) == b.MovesNext &&
			getPieceType(b.Pieces[targetSquare]) == pieceType {
				return true
			}
		}
	}
	opponentPawns := bpawn
	if b.MovesNext == black {
		opponentPawns = wpawn
	}
	for _, offset := range pawnCaptureVectors[opponentPawns] {
		targetSquare := square + offset
		if isOnBoard(targetSquare) && 
		getPieceSide(b.Pieces[targetSquare]) == b.MovesNext {
			return true
		}
	}
	return false
}

func (b *Board) GenAllowedMoves() []Move {
	pseudoMoves := make([]Move, 0, 200)
	pseudoCaptures := make([]Move, 0, 200)
	output := make([]Move, 0, 200)

	for square, piece := range b.Pieces {
		if getPieceSide(piece) != b.MovesNext {
			continue
		}
		maxDistance := 1
		if isRangePiece(piece) {
			maxDistance = 7
		}
		// FIND ALL PSEUDO-VALID MOVES AND CAPTURES FOR ALL PIECE TYPES EXCEPT PAWN
		for _, offset := range moveVectors[getPieceType(piece)] {
			targetSquare := square + offset
			
			for isOnBoard(targetSquare) && maxDistance > 0 {
				if b.Pieces[targetSquare] == empty {
					pseudoMoves = append(pseudoMoves, Move{square, targetSquare})
				} else if getPieceSide(b.Pieces[targetSquare]) != b.MovesNext {
					pseudoCaptures = append(pseudoCaptures, Move{square, targetSquare})
					break;
				} else {
					break;
				}
				targetSquare += offset
				maxDistance--;
			}
		}
		// FIND ALL PSEUDO-VALID PAWN CAPTURES
		for _, offset := range pawnCaptureVectors[piece] {
			targetSquare := square + offset
			if b.Pieces[targetSquare] != empty && getPieceSide(b.Pieces[targetSquare]) != b.MovesNext {
				pseudoCaptures = append(pseudoCaptures, Move{square, targetSquare})
			}
		}
		// FIND ALL PSEUDO-VALID PAWN MOVES
		for _, offset := range pawnMoveVectors[piece] {
			targetSquare := square + offset
			if b.Pieces[targetSquare] == empty {
				pseudoMoves = append(pseudoMoves, Move{square, targetSquare})
			} else {
				break
			}
		}
	}
	// CHECK IS KING SAFE
	for _, move := range append(pseudoCaptures, pseudoMoves...) {
		capturedPiece := b.forwardMove(move)
		if !b.isAttacked(b.Kings[-1 * b.MovesNext]) {
			output = append(output, move)
		}
		b.undoMove(move, capturedPiece)
	}
	return output
}


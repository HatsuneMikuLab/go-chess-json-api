package chess

type PiecesMap map[int]byte
type KingsMap map[int8]int
type Move [2]int

type Board struct {
	Pieces        PiecesMap
	Kings         KingsMap
	MovesNext     int8
	CastlePerm    byte
	EnPassant     byte
	HalfmoveClock byte
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

func (b *Board) GenAllowedMoves() {
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
		for index, offset := range moveVectors[getPieceType(piece)] {
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
		for _, offset := range pawnCaptureVectors[getPieceType(piece)] {
			targetSquare := square + offset
			if b.Pieces[targetSquare] != empty && getPieceSide(b.Pieces[targetSquare]) =! b.MovesNext {
				pseudoCaptures = append(pseudoCaptures, Move{square, targetSquare})
			}
		}
		// FIND ALL PSEUDO-VALID PAWN MOVES
		for _, offset := range pawnMoveVectors[getPieceType(piece)] {
			targetSquare := square + offset
			if b.Pieces[targetSquare] == empty {
				pseudoMoves = append(pseudoMoves, Move{square, targetSquare})
			} else {
				break
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
}


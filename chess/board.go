package chess

type PiecesMap map[int]byte

type Board struct {
	Pieces        PiecesMap
	Kings         [2]byte
	MovesNext     byte
	CastlePerm    byte
	EnPassant     byte
	HalfmoveClock byte
}

func (b *Board) ForwardMove(move [2]int) byte { // [0]FROM [1]TO
	capturedPiece := b.Pieces[move[1]]
	b.Pieces[move[1]] = b.Pieces[move[0]]
	b.Pieces[move[0]] = empty
	return capturedPiece 
}

func (b *Board) UndoMove(move [2]int, capturedPiece byte) { // [0]FROM [1]TO
	b.Pieces[move[0]] = b.Pieces[move[1]]
	if capturedPiece == empty {
		delete(b.Pieces, move[1])
	} else {
		b.Pieces[move[1]] = capturedPiece
	}
}

func (b *Board) GenAllowedMoves() {
	movesWithoutCapture := make([][2]int, 0, 200)
	movesWithCapture := make([][2]int, 0, 100)
	for square, piece := range b.Pieces {
		if getPieceSide(piece) != b.MovesNext {
			continue
		}
		// FIND ALL PSEUDO-VALID MOVES AND CAPTURES FOR ALL PIECE TYPES EXCEPT PAWN
		for index, offset := range moveVectors[getPieceType(piece)] {
			targetSquare := square + offset
			maxDistance := 1
			if isRangePiece(piece) {
				maxDistance = 7
			}
			
			for isOnBoard(targetSquare) && maxDistance > 0 {
				if b.Pieces[targetSquare] == empty {
					movesWithoutCapture = append(movesWithoutCapture, [2]int{square, targetSquare})
				} else if getPieceSide(b.Pieces[targetSquare]) != b.MovesNext {
					movesWithCapture = append(movesWithCapture, [2]int{square, targetSquare})
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
				movesWithCapture = append(movesWithCapture, [2]int{square, targetSquare})
			}
		}
		// FIND ALL PSEUDO-VALID PAWN MOVES
		for _, offset := range pawnMoveVectors[getPieceType(piece)] {
			targetSquare := square + offset
			if b.Pieces[targetSquare] == empty {
				movesWithoutCapture = append(movesWithoutCapture, [2]int{square, targetSquare})
			} else {
				break
			}
		}
	}
}


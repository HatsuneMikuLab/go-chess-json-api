package chess

type Board struct {
	Pieces        PiecesMap `json:"pieces"`
	Kings         KingsMap `json:"kings"`
	MovesNext     Side `json:"moves_next"`
	CastlePerm    byte `json:"castle_perm"`
	EnPassant     byte `json:"en_passant"`
	HalfmoveClock byte `json:"halfmove_clock"`
}

func SetupStartPosition() *Board {
	board := &Board{}
	board.Pieces = startPosition
	board.Kings = KingsMap{
		white: 4, //E1
		black: 116, //E8
	}
	board.MovesNext = white
	board.CastlePerm = 0b1111
	return board
}

func RenderBoard() map[string]string {
	
}

func (b *Board) ForwardMove(move Move) Piece { // [0]FROM [1]TO
	if b.Pieces[move[0]] == king { // TRACKING KING POSITION
		b.Kings[b.MovesNext] = move[1]
	}
	capturedPiece := b.Pieces[move[1]]
	b.Pieces[move[1]] = b.Pieces[move[0]]
	delete(b.Pieces, move[0])
	b.MovesNext = -1 * b.MovesNext
	return capturedPiece 
}

func (b *Board) UndoMove(move Move, capturedPiece Piece) { // [0]FROM [1]TO
	if b.Pieces[move[1]] == king { // TRACKING KING POSITION
		b.Kings[b.MovesNext] = move[0]
	}
	b.Pieces[move[0]] = b.Pieces[move[1]]
	if capturedPiece == empty {
		delete(b.Pieces, move[1])
	} else {
		b.Pieces[move[1]] = capturedPiece
	}
	b.MovesNext = -1 * b.MovesNext
}

func (b *Board) isAttacked(square Square, bySide Side) bool {
	for pieceType, offsetSlice := range moveVectors {
		for _, offset := range offsetSlice {
			iSquare := square + offset
			for rangePieces[pieceType] && isOnBoard(iSquare) && b.Pieces[iSquare] == empty {
				iSquare += offset
			}
			if getPieceSide(b.Pieces[iSquare]) == bySide && getPieceType(b.Pieces[iSquare]) == pieceType {
				return true
			}
		}
	}

	activePawns := wpawn
	if bySide == black {
		activePawns = bpawn
	}
	for _, offset := range pawnCaptureVectors[activePawns] {
		iSquare := square - offset
		if b.Pieces[iSquare] == activePawns {
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
		if getPieceSide(piece) != b.MovesNext { // YOU CANNOT MOVE OPPONENT's PIECES
			continue
		}
		// FIND ALL PSEUDO-VALID MOVES AND CAPTURES FOR ALL PIECE TYPES EXCEPT PAWN
		for _, offset := range moveVectors[getPieceType(piece)] {
			iSquare := square + offset
			for isOnBoard(iSquare) {
				if b.Pieces[iSquare] == empty {
					pseudoMoves = append(pseudoMoves, Move{square, iSquare})
				} else if getPieceSide(b.Pieces[iSquare]) != b.MovesNext {
					pseudoCaptures = append(pseudoCaptures, Move{square, iSquare})
					break;
				} else {
					break;
				}
				if rangePieces[getPieceType(piece)] {
					iSquare += offset
				} else {
					break;
				}
			}
		}
		// FIND ALL PSEUDO-VALID PAWN CAPTURES
		for _, offset := range pawnCaptureVectors[piece] {
			iSquare := square + offset
			if isOnBoard(iSquare) && getPieceSide(b.Pieces[iSquare]) != b.MovesNext {
				pseudoCaptures = append(pseudoCaptures, Move{square, iSquare})
			}
		}
		// FIND ALL PSEUDO-VALID PAWN MOVES
		for _, offset := range pawnMoveVectors[piece] {
			iSquare := square + offset
			if isOnBoard(iSquare) && b.Pieces[iSquare] == empty {
				pseudoMoves = append(pseudoMoves, Move{square, iSquare})
			} else {
				break
			}
		}
	}
	// CHECK IS KING SAFE
	for _, move := range append(pseudoCaptures, pseudoMoves...) {
		capturedPiece := b.ForwardMove(move)
		if !b.isAttacked(b.Kings[-1 * b.MovesNext], b.MovesNext) {
			output = append(output, move)
		}
		b.UndoMove(move, capturedPiece)
	}
	return append(pseudoCaptures, pseudoMoves...)
}


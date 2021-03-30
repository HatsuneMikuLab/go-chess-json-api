package chess

import (
	"reflect"
	"testing"
)

func name2index(square string) int {
	return int(0 - (square[1] - '8') * 16 + (square[0] - 'a') % 16)
}

func TestForwardMove(t *testing.T) {
	board := SetupStartPosition()
	b2b4 := Move{name2index("b2"), name2index("b4")}
	board.ForwardMove(b2b4)
	shouldBe := PiecesMap{
		0: brook, 1: bknight, 2: bbishop, 3: bqueen, 4: bking, 5: bbishop, 6: bknight, 7: brook,
		16: bpawn, 17: bpawn, 18: bpawn, 19: bpawn, 20: bpawn, 21: bpawn, 22: bpawn, 23: bpawn,
		//we made b2-b4 move should be 97 - 16 * 2 = 65
		96: wpawn, 97 + up*2: wpawn, 98: wpawn, 99: wpawn, 100: wpawn, 101: wpawn, 102: wpawn, 103: wpawn,  
		112: wrook, 113: wknight, 114: wbishop, 115: wqueen, 116: wking, 117: wbishop, 118: wknight, 119: wrook,
	}
	if !reflect.DeepEqual(board.Pieces, shouldBe) {
		t.Errorf(
			"TestForwardMove is failed. \n Result after %v should be %v \n but got %v", 
			b2b4, shouldBe, board.Pieces,
		)
	}
}

func TestUndoMove(t *testing.T) {
	board := SetupStartPosition()
	e2e4 := Move{name2index("e2"), name2index("e4")}
	capturedPiece := board.ForwardMove(e2e4)
	board.UndoMove(e2e4, capturedPiece)
	if !reflect.DeepEqual(board.Pieces, startPosition) {
		t.Errorf(
			"TestUndoMove is failed. \n Result after undo move %v should be %v \n but got %v", 
			e2e4, startPosition, board.Pieces,
		)
	}
}
package chess

import (
	"reflect"
	"testing"
)

func TestForcedE4Move(t *testing.T) {
	board := SetupStartPosition()
	board.forwardMove(Move{101, 101 + (up << 1)})
	shouldBe := PiecesMap{
		0: brook, 1: bknight, 2: bbishop, 3: bqueen, 4: bking, 5: bbishop, 6: bknight, 7: brook,
		16: bpawn, 17: bpawn, 18: bpawn, 19: bpawn, 20: bpawn, 21: bpawn, 22: bpawn, 23: bpawn,
		//we made e2-e4 move should be 101 - 16 * 2 = 69
		96: wpawn, 97: wpawn, 98: wpawn, 99: wpawn, 100: wpawn, 69: wpawn, 102: wpawn, 103: wpawn,  
		112: wrook, 113: wknight, 114: wbishop, 115: wqueen, 116: wking, 117: wbishop, 118: wknight, 119: wrook,
	}
	if !reflect.DeepEqual(board.Pieces, shouldBe) {
		t.Errorf(
			"TestForcedE4Move is failed. \n Result after e2-e4 should be \n %v but got %v", 
			shouldBe, board.Pieces,
		)
	}
}
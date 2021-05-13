package chess

import (
	"reflect"
	"testing"
)

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

func TestIsAttacked(t *testing.T) {
	board := SetupStartPosition()
	board.ForwardMove(Move{name2index("a2"), name2index("a3")}) 
	// FAKE MOVE TO MAKE POSITION WITH WHITE KING UNDER THE CHECK
	board.ForwardMove(Move{name2index("d7"), name2index("d2")}) 
	if !board.isAttacked(name2index("e1"), -1 * board.MovesNext) {
		t.Errorf("TestIsAttacked is failed. \n Should return true but got false in Position: %v", board.Render())
	}
	// KILL A PAWN TO PROTECT WHITE KING
	board.ForwardMove(Move{name2index("d1"), name2index("d2")})
	if board.isAttacked(name2index("e1"), board.MovesNext) {
		t.Errorf("TestIsAttacked is failed. \n Should return false but got true in Position: %v", board.Render())
	}
}

func TestGenAllowedMoves(t *testing.T) {
	board := SetupStartPosition()
	allowedMoves := make(map[Move]bool, 200)

	for _, move := range board.GenAllowedMoves() {
		allowedMoves[move] = true
	}
	// ALL POSSIBLE FIRST MOVES
	shouldBe := map[Move]bool {
		{name2index("b1"), name2index("a3")}: true,
		{name2index("b1"), name2index("c3")}: true,
		{name2index("g1"), name2index("f3")}: true,
		{name2index("g1"), name2index("h3")}: true,
		{name2index("a2"), name2index("a3")}: true,
		{name2index("a2"), name2index("a4")}: true,
		{name2index("b2"), name2index("b3")}: true,
		{name2index("b2"), name2index("b4")}: true,
		{name2index("c2"), name2index("c3")}: true,
		{name2index("c2"), name2index("c4")}: true,
		{name2index("d2"), name2index("d3")}: true,
		{name2index("d2"), name2index("d4")}: true,
		{name2index("e2"), name2index("e3")}: true,
		{name2index("e2"), name2index("e4")}: true,
		{name2index("f2"), name2index("f3")}: true,
		{name2index("f2"), name2index("f4")}: true,
		{name2index("g2"), name2index("g3")}: true,
		{name2index("g2"), name2index("g4")}: true,
		{name2index("h2"), name2index("h3")}: true,
		{name2index("h2"), name2index("h4")}: true,
	}
	if !reflect.DeepEqual(allowedMoves, shouldBe) {
		t.Errorf(
			"TestGenAllowedMoves is failed. \n Should be %v but got %v",
			shouldBe, allowedMoves, 
		)
	}
}
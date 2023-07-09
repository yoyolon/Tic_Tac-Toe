package main

import (
	"testing"
)

// (0,0)に"o"を置くテスト
func TestPutToken_00_o(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.Put(0, 0, "o")
	if b.Get(0, 0) != "o" {
		t.Errorf("Failed to put \"o\" on (0,0)")
	}
}

// (0,0)に"x"を置くテスト
func TestPutToken_00_x(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.Put(0, 0, "x")
	if b.Get(0, 0) != "x" {
		t.Errorf("Failed to put \"x\" on (0,0)")
	}
}

// (1,1)に"o"を置くテスト
func TestPutToken_11_o(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.Put(1, 1, "o")
	if b.Get(1, 1) != "o" {
		t.Errorf("Failed to put \"o\" on (1,1)")
	}
}

// (1,1)に"x"を置くテスト
func TestPutToken_11_x(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.Put(1, 1, "x")
	if b.Get(1, 1) != "x" {
		t.Errorf("Failed to put \"x\" on (1,1)")
	}
}

// (2,2)に"x"を置くテスト
func TestPutToken_22_x(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.Put(2, 2, "x")
	if b.Get(2, 2) != "x" {
		t.Errorf("Failed to put \"x\" on (2,2)")
	}
}

// (2,2)に"o"を置くテスト
func TestPutToken_22_o(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.Put(2, 2, "o")
	if b.Get(2, 2) != "o" {
		t.Errorf("Failed to put \"o\" on (2,2)")
	}
}

// (0,0)-(2,2)に順番に"o"を置くテスト
func TestPutToken_o(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.Put(i, j, "o")
			if b.Get(i, j) != "o" {
				t.Errorf("Failed to put \"o\" on (%d,%d)", i, j)
			}
		}
	}
}

// (0,0)-(2,2)に順番に"x"を置くテスト
func TestPutToken_x(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.Put(i, j, "x")
			if b.Get(i, j) != "x" {
				t.Errorf("Failed to put \"x\" on (%d,%d)", i, j)
			}
		}
	}
}

// 駒の存在しない座標は"."を返すテスト
func TestGetToken_empty(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.Get(i, j) != "." {
				t.Errorf("Board must be mepty")
			}
		}
	}
}

// 既に駒が存在する箇所には駒を配置できないテスト
func TestPutToken_not_overwrite(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.Put(i, j, "o")
			b.Put(i, j, "x") // 駒は上書き不可能
			if b.Get(i, j) != "o" {
				t.Errorf("Token can not be overwritten")
			}
		}
	}
}

// 勝敗判定: 勝者なし
// 無地の盤面では勝敗が存在しない
func TestCheckWin_empty(t *testing.T) {
	b := &Board{
		tokens: []int{
			0, 0, 0,
			0, 0, 0,
			0, 0, 0},
	}
	if b.CheckWin() != 0 {
		t.Errorf("This case is not a winner.")
	}
}

// 引き分け01
func TestCheckWin_draw_01(t *testing.T) {
	b := &Board{
		tokens: []int{
			1, 2, 1,
			1, 2, 2,
			2, 1, 1},
	}
	if b.CheckWin() != 0 {
		t.Errorf("This case is not a winner.")
	}
}

// 引き分け02(player1がリーチ)
func TestCheckWin_draw_02(t *testing.T) {
	b := &Board{
		tokens: []int{
			1, 1, 0,
			0, 1, 1,
			1, 0, 0},
	}
	if b.CheckWin() != 0 {
		t.Errorf("This case is not a winner.")
	}
}

// 引き分け03(player2がリーチ)
func TestCheckWin_draw_03(t *testing.T) {
	b := &Board{
		tokens: []int{
			0, 2, 2,
			2, 0, 2,
			2, 2, 0},
	}
	if b.CheckWin() != 0 {
		t.Errorf("This case is not a winner.")
	}
}

// 勝敗判定: プレイヤー1が勝利するケース
// プレイヤー1が勝利(縦)
func TestCheckWin_player1_win_01(t *testing.T) {
	b := &Board{
		tokens: []int{
			1, 0, 0,
			1, 0, 0,
			1, 0, 0},
	}
	if b.CheckWin() != 1 {
		t.Errorf("Player1 should be winner.")
	}
}

// プレイヤー1が勝利(横)
func TestCheckWin_player1_win_02(t *testing.T) {
	b := &Board{
		tokens: []int{
			0, 0, 0,
			1, 1, 1,
			0, 0, 0},
	}
	if b.CheckWin() != 1 {
		t.Errorf("Player1 should be winner.")
	}
}

// プレイヤー1が勝利(斜め)
func TestCheckWin_player1_win_03(t *testing.T) {
	b := &Board{
		tokens: []int{
			1, 0, 0,
			0, 1, 0,
			0, 0, 1},
	}
	if b.CheckWin() != 1 {
		t.Errorf("Player1 should be winner.")
	}
}

// 勝敗判定: プレイヤー2が勝利するケース
// プレイヤー2が勝利(縦)
func TestCheckWin_player2_win_01(t *testing.T) {
	b := &Board{
		tokens: []int{
			0, 0, 2,
			0, 0, 2,
			0, 0, 2},
	}
	if b.CheckWin() != 2 {
		t.Errorf("Player2 should be winner.")
	}
}

// プレイヤー2が勝利(横)
func TestCheckWin_player2_win_02(t *testing.T) {
	b := &Board{
		tokens: []int{
			0, 0, 0,
			0, 0, 0,
			2, 2, 2},
	}
	if b.CheckWin() != 2 {
		t.Errorf("Player2 should be winner.")
	}
}

// プレイヤー2が勝利(斜め)
func TestCheckWin_player2_win_03(t *testing.T) {
	b := &Board{
		tokens: []int{
			0, 0, 2,
			0, 2, 0,
			2, 0, 0},
	}
	if b.CheckWin() != 2 {
		t.Errorf("Player2 should be winner.")
	}
}

// 勝敗判定: 自作のケース
// player1が勝利
func TestCheckWin_random_01(t *testing.T) {
	b := &Board{
		tokens: []int{
			1, 1, 1,
			2, 1, 0,
			2, 0, 2},
	}
	if b.CheckWin() != 1 {
		t.Errorf("Player1 should be winner.")
	}
}

// 引き分け
func TestCheckWin_random_02(t *testing.T) {
	b := &Board{
		tokens: []int{
			1, 1, 2,
			2, 2, 1,
			1, 2, 0},
	}
	if b.CheckWin() != 0 {
		t.Errorf("This case is not a winner.")
	}
}

// player2が勝利
func TestCheckWin_random_03(t *testing.T) {
	b := &Board{
		tokens: []int{
			1, 0, 1,
			2, 2, 2,
			0, 0, 1},
	}
	if b.CheckWin() != 2 {
		t.Errorf("Player2 should be winner.")
	}
}

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
		t.Errorf("....")
	}
}

// (0,0)に"x"を置くテスト
func TestPutToken_00_x(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.Put(0, 0, "x")
	if b.Get(0, 0) != "x" {
		t.Errorf("....")
	}
}

// (1,1)に"o"を置くテスト
func TestPutToken_11_o(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.Put(1, 1, "o")
	if b.Get(1, 1) != "o" {
		t.Errorf("....")
	}
}

// (1,1)に"x"を置くテスト
func TestPutToken_11_x(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.Put(1, 1, "x")
	if b.Get(1, 1) != "x" {
		t.Errorf("....")
	}
}

// (2,2)に"x"を置くテスト
func TestPutToken_22_x(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.Put(2, 2, "x")
	if b.Get(2, 2) != "x" {
		t.Errorf("....")
	}
}

// (2,2)に"o"を置くテスト
func TestPutToken_22_o(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.Put(2, 2, "o")
	if b.Get(2, 2) != "o" {
		t.Errorf("....")
	}
}

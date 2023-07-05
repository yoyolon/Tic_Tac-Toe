package main

import "fmt"

type Board struct {
	tokens []int
}

func (b *Board) board_initializtion() {
	b.tokens = make([]int, 10)
	for i := 0; i < 10; i++ {
		b.tokens[i] = 0
	}
}

func (b *Board) put(x, y int, u string) {
	if x > 2 || x < 0 || y > 2 || y < 0 {
		fmt.Printf("ここだめ〜\n")
		return //未完成
	}

	if b.tokens[x+3*y] != 0 {
		fmt.Printf("駒はもう存在するので、もう一回入力してください\n")
		return //未完成
	}

	if u == "o" {
		b.tokens[x+3*y] = 1
	} else {
		b.tokens[x+3*y] = 2
	}
}

func (b *Board) get(x, y int) string {
	if b.tokens[x+3*y] == 1 {
		return "o"
	} else if b.tokens[x+3*y] == 2 {
		return "x"
	} else if b.tokens[x+3*y] == 0 {
		return "."
	}
	return "error" //未完成
}

func (b *Board) show_board() string {
	var result string
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			result += b.get(x, y)
			result += " "
		}
		result += "\n"
	}
	return result
}

func win() {

}

func main() {
	var x, y int
	board := Board{}

	board.board_initializtion()

	for i := 0; i <= 9; i++ {
		fmt.Printf("Player1: Input (x,y) ")
		fmt.Scanf("%d %d", &x, &y)
		board.put(x, y, "o")
		fmt.Printf(" x,y\n")
		fmt.Printf(board.show_board())

		fmt.Printf("Player2: Input (x,y) ")
		fmt.Scanf("%d %d", &x, &y)
		board.put(x, y, "x")
		fmt.Printf(" x,y\n")
		fmt.Printf(board.show_board())
	}
}

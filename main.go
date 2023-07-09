package main

import "fmt"

type Board struct {
	tokens []int
}

func (b *Board) BoardIinitializtion() { //碁盤を [........] 初期化
	b.tokens = make([]int, 9)
	for i := 0; i < 9; i++ {
		b.tokens[i] = 0
	}
}

func (b *Board) Put(x, y int, u string) int { //int type 座標x, yを入力、int type tokens[]に保存
	if x > 2 || x < 0 || y > 2 || y < 0 { //碁盤の中でわない error
		fmt.Printf("碁盤の中でわないので、もう一回入力してください\n")
		return 1
	}

	if b.tokens[x+3*y] != 0 {
		fmt.Printf("駒はもう存在していますので、もう一回入力してください\n") //重複　error
		return 1
	}

	if u == "o" { //player1
		b.tokens[x+3*y] = 1
		return 0 //順調
	} else if u == "x" { //player2
		b.tokens[x+3*y] = 2
		return 0 //順調
	}
	return 2 //予期以外　error
}

func (b *Board) Get(x, y int) string { //int type tokens[] --> 駒文字
	if b.tokens[x+3*y] == 1 { //player1
		return "o"
	} else if b.tokens[x+3*y] == 2 { //palyer2
		return "x"
	} else if b.tokens[x+3*y] == 0 { //駒は存在しない
		return "."
	}
	return "error" //予期以外
}

func (b *Board) BoardShow() string {
	var result string
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			result += b.Get(x, y)
			result += " "
		}
		result += "\n"
	}
	return result
}

func (b *Board) CheckWin() int { //flag==1 --> player1 win; flag==2 --> play2 win
	var flag int
	flag = 0

	for i := 0; i < 3; i++ {
		if b.tokens[i+0] == b.tokens[i+3] && b.tokens[i+0] == b.tokens[i+6] { //row check
			if b.tokens[i+0] == 1 {
				flag = 1
			} else if b.tokens[i+0] == 2 {
				flag = 2
			}
		} else if b.tokens[0+3*i] == b.tokens[1+3*i] && b.tokens[0+3*i] == b.tokens[2+3*i] { //column check
			if b.tokens[0+3*i] == 1 {
				flag = 1
			} else if b.tokens[0+3*i] == 2 {
				flag = 2
			}
		} else if (b.tokens[0] == b.tokens[4] && b.tokens[0] == b.tokens[8]) || (b.tokens[6] == b.tokens[4] && b.tokens[6] == b.tokens[2]) { //Xcheck 座標(0,0), (1,1), (2,2)or(0,2), (1,1), (2,0)
			if b.tokens[4] == 1 { //座標(1,1)
				flag = 1
			} else if b.tokens[4] == 2 {
				flag = 2
			}
		}
	}

	return flag
}

func main() {
	var x, y int
	board := Board{}

	board.BoardIinitializtion() //碁盤を初期化

	fmt.Printf("please input like: x[space]y\n")

	for i := 0; i < 5; i++ {
		fmt.Printf("Player1: Input (x,y) ") //player1のターン
		fmt.Scanf("%d %d", &x, &y)
		for board.Put(x-1, y-1, "o") == 1 { //errorの場合、もう一度入力
			fmt.Printf(board.BoardShow())
			fmt.Scanf("%d %d", &x, &y)
		}
		fmt.Printf(board.BoardShow())
		if board.CheckWin() == 1 { //player1 win check
			fmt.Printf("play1 win\n")
			break
		}

		if i == 4 { //９回入力した　endします
			fmt.Printf("draw\n")
			break
		}

		fmt.Printf("Player2: Input (x,y) ") //player2のターン
		fmt.Scanf("%d %d", &x, &y)
		for board.Put(x-1, y-1, "x") == 1 { //errorの場合、もう一度入力
			fmt.Printf(board.BoardShow())
			fmt.Scanf("%d %d", &x, &y)
		}
		fmt.Printf(board.BoardShow())
		if board.CheckWin() == 2 { //player2 win check
			fmt.Printf("play2 win\n")
			break
		}
	}
}

package main

type GameBoard struct {
	Board [50][50]int
}

// func GameofLife() {

// 	for x := 0; x < 10; x++ {
// 		for y := 0; y < 10; y++ {
// 			fmt.Print(board[x][y])
// 		}
// 		fmt.Print("\n")
// 	}
// 	fmt.Print("\n")
// 	// fmt.Println(CellEvolve(board, 1, 1))
// 	board = Evolve(board)
// 	for x := 0; x < 10; x++ {
// 		for y := 0; y < 10; y++ {
// 			fmt.Print(board[x][y])
// 		}
// 		fmt.Print("\n")
// 	}
// 	fmt.Print("\n")
// 	// fmt.Println(CellEvolve(board, 1, 1))
// 	board = Evolve(board)
// 	for x := 0; x < 10; x++ {
// 		for y := 0; y < 10; y++ {
// 			fmt.Print(board[x][y])
// 		}
// 		fmt.Print("\n")
// 	}
// }

func Evolve(board [50][50]int) [50][50]int {
	var newBoard [50][50]int
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			newBoard[x][y] = CellEvolve(board, x, y)
		}
	}
	return newBoard
}

func CellEvolve(board [50][50]int, x, y int) int {
	var xCheck, yCheck, xCheckMax, yCheckMax, AliveNeigbours int
	xCheck = -1
	xCheckMax = 1
	yCheck = -1
	yCheckMax = 1
	if x == 0 {
		xCheck = 0
	}
	if x == 49 {
		xCheckMax = 0
	}
	if y == 0 {
		yCheck = 0
	}
	if y == 49 {
		yCheckMax = 0
	}
	AliveNeigbours = 0
	for x2 := xCheck; x2 <= xCheckMax; x2++ {
		for y2 := yCheck; y2 <= yCheckMax; y2++ {
			// fmt.Print(board[x+x2][y+y2])
			// fmt.Println(y + y2)

			if y+y2 == y && x+x2 == x {

			} else {
				if board[x+x2][y+y2] == 1 {
					AliveNeigbours++
				}
			}
		}
		// fmt.Print("\n")
	}
	// fmt.Println(AliveNeigbours)
	if board[x][y] == 1 && AliveNeigbours < 2 {
		return 0
	} else if board[x][y] == 1 && (AliveNeigbours == 3 || AliveNeigbours == 2) {
		return 1
	} else if board[x][y] == 1 && AliveNeigbours > 3 {
		return 0
	} else if AliveNeigbours == 3 {
		return 1
	}
	return 0
}

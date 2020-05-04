package main

import (
	"fmt"
)

type Board struct {
	board [YMAX][XMAX]int
}

type KingdomInfo struct {
	tileTotal [4]int
	leader    [4]int // empty (0), black (1), blue (2), green (3), red (4)
}

func (b *Board) Init(mapchoice int) {
	if mapchoice == MAPSTANDARD {
		(*b).InitMapStandard()
	} else if mapchoice == MAPADVANCE {
		(*b).InitMapAdvance()
	} else {
		(*b).InitMapTest()
	}
}

func (b *Board) InitMapStandard() {
	for j := 0; j < YMAX; j++ {
		for i := 0; i < XMAX; i++ {
			(*b).SetEmpty(i, j)
		}
	}
	for i := 4; i < 9; i++ {
		(*b).SetRiver(i, 0)
	}
	(*b).SetRiver(12, 0)

	(*b).SetRiver(4, 1)
	(*b).SetRiver(12, 1)

	(*b).SetRiver(3, 2)
	(*b).SetRiver(4, 2)
	(*b).SetRiver(12, 2)
	(*b).SetRiver(13, 2)

	for i := 0; i < 4; i++ {
		(*b).SetRiver(i, 3)
	}
	for i := 13; i < 16; i++ {
		(*b).SetRiver(i, 3)
	}

	(*b).SetRiver(14, 4)
	(*b).SetRiver(15, 4)

	(*b).SetRiver(14, 5)

	for i := 0; i < 4; i++ {
		(*b).SetRiver(i, 6)
	}
	for i := 12; i < 15; i++ {
		(*b).SetRiver(i, 6)
	}

	for i := 3; i < 7; i++ {
		(*b).SetRiver(i, 7)
	}
	(*b).SetRiver(12, 7)

	for i := 6; i < 13; i++ {
		(*b).SetRiver(i, 8)
	}

	(*b).SetTemple(10, 0)
	(*b).SetTemple(1, 1)
	(*b).SetTemple(15, 1)
	(*b).SetTemple(5, 2)
	(*b).SetTemple(13, 4)
	(*b).SetTemple(8, 6)
	(*b).SetTemple(1, 7)
	(*b).SetTemple(14, 8)
	(*b).SetTemple(5, 9)
	(*b).SetTemple(10, 10)
}

func (b *Board) InitMapAdvance() { // TODO advance map. unfinished
	for j := 0; j < YMAX; j++ {
		for i := 0; i < XMAX; i++ {
			(*b).board[j][i] = TILE["EMPTY"]
		}
	}
}

func (b *Board) InitMapTest() { // Map for test
	(*b).SetTemple(0, 0)
	(*b).SetTemple(0, 1)
	(*b).SetTemple(0, 2)
	(*b).SetTemple(1, 0)
	(*b).SetTemple(1, 2)
	(*b).SetTemple(2, 0)
	(*b).SetTemple(2, 1)
	(*b).SetTemple(2, 2)
	(*b).SetTemple(3, 0)

	(*b).SetFarm(2, 3)
	(*b).SetTemple(2, 4)
	(*b).SetMarket(1, 4)
	(*b).SetMarket(3, 4)
	(*b).SetMarket(4, 4)
	(*b).SetSettlement(5, 4)
	(*b).SetSettlement(6, 4)

	(*b).SetTile(3, 1, TILE["P1RED"])
	(*b).SetTile(6, 5, TILE["P2GREEN"])
}

func (b *Board) SetTile(x, y, thisTile int) {
	(*b).board[y][x] = thisTile
}

func (b *Board) SetEmpty(x, y int) {
	(*b).board[y][x] = TILE["EMPTY"]
}

func (b *Board) SetRiver(x, y int) {
	(*b).board[y][x] = TILE["RIVER"]
}

func (b *Board) SetFarm(x, y int) {
	(*b).board[y][x] = TILE["BLUE"]
}

func (b *Board) SetTemple(x, y int) {
	(*b).board[y][x] = TILE["RED"]
}

func (b *Board) SetMarket(x, y int) {
	(*b).board[y][x] = TILE["GREEN"]
}

func (b *Board) SetSettlement(x, y int) {
	(*b).board[y][x] = TILE["BLACK"]
}

func (b *Board) IsEmpty(x, y int) bool {
	if (*b).board[y][x] == TILE["EMPTY"] {
		return true
	} else {
		return false
	}
}

func (b *Board) IsRiver(x, y int) bool {
	if (*b).board[y][x] == TILE["RIVER"] {
		return true
	} else {
		return false
	}
}

func (b *Board) IsNeutralTile(x, y int) bool {
	switch (*b).board[y][x] {
	case TILE["BLACK"], TILE["BLUE"], TILE["GREEN"], TILE["RED"]:
		return true
	default:
		return false
	}
}

func (b *Board) IsLeader(x, y int) bool {
	switch (*b).board[y][x] {
	case TILE["P1BLACK"], TILE["P1BLUE"], TILE["P1RED"], TILE["P1GREEN"],
		TILE["P2BLACK"], TILE["P2BLUE"], TILE["P2RED"], TILE["P2GREEN"],
		TILE["P3BLACK"], TILE["P3BLUE"], TILE["P3RED"], TILE["P3GREEN"],
		TILE["P4BLACK"], TILE["P4BLUE"], TILE["P4RED"], TILE["P4GREEN"]:
		return true
	default:
		return false
	}
}

func (b *Board) PlaceTile(thisTile, x, y int) bool {
	canPlaceTile := true

	if b.IsEmpty(x, y) && !b.IsRiver(x, y) {
		(*b).board[y][x] = thisTile
	} else if b.IsEmpty(x, y) && b.IsRiver(x, y) {
		(*b).board[y][x] = thisTile
	} else { // else can't place tile
		canPlaceTile = false
	}

	return canPlaceTile
}

func (b *Board) RemoveTile(x, y int) { // TODO: remove LEADER then add to PLAYER
	switch (*b).board[y][x] {
	case TILE["P1BLUE"], TILE["P2BLUE"], TILE["P3BLUE"], TILE["P4BLUE"]: // remove farmer tile, replace with river
		(*b).board[y][x] = TILE["RIVER"]
	default: // remove normal tile, replace with empty
		(*b).board[y][x] = TILE["EMPTY"]
	}
}

func (b Board) IsLeaderPlaceable(x, y int) bool {
	for j := -1; j < 2; j++ { // assuming center is empty :-)
		for i := -1; i < 2; i++ {
			if inBound(x+i, y+j) && b.board[y+j][x+i] == TILE["RED"] {
				return true
			}
		}
	}
	return false
}

func (b Board) GetKingdomInfo(x, y int) KingdomInfo { // Get tile total using Flood Fill function
	var mark [][]bool
	var kingdomInfo KingdomInfo

	mark = make([][]bool, YMAX) // allocating memory for slice of 2D array
	for j := range mark {
		mark[j] = make([]bool, XMAX)
	}

	b.FloodFill(x, y, mark, &kingdomInfo)

	return kingdomInfo
}

func (b Board) FloodFill(x, y int, mark [][]bool, k *KingdomInfo) {
	if !inBound(x, y) { // quit function if not in bound
		return
	}

	if (b.IsNeutralTile(x, y) || b.IsLeader(x, y)) && mark[y][x] == false { // check connecting neutral & leader tile
		mark[y][x] = true

		switch b.board[y][x] {
		case TILE["BLACK"]:
			(*k).tileTotal[BLACK]++
		case TILE["BLUE"]:
			(*k).tileTotal[BLUE]++
		case TILE["GREEN"]:
			(*k).tileTotal[GREEN]++
		case TILE["RED"]:
			(*k).tileTotal[RED]++

		case TILE["P1BLACK"], TILE["P2BLACK"], TILE["P3BLACK"], TILE["P4BLACK"]:
			(*k).leader[BLACK] = b.board[y][x] / 4
		case TILE["P1BLUE"], TILE["P2BLUE"], TILE["P3BLUE"], TILE["P4BLUE"]:
			(*k).leader[BLUE] = b.board[y][x] / 4
		case TILE["P1GREEN"], TILE["P2GREEN"], TILE["P3GREEN"], TILE["P4GREEN"]:
			(*k).leader[GREEN] = b.board[y][x] / 4
		case TILE["P1RED"], TILE["P2RED"], TILE["P3RED"], TILE["P4RED"]:
			(*k).leader[RED] = b.board[y][x] / 4
		}

		b.FloodFill(x, y+1, mark, k) // up
		b.FloodFill(x+1, y, mark, k) // right
		b.FloodFill(x, y-1, mark, k) // down
		b.FloodFill(x-1, y, mark, k) // left
	}
}

func (b Board) Print() { // 16 wide x 11 height
	fmt.Printf("  ")
	for i := 0; i < XMAX; i++ {
		fmt.Printf("%2c", i+65) // print Alphabet character from unicode
	}
	fmt.Printf("\n")
	for j := 0; j < YMAX; j++ {
		fmt.Printf("%2d ", j+1)
		for i := 0; i < XMAX; i++ {
			PrintTile(b.board[j][i])
			fmt.Printf(" ")
		}
		fmt.Print("\n")
	}
}

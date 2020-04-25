package main

import (
	"fmt"
)

type Board struct {
	board [YMAX][XMAX]int
}

func (b *Board) Init(mapchoice int) {
	if mapchoice == MAPSTANDARD {
		(*b).InitMapStandard()
	} else {
		(*b).InitMapAdvance()
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

func (b *Board) SetEmpty(x, y int) {
	(*b).board[y][x] = TILE["EMPTY"]
}

func (b *Board) SetRiver(x, y int) {
	(*b).board[y][x] = TILE["RIVER"]
}

func (b *Board) SetTemple(x, y int) {
	(*b).board[y][x] = TILE["RED"]
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

func (b Board) Print() { // 16 wide x 11 height
	for j := 0; j < YMAX; j++ {
		for i := 0; i < XMAX; i++ {
			PrintTile(b.board[j][i])
			fmt.Printf(" ")
		}
		fmt.Print("\n")
	}
}
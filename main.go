package main

import (
	"fmt"
)

const XMAX = 16
const YMAX = 11
const MAXCOLOR = 4

type Player struct {
	tiles [MAXCOLOR]int
}

func (p *Player) Init() {		// randomized 5 tiles
	for i, _ :=  range p.tiles {
		(*p).tiles[i] = 1		// should be randomized from a bag
	}
}

type Tile struct {
	// 4 types of tiles. black, red, blue, green
	tile int
}

type Board struct {
	board [YMAX][XMAX]int
}

func (b *Board) Init() {
	for j:=0; j<YMAX; j++ {
		for i:=0; i<XMAX; i++ {
			(*b).board[j][i] = 0
		}
	}
}

func (b *Board) Print() {		// 16 wide x 11 height
	for j:= 0; j < YMAX; j++ {
		for i:=0; i < XMAX; i++ {
			switch b.board[j][i] {
				case 0, 1, 2: fmt.Printf("%d", (*b).board[j][i] )
				case 3, 4, 5: fmt.Printf("%d", (*b).board[j][i] )
				default: fmt.Printf("error")
			}
		}
		fmt.Print("\n")
	}
}

func main () {
        fmt.Printf("Hello Tigris & Euphrates\n")

        var p1, p2 Player
        var b Board
        b.Init()
        p1.Init()
        p2.Init()
        // b.Print()
}

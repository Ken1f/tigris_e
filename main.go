package main

import (
	"fmt"
	"math/rand"
)

const XMAX = 16
const YMAX = 11
const MAXTILES = 153
const MAXCOLOR = 4
const MAXCOLORTILES = MAXTILES / MAXCOLOR

var TILE map[string]int

func init() {
	TILE = make(map[string]int)
	TILE["EMPTY"] = 0
	TILE["BLACK"] = 1
	TILE["BLUE"]  = 2
	TILE["GREEN"] = 3
	TILE["RED"]   = 4
	TILE["RIVER"] = 5
	TILE["WAR"]   = 6
	TILE["GOLD"]  = 7
	TILE["CASTROPHE"] = 8
	// NEED RED + GOLD TILE
}

type Player struct {
	tiles [MAXCOLOR]int
}

func (p *Player) Init() { // randomized 5 tiles
	for i, _ := range p.tiles {
		(*p).tiles[i] = 1 // should be randomized from a bag
	}
}

type Tile struct {
	// 4 types of tiles. black, red, blue, green
	tile int
}

type Bag struct {
	tiles [MAXCOLOR]int
	total int
}

func (b *Bag) Init() {
	for i := 0; i < MAXCOLOR; i++ {
		b.tiles[i] = MAXCOLORTILES
	}
	b.total = MAXTILES
}

func (b *Bag) getTile() int {
	color := rand.Intn(MAXCOLOR)

	if b.tiles[color] == TILE["EMPTY"] && b.total != TILE["EMPTY"] { // if one of the colors is empty, check other color
		for color := rand.Intn(MAXCOLOR); b.tiles[color] != TILE["EMPTY"]; { // repeat until found non-empty color
		}
	}
	b.tiles[color]--
	b.total--
	return color
}

type Board struct {
	board [YMAX][XMAX]int
}

func (b *Board) Init() {
	for j := 0; j < YMAX; j++ {
		for i := 0; i < XMAX; i++ {
			(*b).board[j][i] = 0
		}
	}
}

func (b *Board) Print() { // 16 wide x 11 height
	for j := 0; j < YMAX; j++ {
		for i := 0; i < XMAX; i++ {
			switch b.board[j][i] {
			case TILE["EMPTY"]:
				fmt.Printf(" ")
			case TILE["RED"]:
				fmt.Printf("R")
			case TILE["GREEN"]:
				fmt.Printf("G")
			case TILE["BLACK"]:
				fmt.Printf("B")
			case TILE["BLUE"]:
				fmt.Printf("U")
			case TILE["RIVER"]:
				fmt.Printf("R")
			case TILE["WAR"]:
				fmt.Printf("W")
			case TILE["GOLD"]:
				fmt.Printf("O")
			case TILE["CASTROPHE"]
				fmt.Printf("C")
			default:
				fmt.Printf("error")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	fmt.Printf("Hello Tigris & Euphrates\n")

	var p1, p2 Player
	var b Board
	b.Init()
	p1.Init()
	p2.Init()
	// b.Print()
}

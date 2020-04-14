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
const EMPTY = 0

var TILE map[string]int

func init() {
	TILE = make(map[string]int)
	TILE["EMPTY"] = 0
	TILE["BLACK"] = 1
	TILE["BLUE"]  = 2
	TILE["GREEN"] = 3
	TILE["RED"]   = 4

	TILE["P1BLACK"] = 5
	TILE["P1BLUE"]  = 6
	TILE["P1GREEN"] = 7
	TILE["P1RED"]   = 8

	TILE["P2BLACK"] = 9
	TILE["P2BLUE"]  = 10
	TILE["P2GREEN"] = 11
	TILE["P2RED"]   = 12

	TILE["P3BLACK"] = 13
	TILE["P3BLUE"]  = 14
	TILE["P3GREEN"] = 15
	TILE["P3RED"]   = 16

	TILE["P4BLACK"] = 17
	TILE["P4BLUE"]  = 18
	TILE["P4GREEN"] = 19
	TILE["P4RED"]   = 20

	TILE["RIVER"] = 21
	TILE["WAR"]   = 22
	TILE["GOLD"]  = 23
	TILE["CASTROPHE"] = 24
	// NEED RED + GOLD TILE
}

type Player struct {
	tiles [5]int
	leader [4]int
}

func (p *Player) AddTile(thisTile int) {
	for i, val := range p.tiles {	// add tile to 1st empty index
		if val == EMPTY {
			(*p).tiles[i] = thisTile
			break
		}
	}
}

func (p *Player) DrawTiles(b *Bag) {
	(*p).AddTile((*b).DrawTile()) //
}

func (p *Player) Init(b *Bag, playerNum int) { // randomized 5 tiles
	for i, _ := range p.tiles {
		(*p).tiles[i] = (*b).DrawTile()// take tiles from bag to add to player tile
	}
	for i, _ := range p.leader {
		(*p).leader[i] = (playerNum-1)*4 + MAXCOLOR + 1 // example (Player1 (1) -1=0)*4=0  + MAXCOLOR (4) + 1 =  5
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

func (b *Bag) DrawTile() int {
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
			(*b).board[j][i] = EMPTY
		}
	}
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
	switch	(*b).board[y][x] {
		case	TILE["P1BLACK"], TILE["P1BLUE"], TILE["P1RED"], TILE["P1GREEN"],
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

	if b.IsEmpty(x,y) && !b.IsRiver(x,y) {
		(*b).board[y][x] = thisTile
	} else if b.IsEmpty(x,y) && b.IsRiver(x,y) {
		(*b).board[y][x] = thisTile
	} else	{				// else can't place tile
		canPlaceTile = false
	}

	return canPlaceTile
}

func (b Board) Print() { // 16 wide x 11 height
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

			case TILE["P1RED"]:
				fmt.Printf("r")
			case TILE["P1GREEN"]:
				fmt.Printf("g")
			case TILE["P1BLACK"]:
				fmt.Printf("b")
			case TILE["P1BLUE"]:
				fmt.Printf("u")

			case TILE["P2RED"]:
				fmt.Printf("r")
			case TILE["P2GREEN"]:
				fmt.Printf("g")
			case TILE["P2BLACK"]:
				fmt.Printf("b")
			case TILE["P2BLUE"]:
				fmt.Printf("u")

			case TILE["P3RED"]:
				fmt.Printf("r")
			case TILE["P3GREEN"]:
				fmt.Printf("g")
			case TILE["P3BLACK"]:
				fmt.Printf("b")
			case TILE["P3BLUE"]:
				fmt.Printf("u")

			case TILE["P4RED"]:
				fmt.Printf("r")
			case TILE["P4GREEN"]:
				fmt.Printf("g")
			case TILE["P4BLACK"]:
				fmt.Printf("b")
			case TILE["P4BLUE"]:
				fmt.Printf("u")

			case TILE["RIVER"]:
				fmt.Printf("R")
			case TILE["WAR"]:
				fmt.Printf("W")
			case TILE["GOLD"]:
				fmt.Printf("O")
			case TILE["CASTROPHE"]:
				fmt.Printf("C")
			default:
				fmt.Printf("error")
			}
		}
		fmt.Print("\n")
	}
}

func p(s string) {
	fmt.Printf(s + "\n")
}

func main() {
	fmt.Printf("Hello Tigris & Euphrates\n")

	var p1, p2 Player
	var board Board
	var bag Bag
	p("board init")
	board.Init()
	p("bag init")
	bag.Init()
	p("player init")
	p1.Init(&bag, 1)
	p2.Init(&bag, 2)
	// b.Print()
}

package main

import (
	"fmt"
	"math/rand"
)

const XMAX = 16
const YMAX = 11
const MAXBLACK = 30
const MAXBLUE = 36
const MAXGREEN = 30
const MAXRED = 57
const MAXTILES = MAXBLACK + MAXBLUE + MAXGREEN + MAXRED
const MAXCOLOR = 4

const (
	BLACK = iota
	BLUE
	GREEN
	RED
)

const (
	COLORBLACK   = "\033[1;30m%s\033[0m"
	COLORRED     = "\033[1;31m%s\033[0m"
	COLORGREEN   = "\033[1;32m%s\033[0m"
	COLORYELLOW  = "\033[1;33m%s\033[0m"
	COLORBLUE    = "\033[1;34m%s\033[0m"
	COLORMAGENTA = "\033[1;35m%s\033[0m"
	COLORCYAN    = "\033[1;36m%s\033[0m"
	COLORWHITE   = "\033[1;37m%s\033[0m"
)

const (
	MAPSTANDARD = iota
	MAPADVANCE
)

const (
	PLAYER1 = iota
	PLAYER2
	PLAYER3
	PLAYER4
)

var TILE map[string]int

func init() {
	TILE = make(map[string]int)

	TILE["BLACK"] = 0
	TILE["BLUE"] = 1
	TILE["GREEN"] = 2
	TILE["RED"] = 3

	TILE["P1BLACK"] = 4
	TILE["P1BLUE"] = 5
	TILE["P1GREEN"] = 6
	TILE["P1RED"] = 7

	TILE["P2BLACK"] = 8
	TILE["P2BLUE"] = 9
	TILE["P2GREEN"] = 10
	TILE["P2RED"] = 11

	TILE["P3BLACK"] = 12
	TILE["P3BLUE"] = 13
	TILE["P3GREEN"] = 14
	TILE["P3RED"] = 15

	TILE["P4BLACK"] = 16
	TILE["P4BLUE"] = 17
	TILE["P4GREEN"] = 18
	TILE["P4RED"] = 19

	TILE["EMPTY"] = 20
	TILE["RIVER"] = 21
	TILE["WAR"] = 22
	TILE["GOLD"] = 23
	TILE["CASTROPHE"] = 24

	// NEED RED + GOLD TILE
}

type Player struct {
	playerNum int
	tiles     [5]int
	leader    [4]int
	points    [4]int
}

func (p *Player) AddTile(thisTile int) {
	for i, val := range p.tiles { // add tile to 1st empty index
		if val == TILE["EMPTY"] {
			(*p).tiles[i] = thisTile
			break
		}
	}
}

func (p *Player) DrawTiles(b *Bag) {
	(*p).AddTile((*b).DrawTile()) //
}

func (p *Player) Init(b *Bag, playerNum int) { // randomized 5 tiles
	(*p).playerNum = playerNum
	for i, _ := range p.tiles {
		(*p).tiles[i] = (*b).DrawTile() // take tiles from bag to add to player tile
	}
	for i, _ := range p.leader {
		(*p).leader[i] = playerNum*4 + MAXCOLOR // example PLAYER1 (1) *4=4  + MAXCOLOR (4) =  4
	}
}

func (p *Player) AddPoint(thisTile int) {
	(*p).points[(thisTile-1)%MAXCOLOR]++
}

func (p Player) Print() {
	fmt.Printf("\033[1;30m%d\033[0m", p.leader[0]) // black
	fmt.Printf("\033[1;34m%d\033[0m", p.leader[1]) // blue
	fmt.Printf("\033[1;32m%d\033[0m", p.leader[2]) // green
	fmt.Printf("\033[1;31m%d\033[0m", p.leader[3]) // red

	fmt.Printf(" ")

	fmt.Printf("\033[1;30m%d\033[0m", p.points[0]) // black
	fmt.Printf("\033[1;34m%d\033[0m", p.points[1]) // blue
	fmt.Printf("\033[1;32m%d\033[0m", p.points[2]) // green
	fmt.Printf("\033[1;31m%d\033[0m", p.points[3]) // red

	fmt.Printf(" ")
	for i := 0; i < 5; i++ {
		PrintTile(p.tiles[i])
	}
	fmt.Printf("\n")
}

type Bag struct {
	tiles [MAXCOLOR]int
	total int
}

func (b *Bag) Init() {
	b.tiles[BLACK] = MAXBLACK
	b.tiles[BLUE] = MAXBLUE
	b.tiles[GREEN] = MAXGREEN
	b.tiles[RED] = MAXRED

	b.total = MAXTILES
}

func (b *Bag) DrawTile() int {
	color := rand.Intn(MAXCOLOR + 1)
	if color > RED { // red can represent 3 or 4 since it has 2x the amount of black, blue, or green
		color = RED
	}

	if b.tiles[color] == TILE["EMPTY"] && b.total != TILE["EMPTY"] { // if one of the colors is empty, check other color
		for color := rand.Intn(MAXCOLOR); b.tiles[color] != TILE["EMPTY"]; { // repeat until found non-empty color
		}
	}
	b.tiles[color]--
	b.total--
	return color
}

func (b Bag) RemainingTile() int {
	return b.total
}

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

func inBound(x, y int) bool {
	if x >= 0 && x < XMAX && y >= 0 && y < YMAX {
		return true
	} else {
		return false
	}
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

func p(s string) {
	fmt.Printf(s + "\n")
}

func PrintTile(thisTile int) {
	switch thisTile {
	case TILE["EMPTY"]:
		fmt.Printf(COLORYELLOW, "A")
	case TILE["RED"]:
		fmt.Printf(COLORRED, "T")
	case TILE["GREEN"]:
		fmt.Printf(COLORGREEN, "M")
	case TILE["BLACK"]:
		fmt.Printf(COLORBLACK, "S")
	case TILE["BLUE"]:
		fmt.Printf(COLORBLUE, "F")

	case TILE["P1RED"]:
		fmt.Printf(COLORRED, "1")
	case TILE["P1GREEN"]:
		fmt.Printf(COLORGREEN, "1")
	case TILE["P1BLACK"]:
		fmt.Printf(COLORBLACK, "1")
	case TILE["P1BLUE"]:
		fmt.Printf(COLORMAGENTA, "1")

	case TILE["P2RED"]:
		fmt.Printf(COLORRED, "2")
	case TILE["P2GREEN"]:
		fmt.Printf(COLORGREEN, "2")
	case TILE["P2BLACK"]:
		fmt.Printf(COLORBLACK, "2")
	case TILE["P2BLUE"]:
		fmt.Printf(COLORMAGENTA, "2")

	case TILE["P3RED"]:
		fmt.Printf(COLORRED, "3")
	case TILE["P3GREEN"]:
		fmt.Printf(COLORGREEN, "3")
	case TILE["P3BLACK"]:
		fmt.Printf(COLORBLACK, "3")
	case TILE["P3BLUE"]:
		fmt.Printf(COLORMAGENTA, "3")

	case TILE["P4RED"]:
		fmt.Printf(COLORRED, "4")
	case TILE["P4GREEN"]:
		fmt.Printf(COLORGREEN, "4")
	case TILE["P4BLACK"]:
		fmt.Printf(COLORBLACK, "4")
	case TILE["P4BLUE"]:
		fmt.Printf(COLORMAGENTA, "4")

	case TILE["RIVER"]:
		fmt.Printf(COLORBLUE, "R")
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

func main() {
	fmt.Printf("Hello Tigris & Euphrates\n")

	var p1, p2, p3, p4 Player
	var board Board
	var bag Bag
	p("board init")
	board.Init(MAPSTANDARD)
	p("bag init")
	bag.Init()
	p("player init")
	p1.Init(&bag, PLAYER1)
	p1.Print()
	p2.Init(&bag, PLAYER2)
	p2.Print()
	p3.Init(&bag, PLAYER3)
	p3.Print()
	p4.Init(&bag, PLAYER4)
	p4.Print()
	board.Print()
}

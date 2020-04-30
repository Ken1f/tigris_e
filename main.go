package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
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
	MAPTEST
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

func inBound(x, y int) bool {
	if x >= 0 && x < XMAX && y >= 0 && y < YMAX {
		return true
	} else {
		return false
	}
}

func pr(s string) {
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
		fmt.Printf(COLORBLUE, "1")

	case TILE["P2RED"]:
		fmt.Printf(COLORRED, "2")
	case TILE["P2GREEN"]:
		fmt.Printf(COLORGREEN, "2")
	case TILE["P2BLACK"]:
		fmt.Printf(COLORBLACK, "2")
	case TILE["P2BLUE"]:
		fmt.Printf(COLORBLUE, "2")

	case TILE["P3RED"]:
		fmt.Printf(COLORRED, "3")
	case TILE["P3GREEN"]:
		fmt.Printf(COLORGREEN, "3")
	case TILE["P3BLACK"]:
		fmt.Printf(COLORBLACK, "3")
	case TILE["P3BLUE"]:
		fmt.Printf(COLORBLUE, "3")

	case TILE["P4RED"]:
		fmt.Printf(COLORRED, "4")
	case TILE["P4GREEN"]:
		fmt.Printf(COLORGREEN, "4")
	case TILE["P4BLACK"]:
		fmt.Printf(COLORBLACK, "4")
	case TILE["P4BLUE"]:
		fmt.Printf(COLORBLUE, "4")

	case TILE["RIVER"]:
		fmt.Printf(COLORBLUE, "R")
	case TILE["WAR"]:
		fmt.Printf(COLORMAGENTA, "W")
	case TILE["GOLD"]:
		fmt.Printf("O")
	case TILE["CASTROPHE"]:
		fmt.Printf(COLORWHITE, "C")
	default:
		fmt.Printf("error")
	}
}

func printGame(p []Player, board Board, bag Bag) {	// NOTE: P[] is a slice. Not pointer
	for i, _ := range p {
		p[i].Print()
	}
	board.Print()
	fmt.Print("Remaining Tile ", bag.RemainingTile(), "\n")
}

func printKingdomInfo(kingdomInfo[] int) {
	fmt.Print("Kingdom info:")
	fmt.Printf(COLORBLACK, " BLACK:")
	fmt.Print(kingdomInfo[BLACK])
	fmt.Printf(COLORBLUE, " BLUE:")
	fmt.Print(kingdomInfo[BLUE])
	fmt.Printf(COLORGREEN, " GREEN:")
	fmt.Print(kingdomInfo[GREEN])
	fmt.Printf(COLORRED, " RED:")
	fmt.Print(kingdomInfo[RED])
}

func readInput() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() 
	line := scanner.Text()
	for _, x := range strings.Fields(line) {
		fmt.Println("next field", x)
	}
}

func main() {
	fmt.Printf("Hello Tigris & Euphrates\n")

	var p[4] Player
	var board Board
	var bag Bag
	pr("board init")
	board.Init(MAPSTANDARD)
	pr("bag init")
	bag.Init()
	fmt.Print("Starting Tile ", bag.RemainingTile(), "\n")
	pr("player init")
	p[0].Init(PLAYER1, bag.DrawTiles(5))
	p[1].Init(PLAYER1, bag.DrawTiles(5))
	p[2].Init(PLAYER1, bag.DrawTiles(5))
	p[3].Init(PLAYER1, bag.DrawTiles(5))
	printGame(p[:], board, bag)

	swap := []int {0,1,2} //
	p[0].SwapTiles(bag.DrawTiles(len(swap)), swap)
	swap = []int {0,1,2,3,4} //
	p[1].SwapTiles(bag.DrawTiles(len(swap)), swap)
	//readInput()
	board.Init(MAPTEST)
	printGame(p[:], board, bag)

	kingdomInfo := board.GetKingdomInfo(1,1)
	printKingdomInfo(kingdomInfo)
}
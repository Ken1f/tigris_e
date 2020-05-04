package main

import (
	"math/rand"
)

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

func (b *Bag) DrawTile() int { // draw 1 tile
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

func (b *Bag) DrawTiles(numTile int) []int { // draw multiple tiles
	var tiles []int

	for i := 0; i < numTile; i++ {
		tiles = append(tiles, (*b).DrawTile())
	}
	return tiles
}

func (b Bag) RemainingTile() int {
	return b.total
}

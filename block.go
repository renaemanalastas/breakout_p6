package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	BlockTile rl.Rectangle
	Position rl.Vector2
	Color rl.Color
}

func newBlock(pos rl.Vector2, color rl.Color) Block {
	return Block{
		BlockTile:	rl.NewRectangle(pos.X, pos.Y, TileSize, TileSize),
		Position:	pos,
		Color:		color,
	}
}

func initGrid() []Block{
	blockGrid := make([]Block, 0)
	for i := TileSize; i < 720; i+=TileSize {
		for j := TileSize; j < 240; j+=TileSize+4{
			//x, y, width, height, color
			blockGrid = append(blockGrid, newBlock(rl.NewVector2(float32(i), float32(j)), rl.Green))
		}
		i += 4
	}
	return blockGrid
}

func drawGrid(grid []Block){
	for _, block := range grid {
		rl.DrawRectangleRec(block.BlockTile, block.Color)
	}
}
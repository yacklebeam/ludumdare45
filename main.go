package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

const (
	gWindowWidth  = 800
	gWindowHeight = 600
)

func loadLevel(levelID uint16) {
	// setup game -- these are globally accessible
	gPositionCompMap[0] = positionComp{x: 100, y: 100, width: 20, height: 20}
	gRenderCompMap[0] = renderComp{color: rl.Red}
	gPhysicsCompMap[0] = physicsComp{x: 20, y: 20}
	gCollisionCompMap[0] = collisionComp{xOffset: -5, yOffset: -5, width: 30, height: 30}
	gTextCompMap[0] = textComponent{text: "Sample", color: rl.White, size: 20, xOffset: 30, yOffset: 0}

	gPositionCompMap[56] = positionComp{x: 300, y: 40, width: 20, height: 50}
	gRenderCompMap[56] = renderComp{color: rl.Blue}
	gCollisionCompMap[56] = collisionComp{xOffset: 5, yOffset: -25, width: 10, height: 100}

	gPositionCompMap[404] = positionComp{x: 30, y: 400, width: 50, height: 20}
	gRenderCompMap[404] = renderComp{color: rl.Green}

	gPositionCompMap[76] = positionComp{x: 10, y: 10, width: 50, height: 20}
	gTextCompMap[76] = textComponent{text: "Example text", color: rl.White, size: 20, xOffset: 0, yOffset: 0}
}

func drawTileMap(tilemap *tiled.Map, texture rl.Texture2D) {
	tileMapImage := texture

	tileW := int32(tilemap.TileWidth)
	tileH := int32(tilemap.TileHeight)

	for _, layer := range tilemap.Layers {
		for tileIndex, tile := range layer.Tiles {
			id := int32(tile.ID)

			tileDestX := int32(tileIndex%tilemap.Width) * tileW
			tileDestY := int32(int32(tileIndex)/int32(tilemap.Width)) * tileH

			var tileSrcX int32 = id % (tileMapImage.Width / tileW)
			var tileSrcY int32 = int32(id / (tileMapImage.Width / tileW))

			rl.DrawTexturePro(tileMapImage,
				rl.NewRectangle(float32(tileSrcX*tileW), float32(tileSrcY*tileH), float32(tileW), float32(tileH)),
				rl.NewRectangle(float32(tileDestX), float32(tileDestY), float32(tileW), float32(tileH)), rl.NewVector2(0, 0), 0, rl.White)

		}
	}
}

func main() {

	rl.InitWindow(gWindowWidth, gWindowHeight, "Ludum Dare #45 Game")

	loadLevel(0)

	img := rl.LoadImage("assets/testtileset.png")
	tex := rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	gameMap, _ := tiled.LoadFromFile("assets/testmap.tmx")

	for !rl.WindowShouldClose() {
		t := rl.GetFrameTime()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// ECS System Updates -- organized by impact
		physicsSystem(gPositionCompMap, gPhysicsCompMap, t)

		drawTileMap(gameMap, tex)

		// Rendering
		renderSystem(gPositionCompMap, gRenderCompMap)
		collisionRenderSystem(gPositionCompMap, gCollisionCompMap)
		textRenderSystem(gPositionCompMap, gTextCompMap)

		rl.EndDrawing()
	}
	rl.CloseWindow()
}

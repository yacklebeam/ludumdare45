package system

import rl "github.com/gen2brain/raylib-go/raylib"

// system code is NOT ECS systems
// system is for interfacing with assets, etc
// - texture loading
// - audio loading
// - game data loading

var textureMap map[string]rl.Texture2D
var audioMap map[string]rl.Sound

func init() {
	textureMap = make(map[string]rl.Texture2D)
	audioMap = make(map[string]rl.Sound)
}

func loadTextureFromFile(filename string) {
	// always loads from assets/images
	// this function MUST be run after rl.InitWindow() in the main game
	img := rl.LoadImage("assets/images/" + filename)
	texture := rl.LoadTextureFromImage(img)
	textureMap[filename] = texture
}

func loadAudioFromFile(filename string) {
	// always loads from assets/audio
	// this function MUST be run after rl.InitWindow() in the main game
	sound := rl.LoadSound("assets/audio/" + filename)
	audioMap[filename] = sound
}

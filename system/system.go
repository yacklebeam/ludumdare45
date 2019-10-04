package system

import rl "github.com/gen2brain/raylib-go/raylib"

// system code is NOT ECS systems
// system is for interfacing with assets, etc
// - texture loading
// - audio loading
// - game data loading

var TextureMap map[string]rl.Texture2D
var AudioMap map[string]rl.Sound

func init() {
	TextureMap = make(map[string]rl.Texture2D)
	AudioMap = make(map[string]rl.Sound)
}

func LoadTextureFromFile(filename string) {
	// always loads from assets/images
	// this function MUST be run after rl.InitWindow() in the main game
	img := rl.LoadImage("assets/images/" + filename)
	texture := rl.LoadTextureFromImage(img)
	TextureMap[filename] = texture
}

func LoadAudioFromFile(filename string) {
	// always loads from assets/audio
	// this function MUST be run after rl.InitWindow() in the main game
	sound := rl.LoadSound("assets/audio/" + filename)
	AudioMap[filename] = sound
}

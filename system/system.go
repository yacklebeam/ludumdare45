package system

import rl "github.com/gen2brain/raylib-go/raylib"

// system code is NOT ECS systems
// system is for interfacing with assets, etc
// - texture loading
// - audio loading
// - game data loading

var textureMap map[string]rl.Texture2D
var soundMap map[string]rl.Sound
var musicMap map[string]rl.Music

func init() {
	textureMap = make(map[string]rl.Texture2D)
	soundMap = make(map[string]rl.Sound)
	musicMap = make(map[string]rl.Music)
}

func LoadDefaults() {
	// run this AFTER rl.InitWindow()
	LoadTextureFromFile("missing_texture.png")
}

func LoadTextureFromFile(filename string) {
	// always loads from assets/images
	// this function MUST be run after rl.InitWindow() in the main game
	img := rl.LoadImage("assets/images/" + filename)
	texture := rl.LoadTextureFromImage(img)
	textureMap[filename] = texture
}

func LoadSoundFromFile(filename string) {
	// always loads from assets/audio
	// this function MUST be run after rl.InitWindow() in the main game
	sound := rl.LoadSound("assets/audio/" + filename)
	soundMap[filename] = sound
}

func LoadMusicFromFile(filename string) {
	// always loads from assets/audio
	// this function MUST be run after rl.InitWindow() in the main game
	music := rl.LoadMusicStream("assets/audio/" + filename)
	musicMap[filename] = music
}

func GetTexture(filename string) rl.Texture2D {
	t, exists := textureMap[filename]
	if exists {
		return t
	} else {
		return textureMap["missing_texture.png"]
	}
}

func GetSound(filename string) rl.Sound {
	t, exists := soundMap[filename]
	if exists {
		return t
	} else {
		return rl.NewSound(0, 0, 0)
	}
}

func GetMusic(filename string) rl.Music {
	t, exists := musicMap[filename]
	if exists {
		return t
	} else {
		return musicMap["default.ogg"]
	}
}

package main

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
	UIHeight     int
}

func NewGameData() GameData {
	g := GameData{
		ScreenWidth:  80,
		ScreenHeight: 60,
		TileWidth:    16,
		TileHeight:   16,
		UIHeight:     10,
	}
	return g
}

package main

import (
	"github.com/bytearena/ecs"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Map         GameMap
	World       *ecs.Manager
	WorldTags   map[string]ecs.Tag
	Turn        TurnState
	TurnCounter int
}

// NewGame creates a new Game Object and initializes the data
// This is a pretty solid refactor candidate for later
func NewGame() *Game {
	g := &Game{}
	g.Map = NewGameMap()
	world, tags := InitializeWorld(g.Map.CurrentLevel)
	g.WorldTags = tags
	g.World = world
	g.Turn = PlayerTurn
	g.TurnCounter = 0
	return g
}

// Update is called each tic.
func (g *Game) Update() error {
	g.TurnCounter++

	if g.Turn == PlayerTurn && g.TurnCounter > 20 {
		MovePlayer(g)
	}

	if g.Turn == MonsterTurn {
		UpdateMonster(g)
	}

	return nil
}

// Draw is called each draw cycle and is where we will blit.
func (g *Game) Draw(screen *ebiten.Image) {
	//Draw the Map
	level := g.Map.CurrentLevel
	level.DrawLevel(screen)
	ProcessRenderables(g, level, screen)
	ProcessUserLog(g, screen)
	ProcessHUD(g, screen)
}

// Layout will return the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) {
	gd := NewGameData()
	return gd.TileWidth * gd.ScreenWidth, gd.TileHeight * gd.ScreenHeight
}

func main() {
	g := NewGame()
	ebiten.SetWindowTitle("Tower")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

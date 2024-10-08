package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct{}
type Monster struct{}

type Name struct {
	Label string
}

type Position struct {
	X int
	Y int
}

type Health struct {
	MaxHealth     int
	CurrentHealth int
}

type UserMessage struct {
	AttackMessage    string
	DeadMessage      string
	GameStateMessage string
}

type MeleeWeapon struct {
	Name          string
	MinimumDamage int
	MaximumDamage int
	ToHitBonus    int
}

type Armor struct {
	Name       string
	Defense    int
	ArmorClass int
}

type Renderable struct {
	Image *ebiten.Image
}

type Movable struct{}

func (p *Position) GetManhattanDistance(other *Position) int {
	xDist := math.Abs(float64(p.X - other.X))
	yDist := math.Abs(float64(p.Y - other.Y))
	return int(xDist) + int(yDist)
}

func (p *Position) IsEqual(other *Position) bool {
	return (p.X == other.X && p.Y == other.Y)
}

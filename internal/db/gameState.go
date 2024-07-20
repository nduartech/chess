package db

import (
	"time"
)

func NewGameState(local bool, playerSide bool, botDifficulty int, started time.Time, pgn string, ended bool) *GameState {
	return &GameState{local, playerSide, botDifficulty, started, pgn, ended}
}

func (g *GameState) GetLocal() bool {
	return g.local
}

func (g *GameState) GetPlayerSide() bool {
	return g.playerSide
}

func (g *GameState) GetBotDifficulty() int {
	return g.botDifficulty
}

func (g *GameState) GetStarted() time.Time {
	return g.started
}

func (g *GameState) GetPgn() string {
	return g.pgn
}

func (g *GameState) GetEnded() bool {
	return g.ended
}

func (g *GameState) SetLocal(local bool) {
	g.local = local
}

func (g *GameState) SetPlayerSide(playerSide bool) {
	g.playerSide = playerSide
}

func (g *GameState) SetBotDifficulty(botDifficulty int) {
	g.botDifficulty = botDifficulty
}

func (g *GameState) SetStarted(started time.Time) {
	g.started = started
}

func (g *GameState) SetPgn(pgn string) {
	g.pgn = pgn
}

func (g *GameState) SetEnded(ended bool) {
	g.ended = ended
}

package db

import (
	"context"
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

type DBConn struct {
	db *sql.DB
}

type GameState struct {
	local         bool
	playerSide    bool
	botDifficulty int
	started       time.Time
	pgn           string
	ended         bool
}

type GameStateRow struct {
	ID int
	GameState
}

func NewDBConn(db *sql.DB) *DBConn {
	return &DBConn{db}
}

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

func InitDatabase(dbPath string) (*DBConn, error) {
	var err error
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS games (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
		  local BOOLEAN NOT NULL,
		  player_side BOOLEAN NOT NULL, 
		  bot_difficulty INTEGER,
			started DATETIME NOT NULL,
			pgn TEXT NOT NULL, 
			ended BOOLEAN NOT NULL
		)`,
	)
	if err != nil {
		return nil, err
	}
	return &DBConn{db}, nil
}

func (d *DBConn) AddGame(g *GameState) (int64, error) {
	res, err := d.db.Exec(`INSERT INTO games (local, player_side, bot_difficulty, started, pgn, ended) VALUES (?,?,?,?,?,?);`, g.local, g.playerSide, g.botDifficulty, g.started, g.pgn, g.ended)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (d *DBConn) UpdateGame(id int64, g *GameState) error {
	_, err := d.db.Exec(`UPDATE games SET local=?, player_side=?, bot_difficulty=?, started=?, pgn=?, ended=? WHERE id=?`, g.local, g.playerSide, g.botDifficulty, g.started, g.pgn, g.ended, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DBConn) GetRecentGame() (*GameState, error) {
	var state GameStateRow
	_, err := d.db.QueryContext(context.Background(), "SELECT COUNT(*) FROM games")
	if err != nil {
		return nil, err
	}
	row := d.db.QueryRow("SELECT * FROM games ORDER BY started DESC")
	err = row.Scan(&state.ID, &state.local, &state.playerSide, &state.botDifficulty, &state.started, &state.pgn, &state.ended)
	if err != nil {
		return nil, err
	}
	return &state.GameState, nil
}

func (d *DBConn) ClearGameHistory() error {
	_, err := d.db.Exec("DELETE * FROM games")
	if err != nil {
		return err
	}
	return nil
}

func (d *DBConn) Close() {
	d.db.Close()
}

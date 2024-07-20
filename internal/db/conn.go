package db

import (
	"context"
	"database/sql"
	_ "modernc.org/sqlite"
	"time"
)

type Conn interface {
	Open() error
	AddGame(g *GameState) (int64, error)
	UpdateGame(id int64, g *GameState) error
	GetRecentGame() (*GameStateRow, error)
	ClearGameHistory() error
	Close()
}

type ConnSQLite struct {
	path string
	db   *sql.DB
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
	ID int64
	GameState
}

func NewDBConn(db *sql.DB, path string) *ConnSQLite {
	return &ConnSQLite{path, db}
}

func InitGameDatabase(dbPath string) (*ConnSQLite, error) {
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
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	return NewDBConn(db, dbPath), nil
}

func (d *ConnSQLite) Open() error {
	conn, err := sql.Open("sqlite", d.path)
	if err != nil {
		return err
	}
	d.db = conn
	return nil
}

func (d *ConnSQLite) AddGame(g *GameState) (int64, error) {
	_ = d.Open()
	defer d.Close()
	res, err := d.db.Exec(`INSERT INTO games (local, player_side, bot_difficulty, started, pgn, ended) VALUES (?,?,?,?,?,?);`, g.local, g.playerSide, g.botDifficulty, g.started, g.pgn, g.ended)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (d *ConnSQLite) UpdateGame(id int64, g *GameState) error {
	_ = d.Open()
	defer d.Close()
	_, err := d.db.Exec(`UPDATE games SET local=?, player_side=?, bot_difficulty=?, started=?, pgn=?, ended=? WHERE id=?`, g.local, g.playerSide, g.botDifficulty, g.started, g.pgn, g.ended, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *ConnSQLite) GetRecentGame() (*GameStateRow, error) {
	_ = d.Open()
	defer d.Close()
	var state GameStateRow
	rows, err := d.db.QueryContext(context.Background(), "SELECT COUNT(*) FROM games")
	if err != nil {
		return nil, err
	} else {
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				panic(err)
			}
		}(rows)
	}
	row := d.db.QueryRow("SELECT * FROM games ORDER BY started DESC")
	err = row.Scan(&state.ID, &state.local, &state.playerSide, &state.botDifficulty, &state.started, &state.pgn, &state.ended)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func (d *ConnSQLite) ClearGameHistory() error {
	_ = d.Open()
	defer d.Close()
	_, err := d.db.Exec("DELETE FROM games")
	if err != nil {
		return err
	}
	return nil
}

func (d *ConnSQLite) Close() {
	err := d.db.Close()
	if err != nil {
		panic(err)
	}
}

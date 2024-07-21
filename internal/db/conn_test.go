package db

import (
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testDBPath = "test.db"

func setupTestDB(t *testing.T) *ConnSQLite {
	conn, err := InitGameDatabase(testDBPath)
	require.NoError(t, err)
	return conn
}

func cleanupTestDB(t *testing.T) {
	err := os.Remove(testDBPath)
	require.NoError(t, err)
}

func TestNewDBConn(t *testing.T) {
	db, err := sql.Open("sqlite", testDBPath)
	require.NoError(t, err)
	defer db.Close()

	conn := NewDBConn(db, testDBPath)
	assert.NotNil(t, conn)
	assert.Equal(t, testDBPath, conn.path)
	assert.Equal(t, db, conn.db)
}

func TestInitGameDatabase(t *testing.T) {
	conn, err := InitGameDatabase(testDBPath)
	require.NoError(t, err)
	assert.NotNil(t, conn)
	cleanupTestDB(t)
}

func TestConnSQLite_Open(t *testing.T) {
	conn := setupTestDB(t)
	defer cleanupTestDB(t)

	err := conn.Open()
	require.NoError(t, err)
	assert.NotNil(t, conn.db)
}

func TestConnSQLite_AddGame(t *testing.T) {
	conn := setupTestDB(t)
	defer cleanupTestDB(t)

	game := &GameState{
		Local:         true,
		PlayerSide:    true,
		Turn:          false,
		BotDifficulty: 3,
		Started:       time.Now(),
		Pgn:           "e4 e5",
		Ended:         false,
	}

	id, err := conn.AddGame(game)
	require.NoError(t, err)
	assert.Greater(t, id, int64(0))
}

func TestConnSQLite_UpdateGame(t *testing.T) {
	conn := setupTestDB(t)
	defer cleanupTestDB(t)

	game := &GameState{
		Local:         true,
		PlayerSide:    true,
		Turn:          false,
		BotDifficulty: 3,
		Started:       time.Now(),
		Pgn:           "e4 e5",
		Ended:         false,
	}

	id, err := conn.AddGame(game)
	require.NoError(t, err)

	updatedGame := &GameStateRow{
		ID:        id,
		GameState: *game,
	}
	updatedGame.Pgn = "e4 e5 Nf3"
	updatedGame.Ended = true

	err = conn.UpdateGame(updatedGame)
	require.NoError(t, err)

	retrievedGame, err := conn.GetRecentGame()
	require.NoError(t, err)
	assert.Equal(t, updatedGame.Pgn, retrievedGame.Pgn)
	assert.Equal(t, updatedGame.Ended, retrievedGame.Ended)
}

func TestConnSQLite_GetRecentGame(t *testing.T) {
	conn := setupTestDB(t)
	defer cleanupTestDB(t)

	game1 := &GameState{
		Local:         true,
		PlayerSide:    true,
		Turn:          false,
		BotDifficulty: 3,
		Started:       time.Now().Add(-1 * time.Hour),
		Pgn:           "e4 e5",
		Ended:         false,
	}

	game2 := &GameState{
		Local:         false,
		PlayerSide:    false,
		Turn:          true,
		BotDifficulty: 5,
		Started:       time.Now(),
		Pgn:           "d4 d5",
		Ended:         true,
	}

	_, err := conn.AddGame(game1)
	require.NoError(t, err)

	_, err = conn.AddGame(game2)
	require.NoError(t, err)

	recentGame, err := conn.GetRecentGame()
	require.NoError(t, err)
	assert.Equal(t, game2.Pgn, recentGame.Pgn)
	assert.Equal(t, game2.Started.Unix(), recentGame.Started.Unix())
}

func TestConnSQLite_ClearGameHistory(t *testing.T) {
	conn := setupTestDB(t)
	defer cleanupTestDB(t)

	game := &GameState{
		Local:         true,
		PlayerSide:    true,
		Turn:          false,
		BotDifficulty: 3,
		Started:       time.Now(),
		Pgn:           "e4 e5",
		Ended:         false,
	}

	_, err := conn.AddGame(game)
	require.NoError(t, err)

	err = conn.ClearGameHistory()
	require.NoError(t, err)

	recentGame, err := conn.GetRecentGame()
	assert.Error(t, err)
	assert.Nil(t, recentGame)
}

func TestConnSQLite_Close(t *testing.T) {
	conn := setupTestDB(t)
	defer cleanupTestDB(t)

	err := conn.Open()
	require.NoError(t, err)
	assert.NotNil(t, conn.db)

	conn.Close()
}

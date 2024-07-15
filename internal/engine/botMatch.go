package engine

import (
	"db"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
	"github.com/olahol/melody"
)

func ChessMatchWithBot(difficulty int, playerWhite bool, m *melody.Melody, g *db.GameStateRow, d *db.DBConn) {
	chessBot := NewChessAI(difficulty)
	var game *chess.Game
	if g == nil {
		game = chess.NewGame()
		gs := db.NewGameState(true, playerWhite, difficulty, time.Now(), game.String(), false)
		index, err := d.AddGame(gs)
		if err != nil {
			panic(err)
		}
		g = &db.GameStateRow{index, *gs}
	} else {
		pgn, err := chess.PGN(strings.NewReader(g.GetPgn()))
		if err != nil {
			panic(err)
		}
		game = chess.NewGame(pgn)
	}
	m.HandleConnect(func(session *melody.Session) {

		err := session.Write([]byte("Init:" + fmt.Sprint(game.Position().Board().SquareMap())))
		if err != nil {
			err2 := m.Close()
			if err2 != nil {
				panic(err2)
			}
			panic(err)
		}
		if playerWhite {
			err := session.Write([]byte("Valid Moves:" + fmt.Sprint(game.ValidMoves())))
			if err != nil {
				err2 := m.Close()
				if err2 != nil {
					panic(err2)
				}
				panic(err)
			}
		} else {
			move := chessBot.MakeMove(game)

			err := session.Write([]byte(move.String()))
			if err != nil {
				err2 := m.Close()
				if err2 != nil {
					panic(err2)
				}
				panic(err)
			}
			err = session.Write([]byte("Valid Moves:" + fmt.Sprint(game.ValidMoves())))
			if err != nil {
				err2 := m.Close()
				if err2 != nil {
					panic(err2)
				}
				panic(err)
			}
		}
	})

	m.HandleDisconnect(func(session *melody.Session) {
		d.UpdateGame(g.ID,
			db.NewGameState(
				g.GetLocal(),
				g.GetPlayerSide(),
				g.GetBotDifficulty(),
				g.GetStarted(),
				game.String(),
				(game.Outcome() != chess.NoOutcome),
			),
		)
	})

	m.HandleMessage(func(session *melody.Session, bytes []byte) {

		err := game.MoveStr(string(bytes))
		if err != nil {
			err2 := m.Close()
			if err2 != nil {
				panic(err2)
			}
			panic(err)
		}

		d.UpdateGame(g.ID,
			db.NewGameState(
				g.GetLocal(),
				g.GetPlayerSide(),
				g.GetBotDifficulty(),
				g.GetStarted(),
				game.String(),
				(game.Outcome() != chess.NoOutcome),
			),
		)

		if game.Outcome() == chess.NoOutcome {
			move := chessBot.MakeMove(game)
			err := session.Write([]byte(move.String()))
			if err != nil {
				err2 := m.Close()
				if err2 != nil {
					panic(err2)
				}
				panic(err)
			}
			d.UpdateGame(g.ID,
				db.NewGameState(
					g.GetLocal(),
					g.GetPlayerSide(),
					g.GetBotDifficulty(),
					g.GetStarted(),
					game.String(),
					(game.Outcome() != chess.NoOutcome),
				),
			)

			if game.Outcome() == chess.NoOutcome {
				err = session.Write([]byte("Valid Moves:" + fmt.Sprint(game.ValidMoves())))
				if err != nil {
					err2 := m.Close()
					if err2 != nil {
						panic(err2)
					}
					panic(err)
				}
			} else {
				fmt.Printf("Game completed. %s by %s.\n", game.Outcome(), game.Method())
				err := session.Write([]byte(fmt.Sprintf("Game completed. %s by %s.\n", game.Outcome(), game.Method())))
				if err != nil {
					err2 := m.Close()
					if err2 != nil {
						panic(err2)
					}
					panic(err)
				}

			}
		} else {
			fmt.Printf("Game completed. %s by %s.\n", game.Outcome(), game.Method())
			err := session.Write([]byte(fmt.Sprintf("Game completed. %s by %s.\n", game.Outcome(), game.Method())))
			if err != nil {
				err2 := m.Close()
				if err2 != nil {
					panic(err2)
				}
				panic(err)
			}

		}

	})
}

type ChessAI struct {
	difficulty int
	engine     *uci.Engine
}

func NewChessAI(difficulty int) *ChessAI {
	if difficulty == 0 {
		return &ChessAI{difficulty: difficulty, engine: nil}
	} else {
		return &ChessAI{difficulty: difficulty, engine: getChessUCIEngine(difficulty)}
	}
}

func getChessUCIEngine(difficulty int) *uci.Engine {
	engine, err := uci.New("/home/nathan/Projects/chessGame/Pleco/target/release/pleco")
	if err != nil {
		panic(err)
	}
	if difficulty == 1 {
		// set up engine to use stockfish exe
		eng, err := uci.New("/home/nathan/Projects/chessGame/Pleco/target/release/pleco")
		if err != nil {
			panic(err)
		}
		defer func(eng *uci.Engine) {
			err := eng.Close()
			if err != nil {
				panic(err)
			}
		}(eng)
		// initialize uci with new game
		if err := eng.Run(uci.CmdSetOption{Name: "UCI_LimitStrength", Value: "true"},
			uci.CmdSetOption{Name: "UCI_Elo", Value: "1320"},
			uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
			panic(err)
		}
	} else if difficulty == 2 {
		// set up engine to use stockfish exe
		eng, err := uci.New("/home/nathan/Projects/chessGame/Pleco/target/release/pleco")
		if err != nil {
			panic(err)
		}
		defer func(eng *uci.Engine) {
			err := eng.Close()
			if err != nil {
				panic(err)
			}
		}(eng)
		// initialize uci with new game
		if err := eng.Run(uci.CmdSetOption{Name: "UCI_LimitStrength", Value: "true"},
			uci.CmdSetOption{Name: "UCI_Elo", Value: "2255"},
			uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
			panic(err)
		}
	} else {
		// set up engine to use stockfish exe
		eng, err := uci.New("/home/nathan/Projects/chessGame/Pleco/target/release/pleco")
		if err != nil {
			panic(err)
		}
		defer func(eng *uci.Engine) {
			err := eng.Close()
			if err != nil {
				panic(err)
			}
		}(eng)
		// initialize uci with new game
		if err := eng.Run(uci.CmdSetOption{Name: "UCI_LimitStrength", Value: "true"},
			uci.CmdSetOption{Name: "UCI_Elo", Value: "3190"},
			uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
			panic(err)
		}
	}
	return engine
}

func (ai *ChessAI) MakeMove(game *chess.Game) *chess.Move {
	if ai.engine == nil {
		moves := game.ValidMoves()
		move := moves[rand.Intn(len(moves))]
		if err := game.Move(move); err != nil {
			panic(err)
		}
		return move
	} else {
		// have stockfish play (10 msec per move)
		cmdPos := uci.CmdPosition{Position: game.Position()}
		cmdGo := uci.CmdGo{MoveTime: time.Second / 100}
		if err := ai.engine.Run(cmdPos, cmdGo); err != nil {
			panic(err)
		}
		move := ai.engine.SearchResults().BestMove
		if err := game.Move(move); err != nil {
			panic(err)
		}
		return move
	}
}

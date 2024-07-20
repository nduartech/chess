package engine

import (
	"db"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
	"github.com/olahol/melody"
)

type MelodyInterface interface {
	HandleConnect(fn func(*melody.Session))
	HandleDisconnect(fn func(*melody.Session))
	HandlePong(fn func(*melody.Session))
	HandleMessage(fn func(*melody.Session, []byte))
	HandleMessageBinary(fn func(*melody.Session, []byte))
	HandleSentMessage(fn func(*melody.Session, []byte))
	HandleSentMessageBinary(fn func(*melody.Session, []byte))
	HandleError(fn func(*melody.Session, error))
	HandleClose(fn func(*melody.Session, int, string) error)
	HandleRequest(w http.ResponseWriter, r *http.Request) error
	HandleRequestWithKeys(w http.ResponseWriter, r *http.Request, keys map[string]any) error
	Broadcast(msg []byte) error
	BroadcastFilter(msg []byte, fn func(*melody.Session) bool) error
	BroadcastOthers(msg []byte, s *melody.Session) error
	BroadcastMultiple(msg []byte, sessions []*melody.Session) error
	BroadcastBinary(msg []byte) error
	BroadcastBinaryFilter(msg []byte, fn func(*melody.Session) bool) error
	BroadcastBinaryOthers(msg []byte, s *melody.Session) error
	Sessions() ([]*melody.Session, error)
	Close() error
	CloseWithMsg(msg []byte) error
	Len() int
	IsClosed() bool
}

func ChessMatchWithBot(difficulty int, playerWhite bool, turn bool, m MelodyInterface, g *db.GameStateRow, d db.Conn) {
	chessBot := NewChessAI(difficulty)
	var game *chess.Game
	if g == nil {
		game = chess.NewGame()
		gs := &db.GameState{Local: true, PlayerSide: playerWhite, Turn: turn, BotDifficulty: difficulty, Started: time.Now(), Pgn: game.String()}
		index, err := d.AddGame(gs)
		if err != nil {
			panic(err)
		}
		g = &db.GameStateRow{ID: index, GameState: *gs}
	} else {
		pgn, err := chess.PGN(strings.NewReader(g.Pgn))
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
		if turn {
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
			g.Pgn = game.String()
			g.Turn = !g.Turn
			g.Ended = game.Outcome() != chess.NoOutcome
			err := d.UpdateGame(g)
			if err != nil {
				panic(err)
			}
			err = session.Write([]byte(move.String()))
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
		session.Set("game", game)
	})

	m.HandleDisconnect(func(session *melody.Session) {
		g.Pgn = game.String()
		g.Ended = game.Outcome() != chess.NoOutcome
		err := d.UpdateGame(g)
		if err != nil {
			panic(err)
		}
	})

	m.HandleMessage(func(session *melody.Session, bytes []byte) {
		gameSession, ok := session.Get("game")
		if !ok {
			panic("game not found")
		}
		game := gameSession.(*chess.Game)
		err := game.MoveStr(string(bytes))
		if err != nil {
			err2 := m.Close()
			if err2 != nil {
				panic(err2)
			}
			panic(err)
		}
		g.Pgn = game.String()
		g.Turn = !g.Turn
		g.Ended = game.Outcome() != chess.NoOutcome
		err = d.UpdateGame(g)
		if err != nil {
			panic(err)
		}

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
			g.Pgn = game.String()
			g.Turn = !g.Turn
			g.Ended = game.Outcome() != chess.NoOutcome
			err = d.UpdateGame(g)
			if err != nil {
				panic(err)
			}

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
		session.Set("game", game)
	})
}

type ChessAI interface {
	MakeMove(game *chess.Game) *chess.Move
}

type ChessAIImpl struct {
	difficulty int
	engine     *uci.Engine
}

func NewChessAI(difficulty int) ChessAI {
	if difficulty == 0 {
		return &ChessAIImpl{difficulty: difficulty, engine: nil}
	} else {
		return &ChessAIImpl{difficulty: difficulty, engine: getChessUCIEngine(difficulty)}
	}
}

func getChessUCIEngine(difficulty int) *uci.Engine {
	engine, err := uci.New("stockfish")
	if err != nil {
		panic(err)
	}
	if difficulty == 1 {
		// set up engine to use stockfish exe
		eng, err := uci.New("stockfish")
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
		eng, err := uci.New("stockfish")
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
		eng, err := uci.New("stockfish")
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

func (ai *ChessAIImpl) MakeMove(game *chess.Game) *chess.Move {
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

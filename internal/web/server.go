package web

import (
	"components"
	"db"
	"engine"
	"fmt"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/olahol/melody"
)

var local *db.Conn

func Run() {
	homeComponent := Home()
	newGameComponent := components.NewGameOptions()
	botDifficulty := components.ChooseBotDifficulty()

	m := melody.New()
	s := engine.NewSymphony(m)
	defer func(m *engine.Symphony) {
		err := m.Close()
		if err != nil {
			panic(err)
		}
	}(s)
	checkForGame := func(next http.Handler) http.Handler {
		local, err := db.InitGameDatabase("./.yaca.sqlite3")
		if err != nil {
			panic(err)
		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gameState, err := local.GetRecentGame()
			if err != nil || gameState.Ended {
				next.ServeHTTP(w, r)
				return
			}
			bot := components.NewBotGame(gameState.BotDifficulty, gameState.PlayerSide, gameState.Turn, s, gameState, local)
			templ.Handler(bot).ServeHTTP(w, r)
		})
	}

	http.Handle("/", checkForGame(templ.Handler(homeComponent)))
	http.Handle("/new-game", checkForGame(templ.Handler(newGameComponent)))
	http.Handle("/set-difficulty", checkForGame(templ.Handler(botDifficulty)))

	http.HandleFunc("/bot-game/{difficulty}", func(w http.ResponseWriter, r *http.Request) {
		diff, err := strconv.Atoi(r.PathValue("difficulty"))
		if err != nil {
			fmt.Println(err)
			http.NotFoundHandler().ServeHTTP(w, r)
		}
		page := components.ChooseSideBot(diff)
		checkForGame(templ.Handler(page)).ServeHTTP(w, r)
	})
	http.HandleFunc("/bot-game/{difficulty}/{color}", func(w http.ResponseWriter, r *http.Request) {
		local, err := db.InitGameDatabase("./local.sqlite3")
		if err != nil {
			panic(err)
		}
		gameState, err := local.GetRecentGame()
		if err != nil {
			fmt.Println(err)
			gameState = nil
		}
		if err == nil && gameState.Ended {
			gameState = nil
		}
		diff, err := strconv.Atoi(r.PathValue("difficulty"))
		if err != nil {
			fmt.Println(err)
			http.NotFoundHandler().ServeHTTP(w, r)
		}
		bot := components.NewBotGame(diff, r.PathValue("color") == "white", r.PathValue("color") == "white", s, gameState, local)
		templ.Handler(bot).ServeHTTP(w, r)
	})

	fmt.Println("Listening on :38538")
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)
		if err != nil {
			panic(err)
		}
	})
	err := http.ListenAndServe(":38538", nil)
	if err != nil {
		panic(err)
	}
}

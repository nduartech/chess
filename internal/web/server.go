package web

import (
	"components"
	"db"
	"fmt"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/olahol/melody"
)

var local *db.DBConn

// Need to find a way to maintain state between requests

func Run() {
	homeComponent := Home()
	newGameComponent := components.NewGameOptions()
	botDifficulty := components.ChooseBotDifficulty()

	m := melody.New()
	defer func(m *melody.Melody) {
		err := m.Close()
		if err != nil {
			panic(err)
		}
	}(m)
	checkForGame := func(next http.Handler) http.Handler {
		local, err := db.InitDatabase("./local.sqlite3")
		if err != nil {
			panic(err)
		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gameState, err := local.GetRecentGame()
			if err != nil || gameState.GetEnded() {
				next.ServeHTTP(w, r)
				return
			}
			bot := components.NewBotGame(gameState.GetBotDifficulty(), gameState.GetPlayerSide(), m, gameState, local)
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
		local, err := db.InitDatabase("./local.sqlite3")
		if err != nil {
			panic(err)
		}
		gameState, err := local.GetRecentGame()
		if err != nil {
			fmt.Println(err)
			gameState = nil
		}
		if err == nil && gameState.GetEnded() {
			gameState = nil
		}
		diff, err := strconv.Atoi(r.PathValue("difficulty"))
		if err != nil {
			fmt.Println(err)
			http.NotFoundHandler().ServeHTTP(w, r)
		}
		bot := components.NewBotGame(diff, r.PathValue("color") == "white", m, gameState, local)
		templ.Handler(bot).ServeHTTP(w, r)
	})

	fmt.Println("Listening on :3000")
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)
		if err != nil {
			panic(err)
		}
	})
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

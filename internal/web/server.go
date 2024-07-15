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
	easiestSide := components.ChooseSideBot(0)
	easySide := components.ChooseSideBot(1)
	mediumSide := components.ChooseSideBot(2)
	HardSide := components.ChooseSideBot(3)
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
		defer local.Close()
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gameState, err := local.GetRecentGame()
			if err != nil {
				fmt.Println(err)
				next.ServeHTTP(w, r)
				return
			}
			bot := components.NewBotGame(gameState.GetBotDifficulty(), gameState.GetPlayerSide(), m, gameState)
			templ.Handler(bot).ServeHTTP(w, r)
		})
	}
	// easiestBotW := components.NewBotGame(0, true, m, nil)
	// easiestBotB := components.NewBotGame(0, false, m, nil)
	// easyBotW := components.NewBotGame(1, true, m, nil)
	// easyBotB := components.NewBotGame(1, false, m, nil)
	// mediumBotW := components.NewBotGame(2, true, m, nil)
	// mediumBotB := components.NewBotGame(2, false, m, nil)
	// hardBotW := components.NewBotGame(3, true, m, nil)
	// hardBotB := components.NewBotGame(3, false, m, nil)
	http.Handle("/", checkForGame(templ.Handler(homeComponent)))
	http.Handle("/new-game", checkForGame(templ.Handler(newGameComponent)))
	http.Handle("/set-difficulty", checkForGame(templ.Handler(botDifficulty)))
	http.Handle("/bot-game/0", checkForGame(templ.Handler(easiestSide)))
	http.Handle("/bot-game/1", checkForGame(templ.Handler(easySide)))
	http.Handle("/bot-game/2", checkForGame(templ.Handler(mediumSide)))
	http.Handle("/bot-game/3", checkForGame(templ.Handler(HardSide)))
	http.HandleFunc("/bot-game/{difficulty}/{color}", func(w http.ResponseWriter, r *http.Request) {
		local, err := db.InitDatabase("./local.sqlite3")
		if err != nil {
			panic(err)
		}
		defer local.Close()
		gameState, err := local.GetRecentGame()
		if err != nil {
			fmt.Println(err)
			gameState = nil
		}
		diff, err := strconv.Atoi(r.PathValue("difficulty"))
		if err != nil {
			fmt.Println(err)
			http.NotFoundHandler().ServeHTTP(w, r)
		}
		bot := components.NewBotGame(diff, r.PathValue("color") == "white", m, gameState)
		templ.Handler(bot).ServeHTTP(w, r)
	})
	// http.Handle("/bot-game/0/black", templ.Handler(easiestBotB))
	// http.Handle("/bot-game/1/white", templ.Handler(easyBotW))
	// http.Handle("/bot-game/1/black", templ.Handler(easyBotB))
	// http.Handle("/bot-game/2/white", templ.Handler(mediumBotW))
	// http.Handle("/bot-game/2/black", templ.Handler(mediumBotB))
	// http.Handle("/bot-game/3/white", templ.Handler(hardBotW))
	// http.Handle("/bot-game/3/black", templ.Handler(hardBotB))

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

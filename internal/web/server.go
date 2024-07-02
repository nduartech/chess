package web

import (
	"components"
	"fmt"
	"github.com/a-h/templ"
	"github.com/olahol/melody"
	"net/http"
)

func Run() {
	m := melody.New()
	defer func(m *melody.Melody) {
		err := m.Close()
		if err != nil {
			panic(err)
		}
	}(m)
	homeComponent := Home()
	newGameComponent := components.NewGameOptions()
	botDifficulty := components.ChooseBotDifficulty()
	easiestSide := components.ChooseSideBot(0)
	easySide := components.ChooseSideBot(1)
	mediumSide := components.ChooseSideBot(2)
	HardSide := components.ChooseSideBot(3)
	easiestBotW := components.NewBotGame(0, true, m)
	easiestBotB := components.NewBotGame(0, false, m)
	easyBotW := components.NewBotGame(1, true, m)
	easyBotB := components.NewBotGame(1, false, m)
	mediumBotW := components.NewBotGame(2, true, m)
	mediumBotB := components.NewBotGame(2, false, m)
	hardBotW := components.NewBotGame(3, true, m)
	hardBotB := components.NewBotGame(3, false, m)
	http.Handle("/", templ.Handler(homeComponent))
	http.Handle("/new-game", templ.Handler(newGameComponent))
	http.Handle("/set-difficulty", templ.Handler(botDifficulty))
	http.Handle("/bot-game/0", templ.Handler(easiestSide))
	http.Handle("/bot-game/1", templ.Handler(easySide))
	http.Handle("/bot-game/2", templ.Handler(mediumSide))
	http.Handle("/bot-game/3", templ.Handler(HardSide))
	http.Handle("/bot-game/0/white", templ.Handler(easiestBotW))
	http.Handle("/bot-game/0/black", templ.Handler(easiestBotB))
	http.Handle("/bot-game/1/white", templ.Handler(easyBotW))
	http.Handle("/bot-game/1/black", templ.Handler(easyBotB))
	http.Handle("/bot-game/2/white", templ.Handler(mediumBotW))
	http.Handle("/bot-game/2/black", templ.Handler(mediumBotB))
	http.Handle("/bot-game/3/white", templ.Handler(hardBotW))
	http.Handle("/bot-game/3/black", templ.Handler(hardBotB))
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

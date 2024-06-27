package main

import (
	"internal/web"
)

func main() {

	//openedInBrowser := os.Getenv("OPENED_BROWSER")
	//if openedInBrowser != "true" {
	//	err := exec.Command("open", "http://localhost:3000").Run()
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	web.Run()
}

//func main() {
//	sampleAI := NewChessAI(1)
//	game := chess.NewGame()
//	playerTurn := false
//	for game.Outcome() == chess.NoOutcome {
//		if playerTurn {
//			moves := game.ValidMoves()
//			move := moves[rand.Intn(len(moves))]
//			fmt.Println(move)
//			game.MakeMove(move)
//		} else {
//			_ = sampleAI.MakeMove(game) //TODO:transmit move to UI for animation
//		}
//		playerTurn = !playerTurn
//	}
//	// print outcome and game PGN
//	fmt.Println(game.Position().Board().Draw())
//	fmt.Printf("Game completed. %s by %s.\n", game.Outcome(), game.Method())
//	fmt.Println(game.String())
//	/*
//		Output:
//
//		 A B C D E F G H
//		8- - - - - - - -
//		7- - - - - - ♚ -
//		6- - - - ♗ - - -
//		5- - - - - - - -
//		4- - - - - - - -
//		3♔ - - - - - - -
//		2- - - - - - - -
//		1- - - - - - - -
//
//		Game completed. 1/2-1/2 by InsufficientMaterial.
//
//		1.Nc3 b6 2.a4 e6 3.d4 Bb7 ...
//	*/
//}

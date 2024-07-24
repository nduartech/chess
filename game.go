package main

import (
	"bytes"
	"errors"
	"fmt"
	"internal/web"
	"log"
	"os"
	"os/exec"
)

func main() {
	if _, err := os.Stat("./Stockfish/src/stockfish"); errors.Is(err, os.ErrNotExist) {
		e := exec.Command("make", "-j", "profile-build")
		e.Dir = "./Stockfish/src"
		var out bytes.Buffer
		e.Stdout = &out
		err := e.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Output: %q\n", out.String())
	}
	web.Run()
}

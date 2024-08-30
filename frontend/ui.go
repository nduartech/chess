package main

import webview "github.com/webview/webview_go"

func main() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("YACA")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://localhost:38538")
	w.Run()
}
